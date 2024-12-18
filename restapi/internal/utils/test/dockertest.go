package test

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"
)

// creates a postgres db in a docker container for testing. add the following code for cleanup after
// calling this function:
//
//	defer func() {
//		if err := pool.Purge(resource); err != nil {
//			log.Fatalf("Could not purge resource: %s", err)
//		}
//		resource.Close()
//	}()
func CreateDockerTestPostgres() (*dockertest.Pool, *dockertest.Resource, error) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Could not construct dockertest pool")
	}

	err = pool.Client.Ping()
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Could not connect to dockertest pool")
	}

	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_PASSWORD", "test")
	os.Setenv("DATABASE_NAME", "test")
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.13-alpine3.20",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", os.Getenv("DATABASE_PASSWORD")),
			fmt.Sprintf("POSTGRES_USER=%s", os.Getenv("DATABASE_USER")),
			fmt.Sprintf("POSTGRES_DB=%s", os.Getenv("DATABASE_NAME")),
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Could not start dockettest resource")
	}
	defer func() {
		if err != nil {
			resource.Close()
			if perr := pool.Purge(resource); perr != nil {
				log.Fatalf("Could not purge resource: %s", perr)
			}
		}
	}()

	hostAndPort := resource.GetHostPort("5432/tcp")
	host, port, err := net.SplitHostPort(hostAndPort)
	if err != nil {
		log.Fatalf("Invalid host:port format: %s", err)
	}

	os.Setenv("DATABASE_HOST", host)
	os.Setenv("DATABASE_PORT", port)
	databaseUrl := fmt.Sprintf("postgres://test:test@%s/test?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseUrl)
	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	var db *sql.DB
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return nil, nil, errors.Wrapf(err, "Could not connect to dockertest postgres instance")
	}

	return pool, resource, nil
}
