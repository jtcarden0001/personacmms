package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateTaskConsumable(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	// setup group, asset, task, and consumable
	_, _, task := setupGroupAssetTask(app)
	consumable, _ := db.CreateConsumable(tp.Consumable{Title: "consumable1"})

	taskConsumable := tp.TaskConsumable{TaskId: task.Id, ConsumableId: consumable.Id, QuantityNote: "note1"}
	createdTaskConsumable, err := app.CreateTaskConsumable(taskConsumable)
	assert.NoError(t, err)
	assert.Equal(t, taskConsumable.TaskId, createdTaskConsumable.TaskId)
	assert.Equal(t, taskConsumable.ConsumableId, createdTaskConsumable.ConsumableId)
}

func TestDeleteTaskConsumable(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	// setup group, asset, task, and consumable
	_, _, task := setupGroupAssetTask(app)
	consumable, _ := db.CreateConsumable(tp.Consumable{Title: "consumable1"})
	db.CreateTaskConsumable(tp.TaskConsumable{TaskId: task.Id, ConsumableId: consumable.Id})

	err := app.DeleteTaskConsumable("group1", "asset1", task.Id.String(), consumable.Id.String())
	assert.NoError(t, err)
}

func TestListTaskConsumables(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	// setup group, asset, task, and consumable
	group, asset, task1 := setupGroupAssetTask(app)
	task2, _ := app.CreateTask(group.Title, asset.Title, tp.Task{Title: "task2"})
	consumable, _ := db.CreateConsumable(tp.Consumable{Title: "consumable1"})
	db.CreateTaskConsumable(tp.TaskConsumable{TaskId: task1.Id, ConsumableId: consumable.Id})
	db.CreateTaskConsumable(tp.TaskConsumable{TaskId: task2.Id, ConsumableId: consumable.Id})

	taskConsumables1, err := app.ListTaskConsumables("group1", "asset1", task1.Id.String())
	assert.NoError(t, err)
	taskConsumables2, err := app.ListTaskConsumables("group1", "asset1", task2.Id.String())
	assert.NoError(t, err)
	tcs := append(taskConsumables1, taskConsumables2...)
	assert.Len(t, tcs, 2)
}
