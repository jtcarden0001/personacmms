package cmmsapp

import st "github.com/jtcarden0001/personacmms/webapi/internal/store"

type App struct {
	db st.Store
}

func New(injectedStore st.Store) *App {
	return &App{db: injectedStore}
}

type AppTest struct {
	db st.StoreTest
}

func NewTest(injectedStore st.StoreTest) *AppTest {
	return &AppTest{db: injectedStore}
}

func (cmms *AppTest) ResetSequence(table string, nextVal int) error {
	return cmms.db.ResetSequence(table, nextVal)
}
