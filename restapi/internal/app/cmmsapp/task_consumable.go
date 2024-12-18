package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// create task consumable aka map a consumable to a task and quantify how much
func (a *App) CreateTaskConsumable(atc tp.TaskConsumable) (tp.TaskConsumable, error) {
	if err := a.validateTaskConsumable(atc); err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.CreateTaskConsumable(atc)
}

// Create a task consumable with namespace validation for the associated task
func (a *App) CreateTaskConsumableWithValidation(groupTitle, assetTitle, taskId, consumableId, quantityNote string) (tp.TaskConsumable, error) {
	// validate group, asset, and task coherency
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	// construct the target tc
	tc := tp.TaskConsumable{
		TaskId:       t.Id,
		ConsumableId: cId,
		QuantityNote: quantityNote,
	}

	err = a.validateTaskConsumable(tc)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.CreateTaskConsumable(tc)
}

// delete a task consumable
func (a *App) DeleteTaskConsumable(groupTitle, assetTitle, taskId, consumableId string) error {
	// validate group, asset, and task coherency
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "consumable id invalid")
	}

	return a.db.DeleteTaskConsumable(t.Id, cId)
}

// list task consumables for a task
func (a *App) ListTaskConsumables(groupTitle, assetTitle, taskId string) ([]tp.TaskConsumable, error) {
	// validate group, asset, and task coherency
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.TaskConsumable{}, err
	}

	tcs, err := a.db.ListTaskConsumablesByTaskId(t.Id)
	if err != nil {
		return []tp.TaskConsumable{}, err
	}

	return tcs, nil
}

// get a task consumable
func (a *App) GetTaskConsumable(groupTitle, assetTitle, taskId, consumableId string) (tp.TaskConsumable, error) {
	// validate group, asset, and task coherency
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, ae.New(ae.CodeInvalid, "consumable id invalid")
	}

	return a.db.GetTaskConsumable(t.Id, cId)
}

// update a task consumable
func (a *App) UpdateTaskConsumable(tc tp.TaskConsumable) (tp.TaskConsumable, error) {
	err := a.validateTaskConsumable(tc)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.UpdateTaskConsumable(tc)
}

// update a task consumable with namespace validation for the associated task
func (a *App) UpdateTaskConsumableWithValidation(groupTitle, assetTitle, taskId, consumableId, quantityNote string) (tp.TaskConsumable, error) {
	// validate group, asset, and task coherency
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	cId, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	// construct the target tc
	tc := tp.TaskConsumable{
		TaskId:       t.Id,
		ConsumableId: cId,
		QuantityNote: quantityNote,
	}

	err = a.validateTaskConsumable(tc)
	if err != nil {
		return tp.TaskConsumable{}, err
	}

	return a.db.UpdateTaskConsumable(tc)
}

func (a *App) validateTaskConsumable(atc tp.TaskConsumable) error {
	// validate task exists
	_, err := a.db.GetTask(atc.TaskId)
	if err != nil {
		return err
	}

	// validate consumable exists
	_, err = a.db.GetConsumableById(atc.ConsumableId)
	if err != nil {
		return err
	}

	if atc.QuantityNote == "" {
		return ae.ErrTaskConsumableQuantityNoteEmpty
	}

	return nil
}
