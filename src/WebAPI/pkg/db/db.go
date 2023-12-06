package db

import (
	"database/sql"

	imp "github.com/jocarde/personacmms/pkg/db/postgres"
)

var Db *sql.DB

func init() {
	Db = imp.GetDB()
}
