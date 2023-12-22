package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type Consumable interface {
	CreateConsumable(string) (int, error)
	DeleteConsumable(int) error
	GetAllConsumable() ([]tp.Consumable, error)
	GetConsumable(int) (tp.Consumable, error)
	UpdateConsumable(int, string) error
}

func (a *App) CreateConsumable(name string) (int, error) {
	return a.db.CreateConsumable(name)
}

func (a *App) DeleteConsumable(id int) error {
	return a.db.DeleteConsumable(id)
}

func (a *App) GetAllConsumable() ([]tp.Consumable, error) {
	return a.db.GetAllConsumable()
}

func (a *App) GetConsumable(id int) (tp.Consumable, error) {
	return a.db.GetConsumable(id)
}

func (a *App) UpdateConsumable(id int, name string) error {
	return a.db.UpdateConsumable(id, name)
}
