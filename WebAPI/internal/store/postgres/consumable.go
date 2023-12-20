package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

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
