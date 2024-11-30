package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetTaskCreate(t *testing.T) {
	dbname := "testassettaskcreate"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTask{
		Title:              "testassettask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}
	returnedAssetTask, err := store.CreateAssetTask(at)
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, at, returnedAssetTask, fieldsToExclude)
}

func TestAssetTaskDelete(t *testing.T) {
	dbname := "testassettaskdelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTask{
		Title:              "testassettask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}

	at, err := store.CreateAssetTask(at)
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	err = store.DeleteAssetTask(at.Id)
	if err != nil {
		t.Errorf("DeleteAssetTask() failed: %v", err)
	}

	// Get
	_, err = store.GetAssetTask(at.Id)
	if err == nil {
		t.Errorf("GetAssetTask() failed: expected error, got nil")
	}
}

func TestAssetTaskList(t *testing.T) {
	dbname := "testassettasklist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	assetTasks, err := store.ListAssetTasks()
	if err != nil {
		t.Errorf("ListAssetTasks() failed: %v", err)
	}

	// create a map of the assetTasks title: tp.assetTask
	assetTaskMap := make(map[string]tp.AssetTask)
	for _, assetTask := range assetTasks {
		assetTaskMap[assetTask.Title] = assetTask
	}

	assetTaskMap["testassettask1"] = tp.AssetTask{
		Title:              "testassettask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}

	assetTaskMap["testassettask2"] = tp.AssetTask{
		Title:              "testassettask2",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "2"),
		TaskTemplateId:     setupTaskTemplate(t, store, "2"),
	}

	// Create the assetTasks
	assetTaskMap["testassettask1"], err = store.CreateAssetTask(assetTaskMap["testassettask1"])
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	assetTaskMap["testassettask2"], err = store.CreateAssetTask(assetTaskMap["testassettask2"])
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	// List
	newAssetTasks, err := store.ListAssetTasks()
	if err != nil {
		t.Errorf("ListAssetTasks() failed: %v", err)
	}

	if len(newAssetTasks) != len(assetTaskMap) {
		t.Errorf("ListAssetTasks() failed: expected %v, got %v", len(assetTaskMap), len(newAssetTasks))
	}

	// create a map of the assetTasks title: tp.assetTask
	newAssetTaskMap := make(map[string]tp.AssetTask)
	for _, assetTask := range newAssetTasks {
		newAssetTaskMap[assetTask.Title] = assetTask
	}

	// compare the 2 maps
	for title, assetTask := range assetTaskMap {
		compEntities(t, assetTask, newAssetTaskMap[title])
	}
}

func TestAssetTaskUpdateGet(t *testing.T) {
	dbname := "testassettaskupdateget"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTask{
		Title:              "testassettask1",
		UniqueInstructions: "test instructions",
		AssetId:            setupAsset(t, store, "1"),
		TaskTemplateId:     setupTaskTemplate(t, store, "1"),
	}

	createAt, err := store.CreateAssetTask(at)
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	// Update
	at.Title = "testassettask1updated"
	at.UniqueInstructions = "test instructions updated"
	at.AssetId = setupAsset(t, store, "2")
	at.TaskTemplateId = setupTaskTemplate(t, store, "2")
	updateAt, err := store.UpdateAssetTask(createAt.Id, at)
	if err != nil {
		t.Errorf("UpdateAssetTask() failed: %v", err)
	}

	fields := convertToSet([]string{"Title", "UniqueInstructions", "AssetId", "TaskTemplateId"})
	compEntitiesFieldsShouldBeDifferent(t, createAt, updateAt, fields)

	// Get
	getAt, err := store.GetAssetTask(updateAt.Id)
	if err != nil {
		t.Errorf("GetAssetTask() failed: %v", err)
	}

	compEntities(t, updateAt, getAt)
}
