package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type TaskTool interface {
	CreateTaskTool(int, int) error
	DeleteTaskTool(int, int) error
	GetAllTaskTool() ([]tp.TaskTool, error)
	GetAllTaskToolByTaskId(int) ([]tp.TaskTool, error)
	GetTaskTool(int, int) (tp.TaskTool, error)
}

func (a *App) CreateTaskTool(taskId int, toolId int) error {
	return a.db.CreateTaskTool(taskId, toolId)
}

func (a *App) DeleteTaskTool(taskId int, toolId int) error {
	return a.db.DeleteTaskTool(taskId, toolId)
}

func (a *App) GetAllTaskTool() ([]tp.TaskTool, error) {
	return a.db.GetAllTaskTool()
}

func (a *App) GetAllTaskToolByTaskId(taskId int) ([]tp.TaskTool, error) {
	return a.db.GetAllTaskToolByTaskId(taskId)
}

func (a *App) GetTaskTool(taskId int, toolId int) (tp.TaskTool, error) {
	return a.db.GetTaskTool(taskId, toolId)
}
