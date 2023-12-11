package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Store struct {
	db *sql.DB
}

func New() *Store {
	pgPass := os.Getenv("postgrespass")
	connStr := fmt.Sprintf("postgresql://postgres:%s@localhost/personacmms?sslmode=disable", pgPass)
	// Connect to database
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return &Store{db: Db}
}
