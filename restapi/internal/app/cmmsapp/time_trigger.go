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

	return tp.TimeTrigger{}, ae.New(ae.CodeNotImplemented, "CreateTimeTrigger not implemented")
}

func (a *App) DeleteTimeTrigger(assetId string, taskId string, timeTriggerId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteTimeTrigger not implemented")
}

func (a *App) GetTimeTrigger(assetId string, taskId string, timeTriggerId string) (tp.TimeTrigger, error) {
	return tp.TimeTrigger{}, ae.New(ae.CodeNotImplemented, "GetTimeTrigger not implemented")
}

func (a *App) ListTimeTriggersByAssetAndTask(assetId string, taskId string) ([]tp.TimeTrigger, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTimeTriggersByAssetAndTask not implemented")
}

func (a *App) ListTimeTriggerUnits() ([]string, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTimeUnits not implemented")
}

func (a *App) UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	return tp.TimeTrigger{}, ae.New(ae.CodeNotImplemented, "UpdateTimeTrigger not implemented")
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

	te, err := a.taskExists(timeTrigger.TaskId)
	if err != nil {
		return errors.Wrapf(err, "validateTimeTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("parent task with id [%s] not found", timeTrigger.TaskId.String()))
	}

	return nil
}

func (a *App) timeTriggerExists(id uuid.UUID) (bool, error) {
	_, err := a.db.GetTimeTrigger(id)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
