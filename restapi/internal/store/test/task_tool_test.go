package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTaskToolCreateGet(t *testing.T) {
	dbname := "testtasktoolcreate"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.TaskTool{
		TaskId: setupTask(t, store, "1"),
		ToolId: setupTool(t, store, "1"),
	}
	returnedTaskTool, err := store.CreateTaskTool(at)
	if err != nil {
		t.Errorf("CreateTaskTool() failed %v", err)
	}

	compEntities(t, at, returnedTaskTool)

	// Get
	at2, err := store.GetTaskTool(at.TaskId, at.ToolId)
	if err != nil {
		t.Errorf("GetTaskTool() failed %v", err)
	}

	compEntities(t, at, at2)
}

func TestTaskToolDelete(t *testing.T) {
	dbname := "testtasktooldelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.TaskTool{
		TaskId: setupTask(t, store, "1"),
		ToolId: setupTool(t, store, "1"),
	}
	returnedTaskTool, err := store.CreateTaskTool(at)
	if err != nil {
		t.Errorf("CreateTaskTool() failed %v", err)
	}

	compEntities(t, at, returnedTaskTool)

	// Delete
	err = store.DeleteTaskTool(at.TaskId, at.ToolId)
	if err != nil {
		t.Errorf("DeleteTaskTool() failed %v", err)
	}

	_, err = store.GetTaskTool(at.TaskId, at.ToolId)
	if err == nil {
		t.Errorf("GetTaskTool() failed: expected error")
	}
}

func TestTaskToolList(t *testing.T) {
	dbname := "testtasktoollist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	ats, err := store.ListTaskTools()
	if err != nil {
		t.Errorf("ListTaskTools() failed %v", err)
	}
	origCount := len(ats)

	// Create
	at1 := tp.TaskTool{
		TaskId: setupTask(t, store, "1"),
		ToolId: setupTool(t, store, "1"),
	}
	at2 := tp.TaskTool{
		TaskId: setupTask(t, store, "2"),
		ToolId: setupTool(t, store, "2"),
	}
	at3 := tp.TaskTool{
		TaskId: setupTask(t, store, "3"),
		ToolId: setupTool(t, store, "3"),
	}
	_, err = store.CreateTaskTool(at1)
	if err != nil {
		t.Errorf("CreateTaskTool() failed %v", err)
	}
	_, err = store.CreateTaskTool(at2)
	if err != nil {
		t.Errorf("CreateTaskTool() failed %v", err)
	}
	_, err = store.CreateTaskTool(at3)
	if err != nil {
		t.Errorf("CreateTaskTool() failed %v", err)
	}

	// List
	ats, err = store.ListTaskTools()
	if err != nil {
		t.Errorf("ListTaskTools() failed %v", err)
	}

	expCount := 3 + origCount
	if len(ats) != expCount {
		t.Errorf("ListTaskTools() failed: expected %d, got %d", expCount, len(ats))
	}
}
