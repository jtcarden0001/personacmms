package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	Db = getDB()
}

func getDB() *sql.DB {
	pgPass := os.Getenv("postgrespass")
	connStr := fmt.Sprintf("postgresql://postgres:%s@localhost/personacmms?sslmode=disable", pgPass)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
