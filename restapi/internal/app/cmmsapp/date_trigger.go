package cmmsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateDateTrigger(assetId string, taskId string, dateTrigger apitp.DateTriggerRequest) (apitp.DateTriggerResponse, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "CreateDateTrigger - error checking task exists")
	}

	err = a.validateDateTrigger(dateTrigger)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "CreateDateTrigger validation failed")
	}

	dtId := uuid.New()
	stDtRequest, err := convertApiDateTriggerRequestToStoreDateTrigger(dtId, task.Id, dateTrigger)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "CreateDateTrigger - error converting to store type")
	}

	stDtResponse, err := a.db.CreateDateTrigger(stDtRequest)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "CreateDateTrigger - error creating dateTrigger")
	}

	return convertStoreDateTriggerToApiDateTriggerResponse(stDtResponse)
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

func (a *App) GetDateTrigger(assetId string, taskId string, dateTriggerId string) (apitp.DateTriggerResponse, error) {
	dateTriggerUuid, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return apitp.DateTriggerResponse{}, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "DeleteDateTrigger - error checking task exists")
	}

	dateTrigger, err := a.db.GetDateTrigger(dateTriggerUuid)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "DeleteDateTrigger - error checking dateTrigger exists")
	}

	if dateTrigger.TaskId != task.Id {
		return apitp.DateTriggerResponse{}, ae.New(ae.CodeNotFound,
			fmt.Sprintf("dateTrigger with id [%s] not found in task with id [%s]",
				dateTriggerId,
				taskId))
	}

	return convertStoreDateTriggerToApiDateTriggerResponse(dateTrigger)
}

func (a *App) ListDateTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.DateTriggerResponse, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListDateTriggersByAssetAndTask - error checking task exists")
	}

	stDtReponses, err := a.db.ListDateTriggersByTask(task.Id)
	if err != nil {
		return []apitp.DateTriggerResponse{}, errors.Wrapf(err, "ListDateTriggersByAssetAndTask - error listing dateTriggers")
	}

	return convertStoreDateTriggerListToApiDateTriggerResponseList(stDtReponses)
}

func (a *App) UpdateDateTrigger(assetId string, taskId string, dateTriggerId string, dateTrigger apitp.DateTriggerRequest) (apitp.DateTriggerResponse, error) {
	// check namespace coherency
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "UpdateDateTrigger - error checking task exists")
	}

	dateTriggerUuid, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return apitp.DateTriggerResponse{}, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid uuid")
	}

	err = a.validateDateTrigger(dateTrigger)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "UpdateDateTrigger validation failed")
	}

	stDtRequest, err := convertApiDateTriggerRequestToStoreDateTrigger(dateTriggerUuid, task.Id, dateTrigger)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "UpdateDateTrigger - error converting to store type")
	}

	stDtReponse, err := a.db.UpdateDateTrigger(stDtRequest)
	if err != nil {
		return apitp.DateTriggerResponse{}, errors.Wrapf(err, "UpdateDateTrigger - error updating dateTrigger")
	}

	return convertStoreDateTriggerToApiDateTriggerResponse(stDtReponse)
}

func (a *App) validateDateTrigger(dateTrigger apitp.DateTriggerRequest) error {
	if dateTrigger.ScheduledDate.Before(time.Now()) {
		return ae.New(ae.CodeInvalid, "scheduled date must be in the future")
	}

	return nil
}

func (a *App) dateTriggerExists(dtId string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(dtId)
	if err != nil || uid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "dateTrigger id must be a valid, non nil uuid")
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

func convertApiDateTriggerRequestToStoreDateTrigger(id uuid.UUID, taskId uuid.UUID, dateTriggerRequest apitp.DateTriggerRequest) (storetp.DateTrigger, error) {
	stDt := storetp.DateTrigger{
		Id:            id,
		ScheduledDate: dateTriggerRequest.ScheduledDate,
		TaskId:        taskId,
	}

	return stDt, nil
}

func convertStoreDateTriggerListToApiDateTriggerResponseList(storeDateTriggers []storetp.DateTrigger) ([]apitp.DateTriggerResponse, error) {
	apiDts := make([]apitp.DateTriggerResponse, len(storeDateTriggers))
	for i, dt := range storeDateTriggers {
		apiDt, err := convertStoreDateTriggerToApiDateTriggerResponse(dt)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting store date trigger to api date trigger")
		}
		apiDts[i] = apiDt
	}

	return apiDts, nil
}

func convertStoreDateTriggerToApiDateTriggerResponse(storeDateTrigger storetp.DateTrigger) (apitp.DateTriggerResponse, error) {
	apiDt := apitp.DateTriggerResponse{
		Id:            storeDateTrigger.Id,
		ScheduledDate: storeDateTrigger.ScheduledDate,
		TaskId:        storeDateTrigger.TaskId,
	}

	return apiDt, nil
}
