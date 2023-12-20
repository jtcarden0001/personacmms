package app

import (
	imp "github.com/jtcarden0001/personacmms/webapi/internal/app/cmmsapp"
	"github.com/jtcarden0001/personacmms/webapi/internal/store"
)

type App interface {
	imp.Equipment
	imp.Tool
}

type AppTest interface {
	App
	imp.EquipmentTest
	imp.ToolTest
}

func New(injectedStore store.Store) *imp.App {
	return imp.New(injectedStore)
}

func NewTest(injectedStore store.StoreTest) *imp.AppTest {
	return imp.NewTest(injectedStore)
}
