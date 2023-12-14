package db

import imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/db/postgres"

func New() *imp.Store {
	return imp.New()
}

func NewTest() StoreTest {
	return imp.NewTest()
}
