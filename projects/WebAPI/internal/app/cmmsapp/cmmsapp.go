package cmmsapp

import "github.com/jtcarden0001/personacmms/projects/webapi/internal/db"

type App struct {
	store db.Store
}

func New(injectedStore db.Store) *App {
	return &App{store: injectedStore}
}

type AppTest struct {
	store db.StoreTest
}

func NewTest(injectedStore db.StoreTest) *AppTest {
	return &AppTest{store: injectedStore}
}
