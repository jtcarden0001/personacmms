package test

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
)

// TODO: there is alot of postgres leaked from the store implementation. The clean approach
// will be to decouple the store test code from the implementation.  This is good enough for now.

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
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
		log.Fatalf("Could not start resource: %s", err)
	}
	defer resource.Close()

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
		log.Fatalf("Could not connect to docker: %s", err)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	// run tests
	m.Run()
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
	schema, err := os.ReadFile("../postgres/ddl/init.sql")
	if err != nil {
		log.Fatalf("Could not read schema file: %s", err)
	}

	err = store.Exec(string(schema))
	if err != nil {
		log.Fatalf("Could not execute schema: %s", err)
	}

	return store
}

func convertToSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return set
}

func compareEntitiesExcludingFields(t *testing.T, expected interface{}, actual interface{}, fields map[string]struct{}) {
	// Compare all properties except for the specified fields
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	for i := 0; i < expectedValue.NumField(); i++ {
		field := expectedValue.Type().Field(i)
		if _, ok := fields[field.Name]; ok {
			continue
		}

		expectedField := expectedValue.Field(i).Interface()
		actualField := actualValue.Field(i).Interface()

		if !reflect.DeepEqual(expectedField, actualField) {
			t.Errorf("Create() failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
		}
	}
}
