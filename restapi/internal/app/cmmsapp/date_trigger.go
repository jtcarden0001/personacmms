package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateDateTrigger(assetId string, taskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
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
	return ae.New(ae.CodeNotImplemented, "validateDateTrigger not implemented")
}

func (a *App) dateTriggerExists(dtId uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "dateTriggerExists not implemented")
}
