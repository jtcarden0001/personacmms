package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

// these interfaces are platform agnostic but I like having them in the same file as the implementation,
// will move to a common location if I field the idea of another implemntation
type Consumable interface {
	CreateConsumable(string) (int, error)
	DeleteConsumable(int) error
	GetAllConsumable() ([]tp.EquipmentCategory, error)
	GetConsumable(int) (tp.EquipmentCategory, error)
	UpdateConsumable(int, string) error
}

type ConsumableTest interface {
	ResetSequenceConsumable(int) error
}
