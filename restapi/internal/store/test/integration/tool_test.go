package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateTool(t *testing.T) {
	t.Parallel()
	dbName := "testcreatetool"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  utest.ToPtr("13mm"),
	}

	returnedTool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, tool, returnedTool, fieldsToExclude)
}

func TestDeleteTool(t *testing.T) {
	t.Parallel()
	dbName := "testdeletetool"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  utest.ToPtr("13mm"),
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

func TestDeleteToolNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testdeletetoolnotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteTool("nonexistent-title")
	if err == nil {
		t.Errorf("DeleteTool() failed: expected error, got nil")
	}
}

func TestListTool(t *testing.T) {
	t.Parallel()
	dbName := "testlisttool"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	tools, err := store.ListTools()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}
	count := len(tools)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  utest.ToPtr("13mm"),
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
	t.Parallel()
	dbName := "testupdategettool"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	tool := tp.Tool{
		Title: "testtool1",
		Size:  utest.ToPtr("13mm"),
	}

	cTool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	tool.Size = utest.ToPtr("14mm")
	uTool, err := store.UpdateTool(tool.Title, tool)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	fieldsShouldBeDifferent := utest.ConvertToSet([]string{"Size"})
	compEntitiesFieldsShouldBeDifferent(t, cTool, uTool, fieldsShouldBeDifferent)

	// Get
	gTool, err := store.GetTool(tool.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	compEntities(t, uTool, gTool)

}

func TestUpdateToolNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testupdatetoolnotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tool := tp.Tool{
		Title: "nonexistent-title",
		Size:  utest.ToPtr("13mm"),
	}
	_, err := store.UpdateTool("nonexistent-title", tool)
	if err == nil {
		t.Errorf("UpdateTool() failed: expected error, got nil")
	}
}
