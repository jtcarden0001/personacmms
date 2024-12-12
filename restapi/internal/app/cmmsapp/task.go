package cmmsapp

import (
	"github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
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

// TODO: complete this vertical
func (a *App) ListTasks(groupTitle string, assetTitle string) ([]tp.Task, error) {
	ats, err := a.db.ListTasks()
	if err != nil {
		return []tp.Task{}, err
	}

	// filter Tsks by group and asset title
	var tasks []tp.Task
	for _, at := range ats {
		// if at.GroupTitle == groupTitle && at.AssetTitle == assetTitle {
		tasks = append(tasks, at)
		// }
	}

	return tasks, nil
}

func (a *App) UpdateTask(groupTitle string, assetTitle string, taskId string, at tp.Task) (tp.Task, error) {
	// TODO: validate Tsk

	// cast taskId to tp.UUID
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, err
	}

	return a.db.UpdateTask(atId, at)
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
