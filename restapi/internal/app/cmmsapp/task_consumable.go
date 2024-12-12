package cmmsapp

import (
	"github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// TODO: implement this vertical

func (a *App) CreateTaskConsumable(atc tp.TaskConsumable) (tp.TaskConsumable, error) {
	if err := a.validateTaskConsumable(atc); err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.CreateTaskConsumable(atc)
}

func (a *App) CreateTaskConsumableWithValidation(groupTitle, assetTitle, taskId, consumableId, quantityNote string) (tp.TaskConsumable, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.CreateTaskConsumable(tp.TaskConsumable{TaskId: atId, ConsumableId: cId, QuantityNote: quantityNote})
}

func (a *App) DeleteTaskConsumable(groupTitle, assetTitle, taskId, consumableId string) error {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return err
	}

	return a.db.DeleteTaskConsumable(atId, cId)
}

func (a *App) ListTaskConsumables(groupTitle, assetTitle, taskId string) ([]tp.TaskConsumable, error) {
	// TODO: implement validation
	atcs, err := a.db.ListTaskConsumables()
	if err != nil {
		return nil, err
	}

	// TODO: filter asset task consumables by asset task id

	return atcs, nil
}

func (a *App) GetTaskConsumable(groupTitle, assetTitle, taskId, consumableId string) (tp.TaskConsumable, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.GetTaskConsumable(atId, cId)
}

func (a *App) UpdateTaskConsumable(atc tp.TaskConsumable) (tp.TaskConsumable, error) {
	return a.db.UpdateTaskConsumable(atc)
}

func (a *App) UpdateTaskConsumableWithValidation(groupTitle, assetTitle, taskId, consumableId, quantityNote string) (tp.TaskConsumable, error) {
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.UpdateTaskConsumable(tp.TaskConsumable{
		TaskId:       atId,
		ConsumableId: cId,
		QuantityNote: quantityNote,
	})
}

func (a *App) validateTaskConsumable(atc tp.TaskConsumable) error {
	if atc.QuantityNote == "" {
		return ae.ErrTaskConsumableQuantityNoteEmpty
	}

	// we could validate that the task and consumable exist here
	// but we get that for free with the db layer so will exclude here
	// a little leaky but avoids unnecessary db calls
	return nil
}
