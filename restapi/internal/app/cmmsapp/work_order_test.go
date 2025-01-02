package cmmsapp

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssociateWorkOrderWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateWorkOrderWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateWorkOrderWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, cAsset.Id, false)
	createdTask, err := app.CreateTask(cAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestAssociateWorkOrderWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	wo := utest.SetupWorkOrder(1, cAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestAssociateWorkOrderWithTask: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		workOrderID   string
		shouldSucceed bool
	}{
		{"valid association", cAsset.Id.String(), createdTask.Id.String(), createdWorkOrder.Id.String(), true},
		{"invalid work order ID", cAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"invalid task ID", cAsset.Id.String(), "invalid", createdWorkOrder.Id.String(), false},
		{"invalid asset ID", "invalid", createdTask.Id.String(), createdWorkOrder.Id.String(), false},
		{"invalid asset and task ID", "invalid", "invalid", createdWorkOrder.Id.String(), false},
		{"invalid asset and work order ID", "invalid", createdTask.Id.String(), "invalid", false},
		{"invalid task and work order ID", cAsset.Id.String(), "invalid", "invalid", false},
		{"invalid asset, task and work order ID", "invalid", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), createdWorkOrder.Id.String(), false},
		{"nil task ID", cAsset.Id.String(), uuid.Nil.String(), createdWorkOrder.Id.String(), false},
		{"nil work order ID", cAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), createdWorkOrder.Id.String(), false},
		{"nil asset and work order ID", uuid.Nil.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"nil task and work order ID", cAsset.Id.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"nil asset, task and work order ID", uuid.Nil.String(), uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdTask.Id.String(), createdWorkOrder.Id.String(), false},
		{"empty task ID", cAsset.Id.String(), "", createdWorkOrder.Id.String(), false},
		{"empty work order ID", cAsset.Id.String(), createdTask.Id.String(), "", false},
		{"empty asset and task ID", "", "", createdWorkOrder.Id.String(), false},
		{"empty asset and work order ID", "", createdTask.Id.String(), "", false},
		{"empty task and work order ID", cAsset.Id.String(), "", "", false},
		{"empty asset, task and work order ID", "", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.AssociateWorkOrderWithTask(tc.assetID, tc.taskID, tc.workOrderID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("AssociateWorkOrderWithTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateWorkOrderWithTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestCreateWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestCreateWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	workOrder := utest.SetupWorkOrder(1, cAsset.Id, false)
	workOrderWithId := utest.SetupWorkOrder(2, cAsset.Id, true)

	testCases := []struct {
		name          string
		assetID       string
		workOrder     tp.WorkOrder
		shouldSucceed bool
	}{
		{"valid work order", cAsset.Id.String(), workOrder, true},
		{"work order with ID", cAsset.Id.String(), workOrderWithId, false},
		{"invalid asset ID", "invalid", workOrder, false},
		{"nil asset ID", uuid.Nil.String(), workOrder, false},
		{"empty asset ID", "", workOrder, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateWorkOrder(tc.assetID, tc.workOrder)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo := utest.SetupWorkOrder(1, cAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestDeleteWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		workOrderID   string
		shouldSucceed bool
	}{
		{"valid work order deletion", cAsset.Id.String(), createdWorkOrder.Id.String(), true},
		{"invalid work order ID", cAsset.Id.String(), "invalid", false},
		{"nil work order ID", cAsset.Id.String(), uuid.Nil.String(), false},
		{"empty work order ID", cAsset.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteWorkOrder(tc.assetID, tc.workOrderID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateWorkOrderWithTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateWorkOrderWithTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateWorkOrderWithTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, cAsset.Id, false)
	createdTask, err := app.CreateTask(cAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDisassociateWorkOrderWithTask: failed during setup. CreateTask() failed: %v", err)
	}

	wo := utest.SetupWorkOrder(1, cAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestDisassociateWorkOrderWithTask: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	_, err = app.AssociateWorkOrderWithTask(cAsset.Id.String(), createdTask.Id.String(), createdWorkOrder.Id.String())
	if err != nil {
		t.Errorf("TestDisassociateWorkOrderWithTask: failed during setup. AssociateWorkOrderWithTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		workOrderID   string
		shouldSucceed bool
	}{
		{"valid disassociation", cAsset.Id.String(), createdTask.Id.String(), createdWorkOrder.Id.String(), true},
		{"invalid work order ID", cAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil work order ID", cAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty work order ID", cAsset.Id.String(), createdTask.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateWorkOrderWithTask(tc.assetID, tc.taskID, tc.workOrderID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateWorkOrderWithTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateWorkOrderWithTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	wo := utest.SetupWorkOrder(1, cAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
	if err != nil {
		t.Errorf("TestGetWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		workOrderID   string
		shouldSucceed bool
	}{
		{"valid work order", cAsset.Id.String(), createdWorkOrder.Id.String(), true},
		{"invalid work order ID", cAsset.Id.String(), "invalid", false},
		{"nil work order ID", cAsset.Id.String(), uuid.Nil.String(), false},
		{"empty work order ID", cAsset.Id.String(), "", false},
		{"non-existent work order", cAsset.Id.String(), uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetWorkOrder(tc.assetID, tc.workOrderID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListWorkOrdersByAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListWorkOrdersByAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListWorkOrdersByAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	wo1 := utest.SetupWorkOrder(1, cAsset.Id, false)
	_, err = app.CreateWorkOrder(cAsset.Id.String(), wo1)
	if err != nil {
		t.Errorf("TestListWorkOrdersByAsset: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	wo2 := utest.SetupWorkOrder(2, cAsset.Id, false)
	_, err = app.CreateWorkOrder(cAsset.Id.String(), wo2)
	if err != nil {
		t.Errorf("TestListWorkOrdersByAsset: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		count         int
		shouldSucceed bool
	}{
		{"valid list", cAsset.Id.String(), 2, true},
		{"invalid asset ID", "invalid", 0, false},
		{"nil asset ID", uuid.Nil.String(), 0, false},
		{"empty asset ID", "", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wos, err := app.ListWorkOrdersByAsset(tc.assetID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListWorkOrdersByAsset() failed: %v", err)
				} else {
					if len(wos) != tc.count {
						t.Errorf("ListWorkOrdersByAsset() failed: expected %d work orders, got %d", tc.count, len(wos))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListWorkOrdersByAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListWorkOrderStatus(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListWorkOrderStatus")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	statuses, err := app.ListWorkOrderStatus()
	if err != nil {
		t.Errorf("ListWorkOrderStatus() failed: %v", err)
	}

	if len(statuses) != len(tp.ValidWorkOrderStatuses) {
		t.Errorf("ListWorkOrderStatus() failed: expected %d statuses, got %d", len(tp.ValidWorkOrderStatuses), len(statuses))
	}
}

func TestUpdateWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUpdateWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	workOrderCount := 5
	var woIds []string
	createdWorkOrders := make(map[string]tp.WorkOrder)
	nilIdWorkOrders := make(map[string]tp.WorkOrder)
	for i := 0; i < workOrderCount; i++ {
		wo := utest.SetupWorkOrder(i, cAsset.Id, false)
		cwo, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
		if err != nil {
			t.Errorf("TestUpdateWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
		}

		woIds = append(woIds, cwo.Id.String())
		createdWorkOrders[cwo.Id.String()] = cwo
		nilIdWorkOrders[cwo.Id.String()] = wo
	}

	testCases := []struct {
		name          string
		assetID       string
		workOrderID   string
		workOrder     tp.WorkOrder
		title         string
		shouldSucceed bool
	}{
		{"valid work order with matching IDs", cAsset.Id.String(), woIds[0], createdWorkOrders[woIds[0]], "valid title1", true},
		{"valid work order with WorkOrder.Id nil", cAsset.Id.String(), woIds[1], nilIdWorkOrders[woIds[1]], "valid title2", true},
		{"mismatching work order ID and WorkOrder.Id", cAsset.Id.String(), woIds[2], createdWorkOrders[woIds[3]], "valid title3", false},
		{"non-existent work order", cAsset.Id.String(), uuid.New().String(), tp.WorkOrder{}, "valid title3", false},

		{"invalid work order ID", cAsset.Id.String(), "invalid", tp.WorkOrder{}, "valid title3", false},
		{"nil work order ID", cAsset.Id.String(), uuid.Nil.String(), tp.WorkOrder{}, "valid title3", false},
		{"empty work order ID", cAsset.Id.String(), "", tp.WorkOrder{}, "valid title3", false},
		{"conflicting id", cAsset.Id.String(), woIds[4], createdWorkOrders[woIds[3]], "valid title3", false},

		{"empty title", cAsset.Id.String(), woIds[1], createdWorkOrders[woIds[1]], "", false},
		{"minimum length title", cAsset.Id.String(), woIds[1], createdWorkOrders[woIds[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", cAsset.Id.String(), woIds[1], createdWorkOrders[woIds[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", cAsset.Id.String(), woIds[1], createdWorkOrders[woIds[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.workOrder.Title = tc.title
			uwo, err := app.UpdateWorkOrder(tc.assetID, tc.workOrderID, tc.workOrder)
			if tc.shouldSucceed && err != nil {

				t.Errorf("UpdateWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {

				t.Errorf("UpdateWorkOrder() should have failed with %s", tc.name)
			}

			if tc.shouldSucceed {
				if uwo.Title != tc.title {
					t.Errorf("UpdateWorkOrder() failed: expected title %s, got %s", tc.title, uwo.Title)
				}
			}
		})
	}
}

func TestValidateWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestValidateWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		id            uuid.UUID
		title         string
		createdDate   time.Time
		status        string
		assetId       uuid.UUID
		workOrder     tp.WorkOrder
		shouldSucceed bool
	}{
		{"valid work order", uuid.New(), "valid title 1", time.Now(), tp.WorkOrderStatusComplete, cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), true},
		{"invalid work order ID", uuid.Nil, "invalid title", time.Now(), "invalid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"empty title", uuid.New(), "", time.Now(), "valid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"too short title", uuid.New(), "a", time.Now(), "valid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"too long title", uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), time.Now(), "valid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"future created date", uuid.New(), "future date", time.Now().Add(24 * time.Hour), "valid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"invalid status", uuid.New(), "invalid status", time.Now(), "invalid status", cAsset.Id, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"invalid asset", uuid.New(), "invalid asset", time.Now(), "valid status", uuid.Nil, utest.SetupWorkOrder(1, cAsset.Id, false), false},
		{"non-existent asset", uuid.New(), "non-existent asset", time.Now(), "valid status", uuid.New(), utest.SetupWorkOrder(1, cAsset.Id, false), false},
	}

	for _, tc := range testCases {
		tc.workOrder.Id = tc.id
		tc.workOrder.Title = tc.title
		tc.workOrder.CreatedDate = tc.createdDate
		tc.workOrder.Status = tc.status
		tc.workOrder.AssetId = tc.assetId
		t.Run(tc.name, func(t *testing.T) {
			err := app.validateWorkOrder(tc.workOrder)
			if !tc.shouldSucceed && err == nil {
				t.Errorf("ValidateWorkOrder() should have failed with %s", tc.name)
			}

			if tc.shouldSucceed && err != nil {
				t.Errorf("ValidateWorkOrder() failed: %v", err)
			}
		})
	}
}

func TestWorkOrderExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestWorkOrderExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	cAsset, err := app.CreateAsset(a)
	if err != nil {

		t.Errorf("TestWorkOrderExists: failed during setup. CreateAsset() failed: %v", err)
	}

	wo := utest.SetupWorkOrder(1, cAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(cAsset.Id.String(), wo)
	if err != nil {

		t.Errorf("TestWorkOrderExists: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name        string
		workOrderID string
		shouldExist bool
	}{
		{"existing work order", createdWorkOrder.Id.String(), true},
		{"non-existent work order", uuid.New().String(), false},
		{"invalid work order ID", "invalid", false},
		{"nil work order ID", uuid.Nil.String(), false},
		{"empty work order ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.workOrderExists(tc.workOrderID)
			if tc.shouldExist && (err != nil || !exists) {
				t.Errorf("WorkOrderExists() failed: %v", err)
			}

			if !tc.shouldExist && (err == nil && exists) {
				t.Errorf("WorkOrderExists() should have failed with %s", tc.name)
			}
		})
	}
}
