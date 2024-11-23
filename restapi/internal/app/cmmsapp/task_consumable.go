package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type PreventativeTaskConsumable interface {
	CreatePreventativeTaskConsumable(int, int, string) error
	DeletePreventativeTaskConsumable(int, int) error
	GetAllPreventativeTaskConsumable() ([]tp.PreventativeTaskConsumable, error)
	GetAllPreventativeTaskConsumableByPreventativeTaskId(int) ([]tp.PreventativeTaskConsumable, error)
	GetPreventativeTaskConsumable(int, int) (tp.PreventativeTaskConsumable, error)
	UpdatePreventativeTaskConsumable(int, int, string) error
}

func (a *App) CreatePreventativeTaskConsumable(preventativeTaskId int, consumableId int, quantity string) error {
	return a.db.CreatePreventativeTaskConsumable(preventativeTaskId, consumableId, quantity)
}

func (a *App) DeletePreventativeTaskConsumable(preventativeTaskId int, consumableId int) error {
	return a.db.DeletePreventativeTaskConsumable(preventativeTaskId, consumableId)
}

func (a *App) GetAllPreventativeTaskConsumable() ([]tp.PreventativeTaskConsumable, error) {
	return a.db.GetAllPreventativeTaskConsumable()
}

func (a *App) GetAllPreventativeTaskConsumableByPreventativeTaskId(preventativeTaskId int) ([]tp.PreventativeTaskConsumable, error) {
	return a.db.GetAllPreventativeTaskConsumableByPreventativeTaskId(preventativeTaskId)
}

func (a *App) GetPreventativeTaskConsumable(preventativeTaskId int, consumableId int) (tp.PreventativeTaskConsumable, error) {
	return a.db.GetPreventativeTaskConsumable(preventativeTaskId, consumableId)
}

func (a *App) UpdatePreventativeTaskConsumable(preventativeTaskId int, consumableId int, quantity string) error {
	return a.db.UpdatePreventativeTaskConsumable(preventativeTaskId, consumableId, quantity)
}
