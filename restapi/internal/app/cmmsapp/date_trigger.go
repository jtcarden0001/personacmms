package cmmsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateDateTrigger(assetId string, taskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	if dateTrigger.Id != uuid.Nil {
		return tp.DateTrigger{}, ae.New(ae.CodeInvalid, "dateTrigger id must be nil on create, we will create an id for you")
	}
	dateTrigger.Id = uuid.New()

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "CreateDateTrigger - error checking task exists")
	}

	if dateTrigger.TaskId != uuid.Nil && dateTrigger.TaskId != task.Id {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("task id mismatch [%s] does not match [%s]",
				dateTrigger.TaskId, task.Id))
	}

	dateTrigger.TaskId = task.Id
	err = a.validateDateTrigger(dateTrigger)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "CreateDateTrigger validation failed")
	}

	return a.db.CreateDateTrigger(dateTrigger)
}

func (a *App) DeleteDateTrigger(assetId string, taskId string, dateTriggerId string) error {
	dtUid, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return errors.Wrapf(err, "DeleteDateTrigger - error checking task exists")
	}

	return a.db.DeleteDateTriggerFromTask(task.Id, dtUid)
}

func (a *App) GetDateTrigger(assetId string, taskId string, dateTriggerId string) (tp.DateTrigger, error) {
	dateTriggerUuid, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "DeleteDateTrigger - error checking task exists")
	}

	dateTrigger, err := a.db.GetDateTrigger(dateTriggerUuid)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "DeleteDateTrigger - error checking dateTrigger exists")
	}

	if dateTrigger.TaskId != task.Id {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("dateTrigger with id [%s] not found in task with id [%s]",
				dateTriggerId,
				taskId))
	}

	return dateTrigger, nil
}

func (a *App) ListDateTriggersByAssetAndTask(assetId string, taskId string) ([]tp.DateTrigger, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListDateTriggersByAssetAndTask - error checking task exists")
	}

	return a.db.ListDateTriggersByTask(task.Id)
}

func (a *App) UpdateDateTrigger(assetId string, taskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	dateTriggerUuid, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	if dateTrigger.Id != uuid.Nil && dateTrigger.Id != dateTriggerUuid {
		return tp.DateTrigger{}, ae.New(ae.CodeInvalid, fmt.Sprintf("dateTrigger id mismatch between [%s] and [%s]", dateTriggerId, dateTrigger.Id))
	}
	dateTrigger.Id = dateTriggerUuid

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "UpdateDateTrigger - error checking task exists")
	}

	if dateTrigger.TaskId != uuid.Nil && dateTrigger.TaskId != task.Id {
		return tp.DateTrigger{}, ae.New(ae.CodeNotFound, fmt.Sprintf("dateTrigger with id [%s] not found in task with id [%s]", dateTrigger.Id, task.Id))
	}

	dateTrigger.TaskId = task.Id
	err = a.validateDateTrigger(dateTrigger)
	if err != nil {
		return tp.DateTrigger{}, errors.Wrapf(err, "UpdateDateTrigger validation failed")
	}

	return a.db.UpdateDateTrigger(dateTrigger)
}

func (a *App) validateDateTrigger(dateTrigger tp.DateTrigger) error {
	if dateTrigger.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "dateTrigger id is required")
	}

	if dateTrigger.ScheduledDate.Before(time.Now()) {
		return ae.New(ae.CodeInvalid, "scheduled date must be in the future")
	}

	_, te, err := a.taskExists(dateTrigger.TaskId.String())
	if err != nil {
		return errors.Wrapf(err, "validateDateTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("parent task with id [%s] not found", dateTrigger.TaskId.String()))
	}

	return nil
}

func (a *App) dateTriggerExists(dtId string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(dtId)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	_, err = a.db.GetDateTrigger(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}
