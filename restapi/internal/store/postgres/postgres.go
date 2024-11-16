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

var prodDbName = "personacmms-prod"
var testDbName = "personacmms-test"

// All of the entity store code is structured almost exactly the same.  Probably a way to reduce code using generics or something.
// TODO: look into above
func New() *Store {
	return getStore(prodDbName)
}

func NewTest() *Store {
	return getStore(testDbName)
}

func getStore(dbName string) *Store {
	// TODO: move this config out of the implementation code
	pgUser := os.Getenv("DATABASE_USER")
	pgPass := os.Getenv("DATABASE_PASSWORD")
	pgHost := os.Getenv("DATABASE_HOST")
	//pgPort := os.Getenv("DATABASE_PORT")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", pgUser, pgPass, pgHost, dbName)

	Db, err := sql.Open(pgUser, connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		name: dbName,
		db:   Db,
	}
}

func (pg *Store) ResetSequence(table string, id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH %d", table, id)
	_, err := pg.db.Exec(query)

	return err
}

func (pg *Store) CleanTable(tableName string) error {
	// very important to prevent accidental deletion of production data
	if pg.name != testDbName {
		return fmt.Errorf("clean table for table %s failed on db: %s, cleaning is only allowable on db: %s", tableName, pg.name, testDbName)
	}

	query := fmt.Sprintf("DELETE FROM %s", tableName)
	_, err := pg.db.Exec(query)
	return err
}
