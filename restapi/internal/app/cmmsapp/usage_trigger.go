package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateUsageTrigger(assetId string, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	if usageTrigger.Id != uuid.Nil {
		return tp.UsageTrigger{}, ae.New(ae.CodeInvalid, "usageTrigger id must be nil on create, we will create an id for you")
	}
	usageTrigger.Id = uuid.New()

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "CreateUsageTrigger - error checking task exists")
	}

	if usageTrigger.TaskId != uuid.Nil && usageTrigger.TaskId != task.Id {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, fmt.Sprintf("task id mismatch [%s] does not match [%s]", usageTrigger.TaskId, task.Id))
	}

	usageTrigger.TaskId = task.Id
	err = a.validateUsageTrigger(usageTrigger)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "CreateUsageTrigger validation failed")
	}

	return a.db.CreateUsageTrigger(usageTrigger)
}

func (a *App) DeleteUsageTrigger(assetId string, taskId string, usageTriggerId string) error {
	utUid, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return errors.Wrapf(err, "DeleteUsageTrigger - error checking task exists")
	}

	return a.db.DeleteUsageTriggerFromTask(task.Id, utUid)
}

func (a *App) GetUsageTrigger(assetId string, taskId string, usageTriggerId string) (tp.UsageTrigger, error) {
	utUid, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "GetUsageTrigger - error checking task exists")
	}

	ut, err := a.db.GetUsageTrigger(utUid)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	if ut.TaskId != task.Id {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("no usageTrigger with id [%s] found for task with id [%s]",
				usageTriggerId,
				taskId))
	}

	return ut, nil
}

func (a *App) ListUsageTriggersByAssetAndTask(assetId string, taskId string) ([]tp.UsageTrigger, error) {
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return nil, err
	}

	return a.db.ListUsageTriggersByTask(task.Id)
}

func (a *App) ListUsageTriggerUnits() ([]string, error) {
	keys := make([]string, 0, len(tp.ValidUsageTriggerUnits))
	for k := range tp.ValidUsageTriggerUnits {
		keys = append(keys, k)
	}

	return keys, nil
}

func (a *App) UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	utUid, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	if usageTrigger.Id != uuid.Nil && usageTrigger.Id != utUid {
		return tp.UsageTrigger{}, ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger id mismatch between [%s] and [%s]", usageTriggerId, usageTrigger.Id))
	}
	usageTrigger.Id = utUid

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "UpdateUsageTrigger - error checking task exists")
	}

	if usageTrigger.TaskId != uuid.Nil && usageTrigger.TaskId != task.Id {
		return tp.UsageTrigger{}, ae.New(ae.CodeNotFound, fmt.Sprintf("usageTrigger with id [%s] not found in task with id [%s]", usageTrigger.Id, task.Id))
	}

	usageTrigger.TaskId = task.Id
	err = a.validateUsageTrigger(usageTrigger)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "UpdateUsageTrigger validation failed")
	}

	return a.db.UpdateUsageTrigger(usageTrigger)
}

func (a *App) validateUsageTrigger(usageTrigger tp.UsageTrigger) error {
	if usageTrigger.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "usageTrigger id is required")
	}

	minUsageTriggerQuantity := 1
	if usageTrigger.Quantity < minUsageTriggerQuantity {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger quantity must be greater than [%d]", minUsageTriggerQuantity))
	}

	if !tp.ValidUsageTriggerUnits[usageTrigger.UsageUnit] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger unit must be one of [%s]", tp.PrintValidUsageTriggerUnits()))
	}

	_, te, err := a.taskExists(usageTrigger.TaskId.String())
	if err != nil {
		return errors.Wrapf(err, "validateUsageTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("validateUsageTrigger - task with id [%s] not found", usageTrigger.TaskId.String()))
	}

	return nil

}

func (a *App) usageTriggerExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	_, err = a.db.GetUsageTrigger(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}
