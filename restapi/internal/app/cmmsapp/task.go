package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateTask(assetId string, task tp.Task) (tp.Task, error) {
	if task.Id != uuid.Nil {
		return tp.Task{}, ae.New(ae.CodeInvalid, "task id must be nil on create, we will create an id for you")
	}
	task.Id = uuid.New()

	return tp.Task{}, ae.New(ae.CodeNotImplemented, "CreateTask not implemented")
}

func (a *App) DeleteTask(assetId string, taskId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteTask not implemented")
}

func (a *App) DisassociateTaskWithWorkOrder(assetId string, taskId string, workOrderId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateTaskWithWorkOrder not implemented")
}

func (a *App) GetTask(assetId string, taskId string) (tp.Task, error) {
	return tp.Task{}, ae.New(ae.CodeNotImplemented, "GetTask not implemented")
}

func (a *App) ListTasksByAsset(assetId string) ([]tp.Task, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTasksByAsset not implemented")
}

func (a *App) UpdateTask(assetId string, taskId string, task tp.Task) (tp.Task, error) {
	return tp.Task{}, ae.New(ae.CodeNotImplemented, "UpdateTask not implemented")
}

func (a *App) validateTask(task tp.Task) error {
	if task.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "task id is required")
	}

	if len(task.Title) < tp.MinEntityTitleLength || len(task.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("task title length must be between [%d] and [%d] characters",
			tp.MinEntityTitleLength,
			tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) taskExists(id uuid.UUID) (bool, error) {
	_, err := a.db.GetTask(id)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
