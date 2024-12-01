package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTaskCreate(t *testing.T) {
	dbname := "testtaskcreate"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.Task{
		Title:              "testtask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}
	returnedTask, err := store.CreateTask(at)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, at, returnedTask, fieldsToExclude)
}

func TestTaskDelete(t *testing.T) {
	dbname := "testtaskdelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.Task{
		Title:              "testtask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
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

func TestTaskList(t *testing.T) {
	dbname := "testtasklist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

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
		Title:              "testtask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}

	taskMap["testtask2"] = tp.Task{
		Title:              "testtask2",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "2"),
		TaskTemplateId:     setupTaskTemplate(t, store, "2"),
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
		compEntities(t, task, newTaskMap[title])
	}
}

func TestTaskUpdateGet(t *testing.T) {
	dbname := "testtaskupdateget"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.Task{
		Title:              "testtask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}

	createAt, err := store.CreateTask(at)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	// Update
	at.Title = "testtask1updated"
	at.UniqueInstructions = "test instructions updated"
	at.AssetId = setupAsset(t, store, "2")
	at.TaskTemplateId = setupTaskTemplate(t, store, "2")
	updateAt, err := store.UpdateTask(createAt.Id, at)
	if err != nil {
		t.Errorf("UpdateTask() failed: %v", err)
	}

	fields := convertToSet([]string{"Title", "UniqueInstructions", "AssetId", "TaskTemplateId"})
	compEntitiesFieldsShouldBeDifferent(t, createAt, updateAt, fields)

	// Get
	getAt, err := store.GetTask(updateAt.Id)
	if err != nil {
		t.Errorf("GetTask() failed: %v", err)
	}

	compEntities(t, updateAt, getAt)
}
