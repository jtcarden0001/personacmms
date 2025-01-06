package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func ConvertApiTaskRequestToStoreTask(taskRequest apitp.TaskRequest) (storetp.Task, error) {
	return storetp.Task{}, ae.New(ae.CodeNotImplemented, "ConvertApiTaskRequestToStoreTask not implemented")
}

func ConvertStoreTaskListToApiTaskResponseList(storeTasks []storetp.Task) ([]apitp.TaskResponse, error) {
	return []apitp.TaskResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreTaskListToApiTaskResponseList not implemented")
}

func ConvertStoreTaskToApiTaskResponse(storeTask storetp.Task) (apitp.TaskResponse, error) {
	return apitp.TaskResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreTaskToApiTaskResponse not implemented")
}
