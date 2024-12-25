package integration

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestTaskCreate(t *testing.T) {
	t.Parallel()
	dbname := "testtaskcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	// Create
	at := tp.Task{
		Title:          "testtask1",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        setupAsset(t, store, "1"),
		TaskTemplateId: utest.ToPtr(setupTaskTemplate(t, store, "1")),
	}
	returnedTask, err := store.CreateTask(at)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, at, returnedTask, fieldsToExclude)
}

func TestTaskDelete(t *testing.T) {
	t.Parallel()
	dbname := "testtaskdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	// Create
	at := tp.Task{
		Title:          "testtask1",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        setupAsset(t, store, "1"),
		TaskTemplateId: utest.ToPtr(setupTaskTemplate(t, store, "1")),
	}

	at, err := store.CreateTask(at)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	err = store.DeleteTask(at.Id)
	if err != nil {
		t.Errorf("DeleteTask() failed: %v", err)
	}

	// Get
	_, err = store.GetTask(at.Id)
	if err == nil {
		t.Errorf("GetTask() failed: expected error, got nil")
	}
}

func TestTaskDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbname := "testtaskdeletenotfound"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	err := store.DeleteTask(uuid.UUID{})
	if err == nil {
		t.Errorf("DeleteTask() should have failed")
	}
}

func TestTaskList(t *testing.T) {
	t.Parallel()
	dbname := "testtasklist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	// List
	tasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("ListTasks() failed: %v", err)
	}

	// create a map of the tasks title: tp.task
	taskMap := make(map[string]tp.Task)
	for _, task := range tasks {
		taskMap[task.Title] = task
	}

	taskMap["testtask1"] = tp.Task{
		Title:          "testtask1",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        setupAsset(t, store, "1"),
		TaskTemplateId: utest.ToPtr(setupTaskTemplate(t, store, "1")),
	}

	taskMap["testtask2"] = tp.Task{
		Title:          "testtask2",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        setupAsset(t, store, "2"),
		TaskTemplateId: utest.ToPtr(setupTaskTemplate(t, store, "2")),
	}

	// Create the tasks
	taskMap["testtask1"], err = store.CreateTask(taskMap["testtask1"])
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	taskMap["testtask2"], err = store.CreateTask(taskMap["testtask2"])
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	// List
	newTasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("ListTasks() failed: %v", err)
	}

	if len(newTasks) != len(taskMap) {
		t.Errorf("ListTasks() failed: expected %v, got %v", len(taskMap), len(newTasks))
	}

	// create a map of the tasks title: tp.task
	newTaskMap := make(map[string]tp.Task)
	for _, task := range newTasks {
		newTaskMap[task.Title] = task
	}

	// compare the 2 maps
	for title, task := range taskMap {
		utest.CompEntities(t, task, newTaskMap[title])
	}
}

func TestTaskUpdateGet(t *testing.T) {
	t.Parallel()
	dbname := "testtaskupdateget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	// Create
	at := tp.Task{
		Title:          "testtask1",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        setupAsset(t, store, "1"),
		TaskTemplateId: utest.ToPtr(setupTaskTemplate(t, store, "1")),
	}

	createAt, err := store.CreateTask(at)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	// Update
	at.Title = "testtask1updated"
	at.Instructions = utest.ToPtr("test instructions updated")
	at.AssetId = setupAsset(t, store, "2")
	at.TaskTemplateId = utest.ToPtr(setupTaskTemplate(t, store, "2"))
	updateAt, err := store.UpdateTask(createAt.Id, at)
	if err != nil {
		t.Errorf("UpdateTask() failed: %v", err)
	}

	fields := utest.ConvertToSet([]string{"Title", "Instructions", "AssetId", "TaskTemplateId"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createAt, updateAt, fields)

	// Get
	getAt, err := store.GetTask(updateAt.Id)
	if err != nil {
		t.Errorf("GetTask() failed: %v", err)
	}

	utest.CompEntities(t, updateAt, getAt)
}

func TestTaskUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbname := "testtaskupdatenotfound"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	at := tp.Task{
		Id:             uuid.UUID{},
		Title:          "testtask1",
		Instructions:   utest.ToPtr("test instructions"),
		AssetId:        uuid.UUID{},
		TaskTemplateId: utest.ToPtr(uuid.UUID{}),
	}
	_, err := store.UpdateTask(at.Id, at)
	if err == nil {
		t.Errorf("UpdateTask() should have failed")
	}
}
