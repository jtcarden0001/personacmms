package app

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/app/cmmsapp"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/db"
)

func New(store db.Store) *imp.App {
	return imp.New(store)
}
