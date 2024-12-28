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

	return tp.UsageTrigger{}, ae.New(ae.CodeNotImplemented, "CreateUsageTrigger not implemented")
}

func (a *App) DeleteUsageTrigger(assetId string, taskId string, usageTriggerId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteUsageTrigger not implemented")
}

func (a *App) GetUsageTrigger(assetId string, taskId string, usageTriggerId string) (tp.UsageTrigger, error) {
	return tp.UsageTrigger{}, ae.New(ae.CodeNotImplemented, "GetUsageTrigger not implemented")
}

func (a *App) ListUsageTriggersByAssetAndTask(assetId string, taskId string) ([]tp.UsageTrigger, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListUsageTriggersByAssetAndTask not implemented")
}

func (a *App) ListUsageTriggerUnits() ([]string, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListUsageUnits not implemented")
}

func (a *App) UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	return tp.UsageTrigger{}, ae.New(ae.CodeNotImplemented, "UpdateUsageTrigger not implemented")
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

	te, err := a.taskExists(usageTrigger.TaskId)
	if err != nil {
		return errors.Wrapf(err, "validateUsageTrigger - unexpected error validating task exists")
	}

	if !te {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("validateUsageTrigger - task with id [%s] not found", usageTrigger.TaskId.String()))
	}

	return nil

}

func (a *App) usageTriggerExists(id uuid.UUID) (bool, error) {
	_, err := a.db.GetUsageTrigger(id)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
