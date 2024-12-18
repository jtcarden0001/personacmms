package test

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	_ "github.com/lib/pq"
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

func InitializeStore(dbName string) st.Store {
	conninfo := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"))
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Fatal(err)
	}

	store := st.NewWithDb(dbName)
	schema, err := os.ReadFile("../../postgres/ddl/init.sql")
	if err != nil {
		log.Fatalf("Could not read schema file: %s", err)
	}

	err = store.Exec(string(schema))
	if err != nil {
		log.Fatalf("Could not execute schema: %s", err)
	}

	return store
}

func CloseStore(store st.Store, dbName string) {
	// cannnot drop the database because the connection is still open
	// store.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	_ = dbName
	err := store.Close()
	if err != nil {
		log.Fatalf("Could not close store: %s", err)
	}
}
