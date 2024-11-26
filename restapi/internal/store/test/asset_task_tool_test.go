package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetTaskToolCreateGet(t *testing.T) {
	dbname := "testassettasktoolcreate"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTaskTool{
		AssetTaskId: setupAssetTask(t, store, "1"),
		ToolId:      setupTool(t, store, "1"),
	}
	returnedAssetTaskTool, err := store.CreateAssetTaskTool(at)
	if err != nil {
		t.Errorf("CreateAssetTaskTool() failed %v", err)
	}

	compEntities(t, at, returnedAssetTaskTool)

	// Get
	at2, err := store.GetAssetTaskTool(at.AssetTaskId, at.ToolId)
	if err != nil {
		t.Errorf("GetAssetTaskTool() failed %v", err)
	}

	compEntities(t, at, at2)
}

func TestAssetTaskToolDelete(t *testing.T) {
	dbname := "testassettasktooldelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTaskTool{
		AssetTaskId: setupAssetTask(t, store, "1"),
		ToolId:      setupTool(t, store, "1"),
	}
	returnedAssetTaskTool, err := store.CreateAssetTaskTool(at)
	if err != nil {
		t.Errorf("CreateAssetTaskTool() failed %v", err)
	}

	compEntities(t, at, returnedAssetTaskTool)

	// Delete
	err = store.DeleteAssetTaskTool(at.AssetTaskId, at.ToolId)
	if err != nil {
		t.Errorf("DeleteAssetTaskTool() failed %v", err)
	}

	_, err = store.GetAssetTaskTool(at.AssetTaskId, at.ToolId)
	if err == nil {
		t.Errorf("GetAssetTaskTool() failed: expected error")
	}
}

func TestAssetTaskToolList(t *testing.T) {
	dbname := "testassettasktoollist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	ats, err := store.ListAssetTaskTools()
	if err != nil {
		t.Errorf("ListAssetTaskTools() failed %v", err)
	}
	origCount := len(ats)

	// Create
	at1 := tp.AssetTaskTool{
		AssetTaskId: setupAssetTask(t, store, "1"),
		ToolId:      setupTool(t, store, "1"),
	}
	at2 := tp.AssetTaskTool{
		AssetTaskId: setupAssetTask(t, store, "2"),
		ToolId:      setupTool(t, store, "2"),
	}
	at3 := tp.AssetTaskTool{
		AssetTaskId: setupAssetTask(t, store, "3"),
		ToolId:      setupTool(t, store, "3"),
	}
	_, err = store.CreateAssetTaskTool(at1)
	if err != nil {
		t.Errorf("CreateAssetTaskTool() failed %v", err)
	}
	_, err = store.CreateAssetTaskTool(at2)
	if err != nil {
		t.Errorf("CreateAssetTaskTool() failed %v", err)
	}
	_, err = store.CreateAssetTaskTool(at3)
	if err != nil {
		t.Errorf("CreateAssetTaskTool() failed %v", err)
	}

	// List
	ats, err = store.ListAssetTaskTools()
	if err != nil {
		t.Errorf("ListAssetTaskTools() failed %v", err)
	}

	expCount := 3 + origCount
	if len(ats) != expCount {
		t.Errorf("ListAssetTaskTools() failed: expected %d, got %d", expCount, len(ats))
	}
}
