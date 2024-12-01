package integration

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestCreateTool(t *testing.T) {
	dbName := "testcreatetool"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  "13mm",
	}

	returnedTool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnedTool.Title != tool.Title {
		t.Errorf("Create() failed: expected %s, got %s", tool.Title, returnedTool.Title)
	}

	if returnedTool.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-empty ID, got empty")
	}

	if returnedTool.Size != tool.Size {
		t.Errorf("Create() failed: expected %s, got %s", tool.Size, returnedTool.Size)
	}
}

func TestDeleteTool(t *testing.T) {
	dbName := "testdeletetool"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  "13mm",
	}

	_, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Delete
	err = store.DeleteTool(tool.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Confirm deletion
	_, err = store.GetTool(tool.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestListTool(t *testing.T) {
	dbName := "testlisttool"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// List
	tools, err := store.ListTools()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}
	count := len(tools)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  "13mm",
	}
	_, err = store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	tool.Title = "testtool2"
	_, err = store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	tools, err = store.ListTools()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(tools) != 2+count {
		t.Errorf("ListTool() failed: expected 2, got %d", len(tools))
	}

	// TODO: test list content
}

func TestUpdateGetTool(t *testing.T) {
	dbName := "testupdategettool"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  "13mm",
	}

	_, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	tool.Size = "14mm"
	_, err = store.UpdateTool(tool.Title, tool)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	returnedTool, err := store.GetTool(tool.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returnedTool.Size != tool.Size {
		t.Errorf("Get() failed: expected %s, got %s", tool.Size, returnedTool.Size)
	}
}
