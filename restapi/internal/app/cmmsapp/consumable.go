package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Consumable interface {
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(string) error
	ListConsumables() ([]tp.Consumable, error)
	GetConsumable(string) (tp.Consumable, error)
	UpdateConsumable(string, tp.Consumable) (tp.Consumable, error)
}

func (a *App) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	return a.db.CreateConsumable(consumable)
}

func (a *App) DeleteConsumable(consumableTitle string) error {
	return a.db.DeleteConsumable(consumableTitle)
}

func (a *App) ListConsumables() ([]tp.Consumable, error) {
	return a.db.ListConsumables()
}

func (a *App) GetConsumable(consumableTitle string) (tp.Consumable, error) {
	return a.db.GetConsumable(consumableTitle)
}

func (a *App) UpdateConsumable(consumableTitle string, consumable tp.Consumable) (tp.Consumable, error) {
	return a.db.UpdateConsumable(consumableTitle, consumable)
}
