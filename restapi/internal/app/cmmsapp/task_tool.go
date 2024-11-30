package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateTaskTool(tool tp.TaskTool) (tp.TaskTool, error) {
	return a.db.CreateTaskTool(tool)
}

func (a *App) CreateTaskToolWithValidation(groupTitle, assetTitle, taskId, toolId string) (tp.TaskTool, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	return a.db.CreateTaskTool(tp.TaskTool{TaskId: atId, ToolId: tId})
}

func (a *App) DeleteTaskTool(groupTitle, assetTitle, taskId, toolId string) error {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return err
	}

	return a.db.DeleteTaskTool(atId, tId)
}

func (a *App) ListTaskTools(groupTitle, assetTitle, taskId string) ([]tp.TaskTool, error) {
	// TODO: implement validation

	atTools, err := a.db.ListTaskTools()
	if err != nil {
		return nil, err
	}

	// TODO: filter asset task tools by asset task id

	return atTools, nil
}

func (a *App) GetTaskTool(groupTitle, assetTitle, taskId, toolId string) (tp.TaskTool, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(taskId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	return a.db.GetTaskTool(atId, tId)
}
