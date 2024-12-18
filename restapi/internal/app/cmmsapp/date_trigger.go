package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// Create a DateTrigger
func (a *App) CreateDateTrigger(groupTitle string, assetTitle string, taskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	if err := a.validateAndInterpolateDateTrigger(groupTitle, assetTitle, taskId, &dateTrigger); err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.CreateDateTrigger(dateTrigger)
}

// Delete a DateTrigger
func (a *App) DeleteDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string) error {
	// Get before delete to provide opportunity to return not found error
	dt, err := a.GetDateTrigger(groupTitle, assetTitle, taskId, dateTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteDateTrigger(dt.Id)
}

// Get a date trigger that is essentially namespaced under the task specificed
func (a *App) GetDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string) (tp.DateTrigger, error) {
	if _, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId); err != nil {
		return tp.DateTrigger{}, err
	}

	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.GetDateTrigger(parsedDtId)
}

// List all date triggers for a particular task
func (a *App) ListDateTriggers(groupTitle string, assetTitle string, taskId string) ([]tp.DateTrigger, error) {
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.DateTrigger{}, err
	}

	return a.db.ListDateTriggersByTaskId(tid)
}

// Update a date trigger
func (a *App) UpdateDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	if err := a.validateAndInterpolateDateTrigger(groupTitle, assetTitle, taskId, &dateTrigger); err != nil {
		return tp.DateTrigger{}, err
	}

	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, err
	}

	if dateTrigger.Id == uuid.Nil {
		dateTrigger.Id = parsedDtId
	} else if dateTrigger.Id != parsedDtId {
		return tp.DateTrigger{}, ae.ErrIdMismatch
	}

	return a.db.UpdateDateTrigger(parsedDtId, dateTrigger)
}

func (a *App) validateAndInterpolateDateTrigger(groupTitle string, assetTitle string, taskId string, dateTrigger *tp.DateTrigger) error {
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}

	dateTrigger.TaskId = tid
	return nil
}

// common function to validate the dependencies (namespaces) of a trigger and return the final namespace (taskId)
func (a *App) validateTriggerDependencies(groupTitle string, assetTitle string, taskId string) (tp.UUID, error) {
	// Get Asset will validate group and asset
	if _, err := a.GetAsset(groupTitle, assetTitle); err != nil {
		return uuid.Nil, err
	}

	task, err := a.GetTask(groupTitle, assetTitle, taskId)
	return task.Id, err
}
