package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
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

// handling nil pointers aka null columns
// ConvertInt64Pointer converts an int64 pointer to sql.NullInt64
func ConvertInt64Pointer(ptr *int64) sql.NullInt64 {
	if ptr != nil {
		return sql.NullInt64{Int64: *ptr, Valid: true}
	}
	return sql.NullInt64{Valid: false}
}

// ConvertStringPointer converts a string pointer to sql.NullString
func ConvertStringPointer(ptr *string) sql.NullString {
	if ptr != nil {
		return sql.NullString{String: *ptr, Valid: true}
	}
	return sql.NullString{Valid: false}
}

// ConvertFloat64Pointer converts a float64 pointer to sql.NullFloat64
func ConvertFloat64Pointer(ptr *float64) sql.NullFloat64 {
	if ptr != nil {
		return sql.NullFloat64{Float64: *ptr, Valid: true}
	}
	return sql.NullFloat64{Valid: false}
}

// ConvertUUIDPointer converts a uuid.UUID pointer to sql.NullString
func ConvertUUIDPointer(ptr *uuid.UUID) sql.NullString {
	if ptr != nil {
		return sql.NullString{String: ptr.String(), Valid: true}
	}
	return sql.NullString{Valid: false}
}

// ConvertBoolPointer converts a bool pointer to sql.NullBool
func ConvertBoolPointer(ptr *bool) sql.NullBool {
	if ptr != nil {
		return sql.NullBool{Bool: *ptr, Valid: true}
	}
	return sql.NullBool{Valid: false}
}

// ConvertTimePointer converts a time.Time pointer to sql.NullTime
func ConvertTimePointer(ptr *time.Time) sql.NullTime {
	if ptr != nil {
		return sql.NullTime{Time: *ptr, Valid: true}
	}
	return sql.NullTime{Valid: false}
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
