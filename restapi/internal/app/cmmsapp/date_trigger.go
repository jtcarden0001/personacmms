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

	return tp.DateTrigger{}, ae.New(ae.CodeNotImplemented, "CreateDateTrigger not implemented")
}

func (a *App) DeleteDateTrigger(assetId string, taskId string, dateTriggerId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteDateTrigger not implemented")
}

func (a *App) GetDateTrigger(assetId string, taskId string, dateTriggerId string) (tp.DateTrigger, error) {
	return tp.DateTrigger{}, ae.New(ae.CodeNotImplemented, "GetDateTrigger not implemented")
}

func (a *App) ListDateTriggersByAssetAndTask(assetId string, taskId string) ([]tp.DateTrigger, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListDateTriggers not implemented")
}

func (a *App) UpdateDateTrigger(assetId string, taskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	return tp.DateTrigger{}, ae.New(ae.CodeNotImplemented, "UpdateDateTrigger not implemented")
}

func (a *App) validateDateTrigger(dateTrigger tp.DateTrigger) error {
	if dateTrigger.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "dateTrigger id is required")
	}

	if dateTrigger.ScheduledDate.Before(time.Now()) {
		return ae.New(ae.CodeInvalid, "scheduled date must be in the future")
	}

	te, err := a.taskExists(dateTrigger.TaskId)
	if err != nil {
		return errors.Wrapf(err, "validateDateTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("parent task with id [%s] not found", dateTrigger.TaskId.String()))
	}

	return nil
}

func (a *App) dateTriggerExists(dtId uuid.UUID) (bool, error) {
	_, err := a.db.GetDateTrigger(dtId)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
