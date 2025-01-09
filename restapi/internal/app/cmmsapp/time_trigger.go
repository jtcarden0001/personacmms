package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateTimeTrigger(assetId string, taskId string, timeTrigger apitp.TimeTriggerRequest) (apitp.TimeTriggerResponse, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "CreateTimeTrigger - error checking task exists")
	}

	err = a.validateTimeTrigger(timeTrigger)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "CreateTimeTrigger validation failed")
	}

	newTimeTriggerId := uuid.New()
	stTTRequest, err := convertApiTimeTriggerRequestToStoreTimeTrigger(task, newTimeTriggerId, timeTrigger)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "CreateTimeTrigger - error converting to store type")
	}

	stTTResponse, err := a.db.CreateTimeTrigger(stTTRequest)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "CreateTimeTrigger - error creating timeTrigger")
	}

	return convertStoreTimeTriggerToApiTimeTriggerResponse(stTTResponse)
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

func (a *App) GetTimeTrigger(assetId string, taskId string, timeTriggerId string) (apitp.TimeTriggerResponse, error) {
	ttUid, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return apitp.TimeTriggerResponse{}, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "GetTimeTrigger - error checking task exists")
	}

	tt, err := a.db.GetTimeTrigger(ttUid)
	if err != nil {
		return apitp.TimeTriggerResponse{}, err
	}

	if tt.Task.Id != task.Id {
		return apitp.TimeTriggerResponse{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("timeTrigger with id [%s] not found in task with id [%s]",
				tt.Id,
				task.Id))
	}

	return convertStoreTimeTriggerToApiTimeTriggerResponse(tt)

}

func (a *App) ListTimeTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.TimeTriggerResponse, error) {
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListTimeTriggersByAssetAndTask - error checking task exists")
	}

	stTTResponses, err := a.db.ListTimeTriggersByTask(task.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "ListTimeTriggersByAssetAndTask failed")
	}

	return convertStoreTimeTriggerListToApiTimeTriggerResponseList(stTTResponses)
}

func (a *App) ListTimeTriggerUnits() ([]string, error) {
	keys := make([]string, 0, len(apitp.ValidTimeTriggerUnits))
	for key := range apitp.ValidTimeTriggerUnits {
		keys = append(keys, key)
	}

	return keys, nil
}

func (a *App) UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger apitp.TimeTriggerRequest) (apitp.TimeTriggerResponse, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "UpdateTimeTrigger - error checking task exists")
	}

	ttUid, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return apitp.TimeTriggerResponse{}, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid uuid")
	}

	err = a.validateTimeTrigger(timeTrigger)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "UpdateTimeTrigger validation failed")
	}

	stTTRequest, err := convertApiTimeTriggerRequestToStoreTimeTrigger(task, ttUid, timeTrigger)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "UpdateTimeTrigger - error converting to store type")
	}

	stTTResponse, err := a.db.UpdateTimeTrigger(stTTRequest)
	if err != nil {
		return apitp.TimeTriggerResponse{}, errors.Wrapf(err, "UpdateTimeTrigger - error updating timeTrigger")
	}

	return convertStoreTimeTriggerToApiTimeTriggerResponse(stTTResponse)
}

func (a *App) validateTimeTrigger(timeTrigger apitp.TimeTriggerRequest) error {
	minTimeTriggerQuantity := 1
	if timeTrigger.Quantity < minTimeTriggerQuantity {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("timeTrigger quantity must be greater or equal to [%d]", minTimeTriggerQuantity))
	}

	if !apitp.ValidTimeTriggerUnits[timeTrigger.TimeUnit] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("timeTrigger unit must be one of [%s]", apitp.PrintValidTimeTriggerUnits()))
	}

	return nil
}

func (a *App) timeTriggerExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil || uid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "timeTrigger id must be a valid, non-nil uuid")
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

func convertApiTimeTriggerRequestToStoreTimeTrigger(task apitp.TaskResponse, ttId uuid.UUID, timeTrigger apitp.TimeTriggerRequest) (storetp.TimeTrigger, error) {
	stTT := storetp.TimeTrigger{
		Id:       ttId,
		Quantity: timeTrigger.Quantity,
		TimeUnit: timeTrigger.TimeUnit,
		Task: storetp.Task{
			Id:      task.Id,
			AssetId: task.AssetId,
		},
	}

	return stTT, nil

}

func convertStoreTimeTriggerListToApiTimeTriggerResponseList(timeTriggers []storetp.TimeTrigger) ([]apitp.TimeTriggerResponse, error) {
	apiTTs := []apitp.TimeTriggerResponse{}
	for _, tt := range timeTriggers {
		apiTT, err := convertStoreTimeTriggerToApiTimeTriggerResponse(tt)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting store time trigger to api time trigger")
		}
		apiTTs = append(apiTTs, apiTT)
	}

	return apiTTs, nil
}

func convertStoreTimeTriggerToApiTimeTriggerResponse(timeTrigger storetp.TimeTrigger) (apitp.TimeTriggerResponse, error) {
	apiTT := apitp.TimeTriggerResponse{
		Id:            timeTrigger.Id,
		Quantity:      timeTrigger.Quantity,
		TimeUnit:      timeTrigger.TimeUnit,
		TaskId:        timeTrigger.Task.Id,
		TaskReference: fmt.Sprintf("/api/v1/assets/%s/tasks/%s", timeTrigger.Task.AssetId, timeTrigger.Task.Id),
	}

	return apiTT, nil
}
