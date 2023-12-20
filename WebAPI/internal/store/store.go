package store

import imp "github.com/jtcarden0001/personacmms/webapi/internal/store/postgres"

// my interfaces are not implementation specific but it's easier keeping everything in one file.
// I will move these to a common location if I field the idea of another implementation or find a need
type Store interface {
	imp.Equipment
	imp.Tool
	imp.EquipmentCategory
}

type StoreTest interface {
	Store
	imp.EquipmentTest
	imp.ToolTest
	imp.EquipmentCategoryTest
}

func New() *imp.Store {
	return imp.New()
}

func NewTest() StoreTest {
	return imp.NewTest()
}
