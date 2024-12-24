package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateTimeTrigger(assetId string, taskId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
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

func (a *App) UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	return tp.TimeTrigger{}, ae.New(ae.CodeNotImplemented, "UpdateTimeTrigger not implemented")
}

func (a *App) validateTimeTrigger(timeTrigger tp.TimeTrigger) error {
	return ae.New(ae.CodeNotImplemented, "validateTimeTrigger not implemented")
}

func (a *App) timeTriggerExists(id uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "timeTriggerExists not implemented")
}
