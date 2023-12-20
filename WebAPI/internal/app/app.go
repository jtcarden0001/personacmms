package app

import (
	imp "github.com/jtcarden0001/personacmms/webapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/webapi/internal/store"
)

type App interface {
	imp.Equipment
	imp.Tool
}

type AppTest interface {
	App
	ResetSequence(string, int) error
}

func New(injectedStore st.Store) *imp.App {
	return imp.New(injectedStore)
}

func NewTest(injectedStore st.StoreTest) *imp.AppTest {
	return imp.NewTest(injectedStore)
}
