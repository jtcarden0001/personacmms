package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type Consumable interface {
	CreateConsumable(string) (int, error)
	DeleteConsumable(int) error
	GetAllConsumable() ([]tp.Consumable, error)
	GetConsumable(int) (tp.Consumable, error)
	UpdateConsumable(int, string) error
}
