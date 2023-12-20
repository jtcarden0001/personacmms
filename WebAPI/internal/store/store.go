package store

import imp "github.com/jtcarden0001/personacmms/webapi/internal/store/postgres"

// these embedded interfaces are not implementation specific but it's easier keeping related code together.
// I will move these to a common location (out of imp) if I field the idea of another implementation or find a need.
type Store interface {
	imp.Equipment
	imp.Tool
	imp.EquipmentCategory
}

type StoreTest interface {
	Store
	ResetSequence(string, int) error
	CleanTestStore() error
}

func New() *imp.Store {
	return imp.New()
}

func NewTest() StoreTest {
	return imp.NewTest()
}
