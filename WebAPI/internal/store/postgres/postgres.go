package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

type Store struct {
	name string
	db   *sql.DB
}

var prodDbName = "personacmms"
var testDbName = "personacmms-test"

func New() *Store {
	return getStore(prodDbName)
}

func NewTest() *Store {
	return getStore(testDbName)
}

func getStore(dbName string) *Store {
	pgPass := os.Getenv("postgrespass")
	connStr := fmt.Sprintf("postgresql://postgres:%s@localhost/%s?sslmode=disable", pgPass, dbName)
	// Connect to database
	Db, err := sql.Open("postgres", connStr)
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

// this might be a little heavy handed to nuke the entire db and recreate it instead of just deleting the data and resetting the sequences
// but will leave unless a reason surfaces to change it.
func (pg *Store) CleanTestStore() error {
	// very important to prevent accidental deletion of production data
	if pg.name != testDbName {
		return fmt.Errorf("cleaning failed on db: %s, cleaning is only allowable on db: %s", pg.name, testDbName)
	}

	ddlPath, err := filepath.Abs("../../../../tools/db/mvp_postgres_ddl.sql")
	if err != nil {
		return err
	}

	ddlBuffer, err := os.ReadFile(ddlPath)
	if err != nil {
		return err
	}

	ddl := string(ddlBuffer)
	_, err = pg.db.Exec(ddl)
	if err != nil {
		return err
	}

	return nil
}
