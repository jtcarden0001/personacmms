package cmmsapp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssociateToolWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateToolWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateToolWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestAssociateToolWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestAssociateToolWithTask: failed during setup. CreateTool() failed: %v", err)
	}

	ts := apitp.ToolSizeRequest{
		Size: nil,
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		toolID        string
		toolSize      apitp.ToolSizeRequest
		shouldSucceed bool
	}{
		{"valid association", createdAsset.Id.String(), createdTask.Id.String(), createdTool.Id.String(), ts, true},
		{"invalid tool ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", ts, false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", createdTool.Id.String(), ts, false},
		{"invalid asset ID", "invalid", createdTask.Id.String(), createdTool.Id.String(), ts, false},
		{"invalid asset and task ID", "invalid", "invalid", createdTool.Id.String(), ts, false},
		{"invalid asset and tool ID", "invalid", createdTask.Id.String(), "invalid", ts, false},
		{"invalid task and tool ID", createdAsset.Id.String(), "invalid", "invalid", ts, false},
		{"invalid asset, task and tool ID", "invalid", "invalid", "invalid", ts, false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), createdTool.Id.String(), ts, false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), createdTool.Id.String(), ts, false},
		{"nil tool ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), ts, false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), createdTool.Id.String(), ts, false},
		{"nil asset and tool ID", uuid.Nil.String(), createdTask.Id.String(), uuid.Nil.String(), ts, false},
		{"nil task and tool ID", createdAsset.Id.String(), uuid.Nil.String(), uuid.Nil.String(), ts, false},
		{"nil asset, task and tool ID", uuid.Nil.String(), uuid.Nil.String(), uuid.Nil.String(), ts, false},
		{"empty asset ID", "", createdTask.Id.String(), createdTool.Id.String(), ts, false},
		{"empty task ID", createdAsset.Id.String(), "", createdTool.Id.String(), ts, false},
		{"empty tool ID", createdAsset.Id.String(), createdTask.Id.String(), "", ts, false},
		{"empty asset and task ID", "", "", createdTool.Id.String(), ts, false},
		{"empty asset and tool ID", "", createdTask.Id.String(), "", ts, false},
		{"empty task and tool ID", createdAsset.Id.String(), "", "", ts, false},
		{"empty asset, task and tool ID", "", "", "", ts, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.AssociateToolWithTask(tc.assetID, tc.taskID, tc.toolID, tc.toolSize)
			if tc.shouldSucceed && err != nil {
				t.Errorf("AssociateToolWithTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateToolWithTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestAssociateToolWithWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateToolWithWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateToolWithWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo := setupApiWorkOrderRequest(1, createdAsset.Id)
	createdWorkOrder, err := app.CreateWorkOrder(createdAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestAssociateToolWithWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestAssociateToolWithWorkOrder: failed during setup. CreateTool() failed: %v", err)
	}

	ts := apitp.ToolSizeRequest{
		Size: nil,
	}

	testCases := []struct {
		name          string
		assetID       string
		workOrderID   string
		toolID        string
		toolSize      apitp.ToolSizeRequest
		shouldSucceed bool
	}{
		{"valid association", createdAsset.Id.String(), createdWorkOrder.Id.String(), createdTool.Id.String(), ts, true},
		{"invalid tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), "invalid", ts, false},
		{"invalid work order ID", createdAsset.Id.String(), "invalid", createdTool.Id.String(), ts, false},
		{"invalid asset ID", "invalid", createdWorkOrder.Id.String(), createdTool.Id.String(), ts, false},
		{"invalid asset and work order ID", "invalid", "invalid", createdTool.Id.String(), ts, false},
		{"invalid asset and tool ID", "invalid", createdWorkOrder.Id.String(), "invalid", ts, false},
		{"invalid work order and tool ID", createdAsset.Id.String(), "invalid", "invalid", ts, false},
		{"invalid asset, work order and tool ID", "invalid", "invalid", "invalid", ts, false},
		{"nil asset ID", uuid.Nil.String(), createdWorkOrder.Id.String(), createdTool.Id.String(), ts, false},
		{"nil work order ID", createdAsset.Id.String(), uuid.Nil.String(), createdTool.Id.String(), ts, false},
		{"nil tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), uuid.Nil.String(), ts, false},
		{"nil asset and work order ID", uuid.Nil.String(), uuid.Nil.String(), createdTool.Id.String(), ts, false},
		{"nil asset and tool ID", uuid.Nil.String(), createdWorkOrder.Id.String(), uuid.Nil.String(), ts, false},
		{"nil work order and tool ID", createdAsset.Id.String(), uuid.Nil.String(), uuid.Nil.String(), ts, false},
		{"nil asset, work order and tool ID", uuid.Nil.String(), uuid.Nil.String(), uuid.Nil.String(), ts, false},
		{"empty asset ID", "", createdWorkOrder.Id.String(), createdTool.Id.String(), ts, false},
		{"empty work order ID", createdAsset.Id.String(), "", createdTool.Id.String(), ts, false},
		{"empty tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), "", ts, false},
		{"empty asset and work order ID", "", "", createdTool.Id.String(), ts, false},
		{"empty asset and tool ID", "", createdWorkOrder.Id.String(), "", ts, false},
		{"empty work order and tool ID", createdAsset.Id.String(), "", "", ts, false},
		{"empty asset, work order and tool ID", "", "", "", ts, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.AssociateToolWithWorkOrder(tc.assetID, tc.workOrderID, tc.toolID, tc.toolSize)
			if tc.shouldSucceed && err != nil {
				t.Errorf("AssociateToolWithWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateToolWithWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestCreateTool(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateTool")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingTool := setupApiToolRequest(1)
	_, err = app.CreateTool(conflictingTool)
	if err != nil {
		t.Errorf("TestCreateTool: failed during setup. CreateTool() failed: %v", err)
	}

	emptyTitleTool := setupApiToolRequest(2)
	emptyTitleTool.Title = ""

	testCases := []struct {
		name          string
		tool          apitp.ToolRequest
		shouldSucceed bool
	}{
		{"valid tool", setupApiToolRequest(3), true},
		{"non nil id", setupApiToolRequest(4), false},
		{"empty title", emptyTitleTool, false},
		{"conflicting title", conflictingTool, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateTool(tc.tool)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateTool() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateTool() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteTool(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteTool")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestDeleteTool: failed during setup. CreateTool() failed: %v", err)
	}

	testCases := []struct {
		name          string
		toolId        string
		shouldSucceed bool
	}{
		{"valid tool deletion", createdTool.Id.String(), true},
		{"invalid tool ID", "invalid", false},
		{"nil tool ID", uuid.Nil.String(), false},
		{"empty tool ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteTool(tc.toolId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteTool() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteTool() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateToolWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateToolWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateToolWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDisassociateToolWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestDisassociateToolWithTask: failed during setup. CreateTool() failed: %v", err)
	}

	ts := apitp.ToolSizeRequest{
		Size: nil,
	}

	_, err = app.AssociateToolWithTask(createdAsset.Id.String(), createdTask.Id.String(), createdTool.Id.String(), ts)
	if err != nil {
		t.Errorf("TestDisassociateToolWithTask: failed during setup. AssociateToolWithTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		toolID        string
		shouldSucceed bool
	}{
		{"valid disassociation", createdAsset.Id.String(), createdTask.Id.String(), createdTool.Id.String(), true},
		{"invalid tool ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", createdTool.Id.String(), false},
		{"invalid asset ID", "invalid", createdTask.Id.String(), createdTool.Id.String(), false},
		{"invalid asset and task ID", "invalid", "invalid", createdTool.Id.String(), false},
		{"invalid asset and tool ID", "invalid", createdTask.Id.String(), "invalid", false},
		{"invalid task and tool ID", createdAsset.Id.String(), "invalid", "invalid", false},
		{"invalid asset, task and tool ID", "invalid", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), createdTool.Id.String(), false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), createdTool.Id.String(), false},
		{"nil tool ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), createdTool.Id.String(), false},
		{"nil asset and tool ID", uuid.Nil.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"nil task and tool ID", createdAsset.Id.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"nil asset, task and tool ID", uuid.Nil.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdTask.Id.String(), createdTool.Id.String(), false},
		{"empty task ID", createdAsset.Id.String(), "", createdTool.Id.String(), false},
		{"empty tool ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
		{"empty asset and task ID", "", "", createdTool.Id.String(), false},
		{"empty asset and tool ID", "", createdTask.Id.String(), "", false},
		{"empty task and tool ID", createdAsset.Id.String(), "", "", false},
		{"empty asset, task and tool ID", "", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateToolWithTask(tc.assetID, tc.taskID, tc.toolID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateToolWithTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateToolWithTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateToolWithWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateToolWithWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateToolWithWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo := setupApiWorkOrderRequest(1, createdAsset.Id)
	createdWorkOrder, err := app.CreateWorkOrder(createdAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestDisassociateToolWithWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestDisassociateToolWithWorkOrder: failed during setup. CreateTool() failed: %v", err)
	}

	ts := apitp.ToolSizeRequest{
		Size: nil,
	}

	_, err = app.AssociateToolWithWorkOrder(createdAsset.Id.String(), createdWorkOrder.Id.String(), createdTool.Id.String(), ts)
	if err != nil {
		t.Errorf("TestDisassociateToolWithWorkOrder: failed during setup. AssociateToolWithWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		workOrderID   string
		toolID        string
		shouldSucceed bool
	}{
		{"valid disassociation", createdAsset.Id.String(), createdWorkOrder.Id.String(), createdTool.Id.String(), true},
		{"invalid tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), "invalid", false},
		{"invalid work order ID", createdAsset.Id.String(), "invalid", createdTool.Id.String(), false},
		{"invalid asset ID", "invalid", createdWorkOrder.Id.String(), createdTool.Id.String(), false},
		{"invalid asset and work order ID", "invalid", "invalid", createdTool.Id.String(), false},
		{"invalid asset and tool ID", "invalid", createdWorkOrder.Id.String(), "invalid", false},
		{"invalid work order and tool ID", createdAsset.Id.String(), "invalid", "invalid", false},
		{"invalid asset, work order and tool ID", "invalid", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdWorkOrder.Id.String(), createdTool.Id.String(), false},
		{"nil work order ID", createdAsset.Id.String(), uuid.Nil.String(), createdTool.Id.String(), false},
		{"nil tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), uuid.Nil.String(), false},
		{"nil asset and work order ID", uuid.Nil.String(), uuid.Nil.String(), createdTool.Id.String(), false},
		{"nil asset and tool ID", uuid.Nil.String(), createdWorkOrder.Id.String(), uuid.Nil.String(), false},
		{"nil work order and tool ID", createdAsset.Id.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"nil asset, work order and tool ID", uuid.Nil.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdWorkOrder.Id.String(), createdTool.Id.String(), false},
		{"empty work order ID", createdAsset.Id.String(), "", createdTool.Id.String(), false},
		{"empty tool ID", createdAsset.Id.String(), createdWorkOrder.Id.String(), "", false},
		{"empty asset and work order ID", "", "", createdTool.Id.String(), false},
		{"empty asset and tool ID", "", createdWorkOrder.Id.String(), "", false},
		{"empty work order and tool ID", createdAsset.Id.String(), "", "", false},
		{"empty asset, work order and tool ID", "", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateToolWithWorkOrder(tc.assetID, tc.workOrderID, tc.toolID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateToolWithWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateToolWithWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetTool(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetTool")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestGetTool: failed during setup. CreateTool() failed: %v", err)
	}

	testCases := []struct {
		name          string
		toolId        string
		shouldSucceed bool
	}{
		{"valid tool", createdTool.Id.String(), true},
		{"invalid tool ID", "invalid", false},
		{"nil tool ID", uuid.Nil.String(), false},
		{"empty tool ID", "", false},
		{"non-existent tool", uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetTool(tc.toolId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetTool() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetTool() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListTools(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListTools")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	tl := setupApiToolRequest(1)
	_, err = app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestListTools: failed during setup. CreateTool() failed: %v", err)
	}

	tl = setupApiToolRequest(2)
	_, err = app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestListTools: failed during setup. CreateTool() failed: %v", err)
	}

	testCases := []struct {
		name          string
		count         int
		shouldSucceed bool
	}{
		{"valid list", 2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := app.ListTools()
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListTools() failed: %v", err)
				} else {
					if len(ts) != tc.count {
						t.Errorf("ListTools() failed: expected %d tools, got %d", tc.count, len(ts))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListTools() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateTool(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateTool")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	toolCount := 5
	var ids []string
	tools := make(map[string]apitp.ToolResponse)
	nilIdTools := make(map[string]apitp.ToolRequest)
	for i := 0; i < toolCount; i++ {
		tl := setupApiToolRequest(i)
		ct, err := app.CreateTool(tl)
		if err != nil {
			t.Errorf("TestUpdateTool: failed during setup. CreateTool() failed: %v", err)
		}

		ids = append(ids, ct.Id.String())
		tools[ct.Id.String()] = ct
		nilIdTools[ct.Id.String()] = tl
	}

	testCases := []struct {
		name          string
		toolId        string
		tool          apitp.ToolRequest
		title         string
		shouldSucceed bool
	}{
		{"valid tool with matching IDs", ids[0], nilIdTools[ids[0]], "valid title1", true},
		{"valid tool with Tool.Id nil", ids[1], nilIdTools[ids[1]], "valid title2", true},
		{"mismatching tool ID and Tool.Id", ids[2], nilIdTools[ids[3]], "valid title3", false},
		{"non-existent tool", uuid.New().String(), apitp.ToolRequest{}, "valid title3", false},

		{"invalid tool ID", "invalid", apitp.ToolRequest{}, "valid title3", false},
		{"nil tool ID", uuid.Nil.String(), apitp.ToolRequest{}, "valid title3", false},
		{"empty tool ID", "", apitp.ToolRequest{}, "valid title3", false},
		{"conflicting id", ids[4], nilIdTools[ids[3]], "valid title3", false},

		{"empty title", ids[1], nilIdTools[ids[1]], "", false},
		{"minimum length title", ids[1], nilIdTools[ids[1]], strings.Repeat("a", apitp.MinEntityTitleLength), true},
		{"maximum length title", ids[1], nilIdTools[ids[1]], strings.Repeat("a", apitp.MaxEntityTitleLength), true},
		{"too long title", ids[1], nilIdTools[ids[1]], strings.Repeat("a", apitp.MaxEntityTitleLength+1), false},
		{"conflicting title", ids[2], nilIdTools[ids[2]], tools[ids[3]].Title, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.tool.Title = tc.title
			_, err := app.UpdateTool(tc.toolId, tc.tool)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateTool() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateTool() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateTool(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateTool")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name          string
		tool          apitp.ToolRequest
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid tool", setupApiToolRequest(1), uuid.New(), "valid title", true},
		{"nil id", setupApiToolRequest(2), uuid.Nil, "valid title", false},

		{"empty title", setupApiToolRequest(3), uuid.New(), "", false},
		{"minimum length title", setupApiToolRequest(4), uuid.New(), strings.Repeat("a", apitp.MinEntityTitleLength), true},
		{"maximum length title", setupApiToolRequest(5), uuid.New(), strings.Repeat("a", apitp.MaxEntityTitleLength), true},
		{"too long title", setupApiToolRequest(6), uuid.New(), strings.Repeat("a", apitp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.tool.Id = tc.id
			tc.tool.Title = tc.title
			err := app.validateTool(tc.tool)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateTool() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateTool() should have failed with %s", tc.name)
			}
		})
	}
}

func TestToolExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestToolExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	tl := setupApiToolRequest(1)
	createdTool, err := app.CreateTool(tl)
	if err != nil {
		t.Errorf("TestToolExists: failed during setup. CreateTool() failed: %v", err)
	}

	testCases := []struct {
		name          string
		toolId        string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid tool", createdTool.Id.String(), true, true},
		{"non-existent tool", uuid.New().String(), false, true},

		{"invalid tool ID", "invalid", false, false},
		{"nil tool ID", uuid.Nil.String(), false, false},
		{"empty tool ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.toolExists(tc.toolId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("toolExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("toolExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("toolExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}

func setupApiToolRequest(identifier int) apitp.ToolRequest {
	return apitp.ToolRequest{
		Title: "Tool " + strconv.Itoa(identifier),
	}
}

func setupApiToolSizeRequest(identifier int, toolId uuid.UUID) apitp.ToolSizeRequest {
	return apitp.ToolSizeRequest{
		Title: "Tool Size " + strconv.Itoa(identifier),
		Size:  utest.ToPtr("Size " + strconv.Itoa(identifier)),
	}
}
