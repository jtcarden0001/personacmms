package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// TODO: switch to *sql.Db calls using contexts

type PostgresStore struct {
	name string
	db   *sql.DB
}

// TODO: All of the entity store code is structured almost exactly the same.  Probably a way to reduce code using generics or something.
func New() *PostgresStore {
	return getStore("")
}

// used for testing
func NewWithDb(dbName string) *PostgresStore {
	return getStore(dbName)
}

func getStore(dbName string) *PostgresStore {
	// TODO: move this config out of the implementation code
	pgUser := os.Getenv("DATABASE_USER")
	pgPass := os.Getenv("DATABASE_PASSWORD")
	pgHost := os.Getenv("DATABASE_HOST")
	pgPort := os.Getenv("DATABASE_PORT")
	if dbName == "" {
		dbName = os.Getenv("DATABASE_NAME")
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, pgHost, pgPort, dbName)

	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		name: dbName,
		db:   Db,
	}
}

func (pg *PostgresStore) Exec(query string) error {
	_, err := pg.db.Exec(query)
	return err
}

func (pg *PostgresStore) Close() error {
	return pg.db.Close()
}
