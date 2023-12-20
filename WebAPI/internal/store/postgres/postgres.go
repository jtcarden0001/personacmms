package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New() *Store {
	return getStore("personacmms")
}

func NewTest() *Store {
	return getStore("personacmms-test")
}

// TODO: need to devise a strategy to scan nll values from the store withough exposing sql semantics throughout the app.
func getStore(dbName string) *Store {
	pgPass := os.Getenv("postgrespass")
	connStr := fmt.Sprintf("postgresql://postgres:%s@localhost/%s?sslmode=disable", pgPass, dbName)
	// Connect to database
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{db: Db}
}

func (pg *Store) ResetSequence(table string, id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH %d", table, id)
	_, err := pg.db.Exec(query)

	return err
}
