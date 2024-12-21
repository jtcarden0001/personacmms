package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func setupGroupAssetTaskTool(app *App) (tp.Group, tp.Asset, tp.Task, tp.Tool) {
	db := app.db.(*mock.MockStore)
	group, _ := db.CreateGroup(tp.Group{Title: "group1"})
	asset, _ := db.CreateAsset(tp.Asset{Title: "asset1", GroupTitle: "group1"})
	task, _ := app.CreateTask(group.Title, asset.Title, tp.Task{Title: "task1"})
	tool, _ := db.CreateTool(tp.Tool{Title: "tool1"})
	return group, asset, task, tool
}

func TestCreateTaskTool(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	_, _, task, tool := setupGroupAssetTaskTool(app)
	taskTool := tp.TaskTool{TaskId: task.Id, ToolId: tool.Id}
	createdTaskTool, err := app.CreateTaskTool(taskTool)
	assert.NoError(t, err)
	assert.Equal(t, taskTool.TaskId, createdTaskTool.TaskId)
}

func TestDeleteTaskTool(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task, tool := setupGroupAssetTaskTool(app)
	db.CreateTaskTool(tp.TaskTool{TaskId: task.Id, ToolId: tool.Id})

	err := app.DeleteTaskTool(group.Title, asset.Title, task.Id.String(), tool.Id.String())
	assert.NoError(t, err)
}

func TestListTaskTools(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task, tool := setupGroupAssetTaskTool(app)
	tool2, _ := db.CreateTool(tp.Tool{Title: "tool2"})
	db.CreateTaskTool(tp.TaskTool{TaskId: task.Id, ToolId: tool.Id})
	db.CreateTaskTool(tp.TaskTool{TaskId: task.Id, ToolId: tool2.Id})

	taskTools, err := app.ListTaskTools(group.Title, asset.Title, task.Id.String())
	assert.NoError(t, err)
	assert.Len(t, taskTools, 2)
}

func TestGetTaskTool(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task, tool := setupGroupAssetTaskTool(app)
	db.CreateTaskTool(tp.TaskTool{TaskId: task.Id, ToolId: tool.Id})

	retrievedTaskTool, err := app.GetTaskTool(group.Title, asset.Title, task.Id.String(), tool.Id.String())
	assert.NoError(t, err)
	assert.Equal(t, task.Id, retrievedTaskTool.TaskId)
	assert.Equal(t, tool.Id, retrievedTaskTool.ToolId)
}
