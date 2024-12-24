package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateUsageTrigger(assetId string, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
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

func (a *App) UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	return tp.UsageTrigger{}, ae.New(ae.CodeNotImplemented, "UpdateUsageTrigger not implemented")
}

func (a *App) validateUsageTrigger(usageTrigger tp.UsageTrigger) error {
	return ae.New(ae.CodeNotImplemented, "validateUsageTrigger not implemented")
}

func (a *App) usageTriggerExists(id uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "usageTriggerExists not implemented")
}
