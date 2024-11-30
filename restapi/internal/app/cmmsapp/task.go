package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateTask(groupTitle string, assetTitle string, at tp.Task) (tp.Task, error) {
	// TODO: validate and populate asset task

	return a.db.CreateTask(at)
}

func (a *App) DeleteTask(groupTitle string, assetTitle string, taskId string) error {
	// TODO: validate asset task

	// cast taskId to tp.UUID
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return err
	}

	return a.db.DeleteTask(atId)
}

func (a *App) ListTasks(groupTitle string, assetTitle string) ([]tp.Task, error) {
	ats, err := a.db.ListTasks()
	if err != nil {
		return []tp.Task{}, err
	}

	// filter asset tasks by group and asset title
	var tasks []tp.Task
	for _, at := range ats {
		// if at.GroupTitle == groupTitle && at.AssetTitle == assetTitle {
		tasks = append(tasks, at)
		// }
	}

	return tasks, nil
}

func (a *App) GetTask(groupTitle string, assetTitle string, taskId string) (tp.Task, error) {
	// TODO: validate asset task

	// cast taskId to tp.UUID
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, err
	}

	return a.db.GetTask(atId)
}

func (a *App) UpdateTask(groupTitle string, assetTitle string, taskId string, at tp.Task) (tp.Task, error) {
	// TODO: validate asset task

	// cast taskId to tp.UUID
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.Task{}, err
	}

	return a.db.UpdateTask(atId, at)
}
