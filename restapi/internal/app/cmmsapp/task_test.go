package cmmsapp

// import (
// 	"testing"

// 	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	"github.com/stretchr/testify/assert"
// )

// func setupGroupAsset(app *App) (tp.Group, tp.Asset) {
// 	db := app.db.(*mock.MockStore)
// 	group, _ := db.CreateGroup(tp.Group{Title: "group1"})
// 	asset, _ := db.CreateAsset(tp.Asset{Title: "asset1", GroupTitle: group.Title})
// 	return group, asset
// }

// func TestCreateTask(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	// create asset and group
// 	group, asset := setupGroupAsset(app)
// 	task := tp.Task{Title: "task1", AssetId: asset.Id}
// 	createdTask, err := app.CreateTask(group.Title, asset.Title, task)
// 	assert.NoError(t, err)
// 	assert.Equal(t, task.Title, createdTask.Title)
// }

// func TestDeleteTask(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}
// 	group, asset := setupGroupAsset(app)
// 	task := tp.Task{Title: "task1", AssetId: asset.Id}
// 	task, _ = db.CreateTask(task)

// 	err := app.DeleteTask(group.Title, asset.Title, task.Id.String())
// 	assert.NoError(t, err)
// }

// func TestListTasks(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset := setupGroupAsset(app)
// 	task1 := tp.Task{Title: "task1", AssetId: asset.Id}
// 	task2 := tp.Task{Title: "task2", AssetId: asset.Id}
// 	db.CreateTask(task1)
// 	db.CreateTask(task2)

// 	tasks, err := app.ListTasks(group.Title, asset.Title)
// 	assert.NoError(t, err)
// 	assert.Len(t, tasks, 2)
// }

// func TestGetTask(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset := setupGroupAsset(app)
// 	task := tp.Task{Title: "task1", AssetId: asset.Id}
// 	task, _ = db.CreateTask(task)

// 	retrievedTask, err := app.GetTask(group.Title, asset.Title, task.Id.String())
// 	assert.NoError(t, err)
// 	assert.Equal(t, task.Title, retrievedTask.Title)
// }

// func TestUpdateTask(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset := setupGroupAsset(app)
// 	task := tp.Task{Title: "task1", AssetId: asset.Id}
// 	task, _ = db.CreateTask(task)

// 	updatedTask := tp.Task{Title: "task1_updated"}
// 	_, err := app.UpdateTask(group.Title, asset.Title, task.Id.String(), updatedTask)
// 	assert.NoError(t, err)

// 	retrievedTask, err := app.GetTask(group.Title, asset.Title, task.Id.String())
// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedTask.Title, retrievedTask.Title)
// }
