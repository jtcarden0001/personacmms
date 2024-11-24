package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type TaskTool interface {
	CreateTaskTool(int, int) error
	DeleteTaskTool(int, int) error
	GetAllTaskTool() ([]tp.TaskTool, error)
	GetAllTaskToolByTaskId(int) ([]tp.TaskTool, error)
	GetTaskTool(int, int) (tp.TaskTool, error)
}

func (a *App) CreateTaskTool(preventativeTaskId int, toolId int) error {
	return a.db.CreateTaskTool(preventativeTaskId, toolId)
}

func (a *App) DeleteTaskTool(preventativeTaskId int, toolId int) error {
	return a.db.DeleteTaskTool(preventativeTaskId, toolId)
}

func (a *App) GetAllTaskTool() ([]tp.TaskTool, error) {
	return a.db.GetAllTaskTool()
}

func (a *App) GetAllTaskToolByTaskId(preventativeTaskId int) ([]tp.TaskTool, error) {
	return a.db.GetAllTaskToolByTaskId(preventativeTaskId)
}

func (a *App) GetTaskTool(preventativeTaskId int, toolId int) (tp.TaskTool, error) {
	return a.db.GetTaskTool(preventativeTaskId, toolId)
}
