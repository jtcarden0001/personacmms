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
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}

func NewTest(injectedStore st.StoreTest) AppTest {
	return imp.NewTest(injectedStore)
}
