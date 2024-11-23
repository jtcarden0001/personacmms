package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type PreventativeTaskTool interface {
	CreatePreventativeTaskTool(int, int) error
	DeletePreventativeTaskTool(int, int) error
	GetAllPreventativeTaskTool() ([]tp.PreventativeTaskTool, error)
	GetAllPreventativeTaskToolByPreventativeTaskId(int) ([]tp.PreventativeTaskTool, error)
	GetPreventativeTaskTool(int, int) (tp.PreventativeTaskTool, error)
}

func (a *App) CreatePreventativeTaskTool(preventativeTaskId int, toolId int) error {
	return a.db.CreatePreventativeTaskTool(preventativeTaskId, toolId)
}

func (a *App) DeletePreventativeTaskTool(preventativeTaskId int, toolId int) error {
	return a.db.DeletePreventativeTaskTool(preventativeTaskId, toolId)
}

func (a *App) GetAllPreventativeTaskTool() ([]tp.PreventativeTaskTool, error) {
	return a.db.GetAllPreventativeTaskTool()
}

func (a *App) GetAllPreventativeTaskToolByPreventativeTaskId(preventativeTaskId int) ([]tp.PreventativeTaskTool, error) {
	return a.db.GetAllPreventativeTaskToolByPreventativeTaskId(preventativeTaskId)
}

func (a *App) GetPreventativeTaskTool(preventativeTaskId int, toolId int) (tp.PreventativeTaskTool, error) {
	return a.db.GetPreventativeTaskTool(preventativeTaskId, toolId)
}
