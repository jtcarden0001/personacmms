package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	return a.db.CreateConsumable(consumable)
}

func (a *App) DeleteConsumable(consumableTitle string) error {
	if _, err := a.GetConsumable(consumableTitle); err != nil {
		return err
	}

	return a.db.DeleteConsumable(consumableTitle)
}

func (a *App) ListConsumables() ([]tp.Consumable, error) {
	return a.db.ListConsumables()
}

func (a *App) GetConsumable(consumableTitle string) (tp.Consumable, error) {
	return a.db.GetConsumable(consumableTitle)
}

func (a *App) UpdateConsumable(oldTitle string, consumable tp.Consumable) (tp.Consumable, error) {
	if consumable.Title == "" {
		consumable.Title = oldTitle
	}
	return a.db.UpdateConsumable(oldTitle, consumable)
}
