package cmmsapp

import (
	"github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateDateTrigger(groupTitle string, assetTitle string, taskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	if err := a.validateAndInterpolateTrigger(groupTitle, assetTitle, taskId, &dateTrigger); err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.CreateDateTrigger(dateTrigger)
}

func (a *App) DeleteDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string) error {
	if _, err := a.validateTriggerAndGetTaskId(groupTitle, assetTitle, taskId); err != nil {
		return err
	}

	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteDateTrigger(parsedDtId)
}

func (a *App) ListDateTriggers(groupTitle string, assetTitle string, taskId string) ([]tp.DateTrigger, error) {
	tid, err := a.validateTriggerAndGetTaskId(groupTitle, assetTitle, taskId)
	if err != nil {
		return nil, err
	}

	return a.db.ListDateTriggersByTaskId(tid)
}

func (a *App) GetDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string) (tp.DateTrigger, error) {
	if _, err := a.validateTriggerAndGetTaskId(groupTitle, assetTitle, taskId); err != nil {
		return tp.DateTrigger{}, err
	}

	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.GetDateTrigger(parsedDtId)
}

func (a *App) UpdateDateTrigger(groupTitle string, assetTitle string, taskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	if err := a.validateAndInterpolateTrigger(groupTitle, assetTitle, taskId, &dateTrigger); err != nil {
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

func (a *App) validateAndInterpolateTrigger(groupTitle string, assetTitle string, taskId string, dateTrigger *tp.DateTrigger) error {
	tid, err := a.validateTriggerAndGetTaskId(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}

	dateTrigger.TaskId = tid
	return nil
}

func (a *App) validateTriggerAndGetTaskId(groupTitle string, assetTitle string, taskId string) (tp.UUID, error) {
	if _, err := a.GetGroup(groupTitle); err != nil {
		return uuid.Nil, err
	}

	if _, err := a.GetAsset(groupTitle, assetTitle); err != nil {
		return uuid.Nil, err
	}

	task, err := a.GetTask(groupTitle, assetTitle, taskId)
	return task.Id, err
}
