package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateTask(assetId string, task apitp.TaskRequest) (apitp.TaskResponse, error) {
	if task.Id != uuid.Nil {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, "task id must be nil on create, we will create an id for you")
	}
	task.Id = uuid.New()

	aUid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	if task.AssetId != uuid.Nil && task.AssetId != aUid {
		return apitp.TaskResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset id mismatch [%s] does not match [%s]", task.AssetId, assetId))
	}

	task.AssetId = aUid
	err = a.validateTask(task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "CreateTask validation failed")
	}

	stTaskRequest, err := convertApiTaskRequestToStoreTask(task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "CreateTask - error converting to store type")
	}

	stTaskResponse, err := a.db.CreateTask(stTaskRequest)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "CreateTask - error creating task")
	}

	return convertStoreTaskToApiTaskResponse(stTaskResponse)
}

func (a *App) DeleteTask(assetId string, taskId string) error {
	tUid, err := uuid.Parse(taskId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "task id must be a valid uuid")
	}

	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	// TODO: ensure cascading deletions of associated consumables, tools, and triggers.

	return a.db.DeleteTaskFromAsset(aUid, tUid)
}

func (a *App) DisassociateTaskWithWorkOrder(assetId string, taskId string, workOrderId string) error {
	// check assetId and task existence and coherency (task belongs to asset)
	t, err := a.GetTask(assetId, taskId)
	if err != nil {
		return err
	}

	wUid, err := uuid.Parse(workOrderId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "work order id must be a valid uuid")
	}

	// if a workorder is found referencing the task, disassociate it.  If not, notfound error
	return a.db.DisassociateWorkOrderWithTask(t.Id, wUid)
}

func (a *App) GetTask(assetId string, taskId string) (apitp.TaskResponse, error) {
	tUid, err := uuid.Parse(taskId)
	if err != nil {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, "task id must be a valid uuid")
	}

	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return apitp.TaskResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	stTaskResponse, err := a.db.GetTask(tUid)
	if err != nil {
		return apitp.TaskResponse{}, err
	}

	if stTaskResponse.AssetId != aUid {
		return apitp.TaskResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("no task with id [%s] found for asset with id [%s]", taskId, assetId))
	}

	return convertStoreTaskToApiTaskResponse(stTaskResponse)
}

func (a *App) ListTasksByAsset(assetId string) ([]apitp.TaskResponse, error) {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return nil, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	stAssetResponses, err := a.db.ListTasksByAsset(aUid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListTasksByAsset failed")
	}

	return convertStoreTaskListToApiTaskResponseList(stAssetResponses)
}

func (a *App) UpdateTask(assetId string, taskId string, task apitp.TaskRequest) (apitp.TaskResponse, error) {
	// check assetId and task existence and coherency (task belongs to asset)
	t, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.TaskResponse{}, err
	}

	if task.Id != uuid.Nil && task.Id != t.Id {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("task id mismatch [%s] and [%s]", task.Id, t.Id))
	}
	task.Id = t.Id

	if task.AssetId != uuid.Nil && task.AssetId != t.AssetId {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("asset id mismatch [%s] and [%s]", task.AssetId, t.AssetId))
	}
	task.AssetId = t.AssetId

	err = a.validateTask(task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask validation failed")
	}

	stTaskRequest, err := convertApiTaskRequestToStoreTask(task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask - error converting to store type")
	}

	stTaskResponse, err := a.db.UpdateTask(stTaskRequest)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask - error updating task")
	}

	return convertStoreTaskToApiTaskResponse(stTaskResponse)
}

func (a *App) validateTask(task apitp.TaskRequest) error {
	if task.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "task id is required")
	}

	if len(task.Title) < apitp.MinEntityTitleLength || len(task.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("task title length must be between [%d] and [%d] characters",
			apitp.MinEntityTitleLength,
			apitp.MaxEntityTitleLength))
	}

	_, aFound, err := a.assetExists(task.AssetId.String())
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", task.AssetId))
	}

	return nil
}

func (a *App) taskExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil || uid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "task id must be a valid uuid")
	}
	_, err = a.db.GetTask(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}

func convertApiTaskRequestToStoreTask(taskRequest apitp.TaskRequest) (storetp.Task, error) {
	return storetp.Task{}, ae.New(ae.CodeNotImplemented, "ConvertApiTaskRequestToStoreTask not implemented")
}

func convertStoreTaskListToApiTaskResponseList(storeTasks []storetp.Task) ([]apitp.TaskResponse, error) {
	return []apitp.TaskResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreTaskListToApiTaskResponseList not implemented")
}

func convertStoreTaskToApiTaskResponse(storeTask storetp.Task) (apitp.TaskResponse, error) {
	return apitp.TaskResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreTaskToApiTaskResponse not implemented")
}
