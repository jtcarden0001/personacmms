package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateUsageTrigger(assetId string, taskId string, usageTrigger apitp.UsageTriggerRequest) (apitp.UsageTriggerResponse, error) {
	if usageTrigger.Id != uuid.Nil {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeInvalid, "usageTrigger id must be nil on create, we will create an id for you")
	}
	usageTrigger.Id = uuid.New()

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "CreateUsageTrigger - error checking task exists")
	}

	if usageTrigger.TaskId != uuid.Nil && usageTrigger.TaskId != task.Id {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("task id mismatch [%s] does not match [%s]", usageTrigger.TaskId, task.Id))
	}

	usageTrigger.TaskId = task.Id
	err = a.validateUsageTrigger(usageTrigger)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "CreateUsageTrigger validation failed")
	}

	stUTRequest, err := convertApiUsageTriggerRequestToStoreUsageTrigger(usageTrigger)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "CreateUsageTrigger - error converting to store type")
	}

	stUTResponse, err := a.db.CreateUsageTrigger(stUTRequest)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "CreateUsageTrigger - error creating usageTrigger")
	}

	return convertStoreUsageTriggerToApiUsageTriggerResponse(stUTResponse)
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

func (a *App) GetUsageTrigger(assetId string, taskId string, usageTriggerId string) (apitp.UsageTriggerResponse, error) {
	utUid, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "GetUsageTrigger - error checking task exists")
	}

	ut, err := a.db.GetUsageTrigger(utUid)
	if err != nil {
		return apitp.UsageTriggerResponse{}, err
	}

	if ut.TaskId != task.Id {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("no usageTrigger with id [%s] found for task with id [%s]",
				usageTriggerId,
				taskId))
	}

	return convertStoreUsageTriggerToApiUsageTriggerResponse(ut)
}

func (a *App) ListUsageTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.UsageTriggerResponse, error) {
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return []apitp.UsageTriggerResponse{}, err
	}

	stUTResponses, err := a.db.ListUsageTriggersByTask(task.Id)
	if err != nil {
		return []apitp.UsageTriggerResponse{}, errors.Wrapf(err, "ListUsageTriggersByAssetAndTask failed")
	}

	return convertStoreUsageTriggerListToApiUsageTriggerResponseList(stUTResponses)
}

func (a *App) ListUsageTriggerUnits() ([]string, error) {
	keys := make([]string, 0, len(apitp.ValidUsageTriggerUnits))
	for k := range apitp.ValidUsageTriggerUnits {
		keys = append(keys, k)
	}

	return keys, nil
}

func (a *App) UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger apitp.UsageTriggerRequest) (apitp.UsageTriggerResponse, error) {
	utUid, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid uuid")
	}

	if usageTrigger.Id != uuid.Nil && usageTrigger.Id != utUid {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger id mismatch between [%s] and [%s]", usageTriggerId, usageTrigger.Id))
	}
	usageTrigger.Id = utUid

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "UpdateUsageTrigger - error checking task exists")
	}

	if usageTrigger.TaskId != uuid.Nil && usageTrigger.TaskId != task.Id {
		return apitp.UsageTriggerResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("usageTrigger with id [%s] not found in task with id [%s]", usageTrigger.Id, task.Id))
	}

	usageTrigger.TaskId = task.Id
	err = a.validateUsageTrigger(usageTrigger)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "UpdateUsageTrigger validation failed")
	}

	stUTRequest, err := convertApiUsageTriggerRequestToStoreUsageTrigger(usageTrigger)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "UpdateUsageTrigger - error converting to store type")
	}

	stUTResponse, err := a.db.UpdateUsageTrigger(stUTRequest)
	if err != nil {
		return apitp.UsageTriggerResponse{}, errors.Wrapf(err, "UpdateUsageTrigger - error updating usageTrigger")
	}

	return convertStoreUsageTriggerToApiUsageTriggerResponse(stUTResponse)
}

func (a *App) validateUsageTrigger(usageTrigger apitp.UsageTriggerRequest) error {
	if usageTrigger.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "usageTrigger id is required")
	}

	minUsageTriggerQuantity := 1
	if usageTrigger.Quantity < minUsageTriggerQuantity {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger quantity must be greater than [%d]", minUsageTriggerQuantity))
	}

	if !apitp.ValidUsageTriggerUnits[usageTrigger.UsageUnit] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("usageTrigger unit must be one of [%s]", apitp.PrintValidUsageTriggerUnits()))
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
	if err != nil || uid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "usageTrigger id must be a valid, non-nil uuid")
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

func convertApiUsageTriggerRequestToStoreUsageTrigger(ut apitp.UsageTriggerRequest) (storetp.UsageTrigger, error) {
	return storetp.UsageTrigger{}, ae.New(ae.CodeNotImplemented, "convertApiUsageTriggerRequestToStoreUsageTrigger not implemented")
}

func convertStoreUsageTriggerListToApiUsageTriggerResponseList(uts []storetp.UsageTrigger) ([]apitp.UsageTriggerResponse, error) {
	return []apitp.UsageTriggerResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreUsageTriggerListToApiUsageTriggerResponseList not implemented")
}

func convertStoreUsageTriggerToApiUsageTriggerResponse(ut storetp.UsageTrigger) (apitp.UsageTriggerResponse, error) {
	return apitp.UsageTriggerResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreUsageTriggerToApiUsageTriggerResponse not implemented")
}
