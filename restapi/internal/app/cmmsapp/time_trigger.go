package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateTimeTrigger(assetId string, taskId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	if timeTrigger.Id != uuid.Nil {
		return tp.TimeTrigger{}, ae.New(ae.CodeInvalid, "timeTrigger id must be nil on create, we will create an id for you")
	}
	timeTrigger.Id = uuid.New()

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "CreateTimeTrigger - error checking task exists")
	}

	if timeTrigger.TaskId != uuid.Nil && timeTrigger.TaskId != task.Id {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound, fmt.Sprintf("task id mismatch [%s] does not match [%s]", timeTrigger.TaskId, task.Id))
	}

	timeTrigger.TaskId = task.Id
	err = a.validateTimeTrigger(timeTrigger)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "CreateTimeTrigger validation failed")
	}

	return a.db.CreateTimeTrigger(timeTrigger)
}

func (a *App) DeleteTimeTrigger(assetId string, taskId string, timeTriggerId string) error {
	ttUid, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return errors.Wrapf(err, "DeleteTimeTrigger - error checking task exists")
	}

	return a.db.DeleteTimeTriggerFromTask(task.Id, ttUid)
}

func (a *App) GetTimeTrigger(assetId string, taskId string, timeTriggerId string) (tp.TimeTrigger, error) {
	ttUid, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "GetTimeTrigger - error checking task exists")
	}

	tt, err := a.db.GetTimeTrigger(ttUid)
	if err != nil {
		return tp.TimeTrigger{}, err
	}

	if tt.TaskId != task.Id {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("timeTrigger with id [%s] not found in task with id [%s]",
				tt.Id,
				task.Id))
	}

	return tt, nil

}

func (a *App) ListTimeTriggersByAssetAndTask(assetId string, taskId string) ([]tp.TimeTrigger, error) {
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListTimeTriggersByAssetAndTask - error checking task exists")
	}

	return a.db.ListTimeTriggersByTask(task.Id)
}

func (a *App) ListTimeTriggerUnits() ([]string, error) {
	keys := make([]string, 0, len(tp.ValidTimeTriggerUnits))
	for key := range tp.ValidTimeTriggerUnits {
		keys = append(keys, key)
	}

	return keys, nil
}

func (a *App) UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	ttUid, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	if timeTrigger.Id != uuid.Nil && timeTrigger.Id != ttUid {
		return tp.TimeTrigger{}, ae.New(ae.CodeInvalid,
			fmt.Sprintf("timeTrigger id mismatch between [%s] and [%s]",
				timeTriggerId, timeTrigger.Id))
	}
	timeTrigger.Id = ttUid

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "UpdateTimeTrigger - error checking task exists")
	}

	if timeTrigger.TaskId != uuid.Nil && timeTrigger.TaskId != task.Id {
		return tp.TimeTrigger{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("task id mismatch [%s] does not match [%s]",
				timeTrigger.TaskId, task.Id))
	}

	timeTrigger.TaskId = task.Id
	err = a.validateTimeTrigger(timeTrigger)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "UpdateTimeTrigger validation failed")
	}

	return a.db.UpdateTimeTrigger(timeTrigger)
}

func (a *App) validateTimeTrigger(timeTrigger tp.TimeTrigger) error {
	if timeTrigger.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "timeTrigger id is required")
	}

	minTimeTriggerQuantity := 1
	if timeTrigger.Quantity < minTimeTriggerQuantity {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("timeTrigger quantity must be greater or equal to [%d]", minTimeTriggerQuantity))
	}

	if !tp.ValidTimeTriggerUnits[timeTrigger.TimeUnit] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("timeTrigger unit must be one of [%s]", tp.PrintValidTimeTriggerUnits()))
	}

	_, te, err := a.taskExists(timeTrigger.TaskId.String())
	if err != nil {
		return errors.Wrapf(err, "validateTimeTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("parent task with id [%s] not found", timeTrigger.TaskId.String()))
	}

	return nil
}

func (a *App) timeTriggerExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	_, err = a.db.GetTimeTrigger(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}
