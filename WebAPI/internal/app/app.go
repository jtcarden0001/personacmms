package app

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/app/cmmsapp"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/store"
)

func New(db store.Store) *imp.App {
	return imp.New(db)
}

func NewTest(db store.StoreTest) *imp.AppTest {
	return imp.NewTest(db)
}
