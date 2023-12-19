package store

import imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/store/postgres"

func New() *imp.Store {
	return imp.New()
}

func NewTest() StoreTest {
	return imp.NewTest()
}

// my interfaces are not implementation specific but it's easier keeping everything in one file.
// I will move these to a common location if I field the idea of another implemntation or find a need
type Store interface {
	imp.Equipment
	imp.Tool
	imp.EquipmentCategory
}

type StoreTest interface {
	imp.Equipment
	imp.EquipmentTest
	imp.Tool
	imp.ToolTest
	imp.EquipmentCategory
	imp.EquipmentCategoryTest
}
