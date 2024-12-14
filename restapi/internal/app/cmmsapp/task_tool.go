package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// create tan association between a task and a tool
func (a *App) CreateTaskTool(tool tp.TaskTool) (tp.TaskTool, error) {
	err := a.validateTaskTool(tool)
	if err != nil {
		return tp.TaskTool{}, err
	}

	return a.db.CreateTaskTool(tool)
}

// create an association between a task and a tool with task namespace validation
func (a *App) CreateTaskToolWithValidation(groupTitle, assetTitle, taskId, toolId string) (tp.TaskTool, error) {
	// validate the namespace coherency of the task
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	// construct the target task tool
	tt := tp.TaskTool{
		TaskId: t.Id,
		ToolId: tId,
	}

	// TODO: some inefficiency here as task is being queried from the db twice, room for improvement
	err = a.validateTaskTool(tt)
	if err != nil {
		return tp.TaskTool{}, err
	}

	return a.db.CreateTaskTool(tt)
}

// delete the relationship between a task and a tool
func (a *App) DeleteTaskTool(groupTitle, assetTitle, taskId, toolId string) error {
	// validate the tasktool exists
	tt, err := a.GetTaskTool(groupTitle, assetTitle, taskId, toolId)
	if err != nil {
		return err
	}

	return a.db.DeleteTaskTool(tt.TaskId, tt.ToolId)
}

// list all tools associates with a task
func (a *App) ListTaskTools(groupTitle, assetTitle, taskId string) ([]tp.TaskTool, error) {
	// namespace validation
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.TaskTool{}, err
	}

	atTools, err := a.db.ListTaskToolsByTaskId(t.Id)
	if err != nil {
		return []tp.TaskTool{}, err
	}

	return atTools, nil
}

// get a tool associated with a task
func (a *App) GetTaskTool(groupTitle, assetTitle, taskId, toolId string) (tp.TaskTool, error) {
	// namespace validation
	task, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.TaskTool{}, err
	}

	return a.db.GetTaskTool(task.Id, tId)
}

func (a *App) validateTaskTool(tool tp.TaskTool) error {
	// validate task exists
	_, err := a.db.GetTask(tool.TaskId)
	if err != nil {
		return err
	}

	// validate tool exists
	_, err = a.db.GetToolById(tool.ToolId)
	if err != nil {
		return err
	}

	return nil
}
