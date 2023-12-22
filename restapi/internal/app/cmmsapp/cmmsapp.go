package cmmsapp

import st "github.com/jtcarden0001/personacmms/webapi/internal/store"

type App struct {
	db st.Store
}

func New(injectedStore st.Store) *App {
	return &App{db: injectedStore}
}

func NewTest(injectedStore st.StoreTest) *App {
	return &App{db: injectedStore}
}
