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
	aUid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.TaskResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	err = a.validateTaskAndAsset(aUid, task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "CreateTask validation failed")
	}

	newTaskId := uuid.New()
	stTaskRequest, err := convertApiTaskRequestToStoreTask(aUid, newTaskId, task)
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

	// TODO: inefficiency here, we are validating asset above with GetTask and again in validateTaskAndAsset
	err = a.validateTaskAndAsset(t.AssetId, task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask validation failed")
	}

	stTaskRequest, err := convertApiTaskRequestToStoreTask(t.AssetId, t.Id, task)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask - error converting to store type")
	}

	stTaskResponse, err := a.db.UpdateTask(stTaskRequest)
	if err != nil {
		return apitp.TaskResponse{}, errors.Wrapf(err, "UpdateTask - error updating task")
	}

	return convertStoreTaskToApiTaskResponse(stTaskResponse)
}

func (a *App) validateTaskAndAsset(assetId uuid.UUID, task apitp.TaskRequest) error {
	if len(task.Title) < apitp.MinEntityTitleLength || len(task.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("task title length must be between [%d] and [%d] characters",
			apitp.MinEntityTitleLength,
			apitp.MaxEntityTitleLength))
	}

	_, aFound, err := a.assetExists(assetId.String())
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
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

func convertApiTaskRequestToStoreTask(assetId uuid.UUID, taskId uuid.UUID, taskRequest apitp.TaskRequest) (storetp.Task, error) {
	stTask := storetp.Task{
		Id:           taskId,
		Title:        taskRequest.Title,
		Instructions: taskRequest.Instructions,
		AssetId:      assetId,
	}

	return stTask, nil
}

func convertStoreTaskListToApiTaskResponseList(storeTasks []storetp.Task) ([]apitp.TaskResponse, error) {
	apiTasks := make([]apitp.TaskResponse, len(storeTasks))
	for i, stTask := range storeTasks {
		apiTask, err := convertStoreTaskToApiTaskResponse(stTask)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting store task to api task")
		}
		apiTasks[i] = apiTask
	}
	return apiTasks, nil
}

func convertStoreTaskToApiTaskResponse(storeTask storetp.Task) (apitp.TaskResponse, error) {
	apiTask := apitp.TaskResponse{
		Id:             storeTask.Id,
		Title:          storeTask.Title,
		Instructions:   storeTask.Instructions,
		AssetId:        storeTask.AssetId,
		AssetReference: fmt.Sprintf("/api/v1/assets/%s", storeTask.AssetId),
	}

	return apiTask, nil
}
