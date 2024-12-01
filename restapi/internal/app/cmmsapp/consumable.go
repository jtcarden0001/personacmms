package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

// A Consumable is a consumable item that is required to complete a maintenance task for an asset.

// Create a Consumable
func (a *App) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	return a.db.CreateConsumable(consumable)
}

// Delete a Consumable
func (a *App) DeleteConsumable(consumableTitle string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetConsumable(consumableTitle); err != nil {
		return err
	}

	return a.db.DeleteConsumable(consumableTitle)
}

// List all Consumables
func (a *App) ListConsumables() ([]tp.Consumable, error) {
	return a.db.ListConsumables()
}

// Get a Consumable
func (a *App) GetConsumable(consumableTitle string) (tp.Consumable, error) {
	return a.db.GetConsumable(consumableTitle)
}

// Update a Consumable
func (a *App) UpdateConsumable(oldTitle string, consumable tp.Consumable) (tp.Consumable, error) {
	if consumable.Title == "" {
		consumable.Title = oldTitle
	}

	return a.db.UpdateConsumable(oldTitle, consumable)
}
