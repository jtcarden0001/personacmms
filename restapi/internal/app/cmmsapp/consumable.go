package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

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
