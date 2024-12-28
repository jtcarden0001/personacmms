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

	aUid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.Task{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}
	task.AssetId = aUid

	err = a.validateTask(task)
	if err != nil {
		return tp.Task{}, errors.Wrapf(err, "CreateTask validation failed")
	}

	return a.db.CreateTask(task)
}

func (a *App) DeleteTask(assetId string, taskId string) error {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	tUid, err := uuid.Parse(taskId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "task id must be a valid uuid")
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

func (a *App) GetTask(assetId string, taskId string) (tp.Task, error) {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return tp.Task{}, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return tp.Task{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	tUid, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, ae.New(ae.CodeInvalid, "task id must be a valid uuid")
	}

	t, err := a.db.GetTask(tUid)
	if err != nil {
		return tp.Task{}, err
	}

	if t.AssetId != aUid {
		return tp.Task{}, ae.New(ae.CodeNotFound, fmt.Sprintf("no task with id [%s] found for asset with id [%s]", taskId, assetId))
	}

	return t, nil
}

func (a *App) ListTasksByAsset(assetId string) ([]tp.Task, error) {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return nil, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	return a.db.ListTasksByAsset(aUid)
}

func (a *App) UpdateTask(assetId string, taskId string, task tp.Task) (tp.Task, error) {
	// check assetId and task existence and coherency (task belongs to asset)
	t, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.Task{}, err
	}

	if task.Id != uuid.Nil && task.Id != t.Id {
		return tp.Task{}, ae.New(ae.CodeInvalid, fmt.Sprintf("task id mismatch [%s] and [%s]", task.Id, t.Id))
	}
	task.Id = t.Id

	if task.AssetId != uuid.Nil && task.AssetId != t.AssetId {
		return tp.Task{}, ae.New(ae.CodeInvalid, fmt.Sprintf("asset id mismatch [%s] and [%s]", task.AssetId, t.AssetId))
	}
	task.AssetId = t.AssetId

	err = a.validateTask(task)
	if err != nil {
		return tp.Task{}, errors.Wrapf(err, "UpdateTask validation failed")
	}

	return a.db.UpdateTask(task)
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

	if _, err := a.db.GetAsset(task.AssetId); err != nil {
		return err
	}

	return nil
}

func (a *App) taskExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
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
