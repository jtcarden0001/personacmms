package cmmsapp

import "github.com/jtcarden0001/personacmms/webapi/internal/store"

type App struct {
	db store.Store
}

func New(injectedStore store.Store) *App {
	return &App{db: injectedStore}
}

type AppTest struct {
	db store.StoreTest
}

func NewTest(injectedStore store.StoreTest) *AppTest {
	return &AppTest{db: injectedStore}
}
