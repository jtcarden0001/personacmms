package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Store struct {
	name string
	db   *sql.DB
}

// TODO: All of the entity store code is structured almost exactly the same.  Probably a way to reduce code using generics or something.
func New() *Store {
	return getStore("")
}

// used for testing
func NewWithDb(dbName string) *Store {
	return getStore(dbName)
}

func getStore(dbName string) *Store {
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

	return &Store{
		name: dbName,
		db:   Db,
	}
}

func (pg *Store) Exec(query string) error {
	_, err := pg.db.Exec(query)
	return err
}

func (pg *Store) Close() error {
	return pg.db.Close()
}
