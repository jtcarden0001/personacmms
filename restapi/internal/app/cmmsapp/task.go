package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// CreateTask creates a new task for an asset
func (a *App) CreateTask(groupTitle string, assetTitle string, task tp.Task) (tp.Task, error) {
	if err := a.validateAndInterpolateTask(groupTitle, assetTitle, &task); err != nil {
		return tp.Task{}, err
	}

	return a.db.CreateTask(task)
}

// DeleteTask deletes a task for an asset
func (a *App) DeleteTask(groupTitle string, assetTitle string, taskId string) error {
	task, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}

	return a.db.DeleteTask(task.Id)
}

// GetTask retrieves a task for an asset
func (a *App) GetTask(groupTitle string, assetTitle string, taskId string) (tp.Task, error) {
	// validate the asset is in the group
	asset, err := a.GetAsset(groupTitle, assetTitle)
	if err != nil {
		return tp.Task{}, err
	}

	// validate valid id format
	tId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, ae.New(ae.CodeInvalid, "task id invalid")
	}

	return a.db.GetTaskByAssetId(asset.Id, tId)
}

// list tasks for an asset
func (a *App) ListTasks(groupTitle string, assetTitle string) ([]tp.Task, error) {
	asset, err := a.GetAsset(groupTitle, assetTitle)
	if err != nil {
		return []tp.Task{}, err
	}

	tasks, err := a.db.ListTasksByAssetId(asset.Id)
	if err != nil {
		return []tp.Task{}, err
	}

	return tasks, nil
}

// UpdateTask updates a task for an asset
func (a *App) UpdateTask(groupTitle string, assetTitle string, taskId string, at tp.Task) (tp.Task, error) {
	err := a.validateAndInterpolateTask(groupTitle, assetTitle, &at)
	if err != nil {
		return tp.Task{}, err
	}

	tId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, ae.New(ae.CodeInvalid, "task id invalid")
	}

	return a.db.UpdateTask(tId, at)
}

func (a *App) validateAndInterpolateTask(groupTitle, assetTitle string, task *tp.Task) error {
	// validate the asset is in the group
	asset, err := a.GetAsset(groupTitle, assetTitle)
	if err != nil {
		return err
	}

	if task.AssetId != uuid.Nil && task.AssetId != asset.Id {
		return ae.New(ae.CodeInvalid, "asset id mismatch")
	}

	if task.Title == "" {
		return ae.New(ae.CodeInvalid, "task title required")
	}

	// task.AssetId either nil or asset.Id
	task.AssetId = asset.Id
	return nil
}
