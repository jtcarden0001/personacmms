package postgres_test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

// TODO: TestToolAssociateWithTask

// TODO: TestToolAssociateWithOrder

func TestToolCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtoolcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tool := utest.SetupTool(1, true)

	// test
	createdTool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("CreateTool() failed: %v", err)
	}

	utest.CompEntities(t, tool, createdTool)
}

func TestToolDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtooldelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tool := utest.SetupTool(1, true)
	_, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("TestToolDelete: failed during setup. CreateTool() failed: %v", err)
	}

	// test
	err = store.DeleteTool(tool.Id)
	if err != nil {
		t.Errorf("TestToolDelete: DeleteTool() failed: %v", err)
	}

	_, err = store.GetTool(tool.Id)
	if err == nil {
		t.Errorf("TestToolDelete: GetTool() returned nil error after deletion")
	}
}

// TODO: TestToolDisassociateWithTask

// TODO: TestToolDisassociateWithOrder

func TestToolGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtoolget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tool := utest.SetupTool(1, true)
	_, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("TestToolGet: failed during setup. CreateTool() failed: %v", err)
	}

	// test
	getTool, err := store.GetTool(tool.Id)
	if err != nil {
		t.Errorf("GetTool() failed: %v", err)
	}

	utest.CompEntities(t, tool, getTool)
}

func TestToolList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtoollist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tool1 := utest.SetupTool(1, true)
	tool2 := utest.SetupTool(2, true)
	tool3 := utest.SetupTool(3, true)

	_, err := store.CreateTool(tool1)
	if err != nil {
		t.Errorf("TestToolList: failed during setup. CreateTool() failed: %v", err)
	}
	_, err = store.CreateTool(tool2)
	if err != nil {
		t.Errorf("TestToolList: failed during setup. CreateTool() failed: %v", err)
	}
	_, err = store.CreateTool(tool3)
	if err != nil {
		t.Errorf("TestToolList: failed during setup. CreateTool() failed: %v", err)
	}

	// test
	tools, err := store.ListTools()
	if err != nil {
		t.Errorf("ListTools() failed: %v", err)
	}

	if len(tools) != 3 {
		t.Errorf("ListTools() failed: expected 3 tools, got %d", len(tools))
	}

	toolMap := map[string]tp.Tool{
		tool1.Title: tool1,
		tool2.Title: tool2,
		tool3.Title: tool3,
	}

	for _, tool := range tools {
		expectedTool, ok := toolMap[tool.Title]
		if !ok {
			t.Errorf("ListTools() failed: unexpected tool with Title %v", tool.Title)
		}
		utest.CompEntities(t, expectedTool, tool)
	}
}

// TODO: TestToolListByTask

// TODO: TestToolListByWorkOrder

func TestToolUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtoolupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tool := utest.SetupTool(1, true)
	createTool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("TestToolUpdate: failed during setup. CreateTool() failed: %v", err)
	}

	// test
	tool.Title = "Updated Tool Title"

	updatedTool, err := store.UpdateTool(tool)
	if err != nil {
		t.Errorf("UpdateTool() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createTool, updatedTool, differentFields)
}
