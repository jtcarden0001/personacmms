package cmmsapp

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestCreateTask: failed during setup. CreateAsset() failed: %v", err)
	}

	conflictingTask := utest.SetupTask(1, createdAsset.Id, false)
	_, err = app.CreateTask(createdAsset.Id.String(), conflictingTask)
	if err != nil {
		t.Errorf("TestCreateTask: failed during setup. CreateTask() failed: %v", err)
	}

	emptyTitleTask := utest.SetupTask(2, createdAsset.Id, false)
	emptyTitleTask.Title = ""

	testCases := []struct {
		name          string
		assetID       string
		task          tp.Task
		shouldSucceed bool
	}{
		{"valid task", createdAsset.Id.String(), utest.SetupTask(3, createdAsset.Id, false), true},
		{"non nil id", createdAsset.Id.String(), utest.SetupTask(4, createdAsset.Id, true), false},
		{"empty title", createdAsset.Id.String(), emptyTitleTask, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateTask(tc.assetID, tc.task)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteTask: failed during setup. CreateAsset() failed: %v", err)
	}

	ta := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), ta)
	if err != nil {
		t.Errorf("TestDeleteTask: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		shouldSucceed bool
	}{
		{"valid task deletion", createdAsset.Id.String(), createdTask.Id.String(), true},
		{"invalid task ID", createdAsset.Id.String(), "invalid", false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"empty task ID", createdAsset.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteTask(tc.assetID, tc.taskID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateTaskWithWorkOrder(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateTaskWithWorkOrder")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateTaskWithWorkOrder: failed during setup. CreateAsset() failed: %v", err)
	}

	ta := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), ta)
	if err != nil {
		t.Errorf("TestDisassociateTaskWithWorkOrder: failed during setup. CreateTask() failed: %v", err)
	}

	w := utest.SetupWorkOrder(1, createdAsset.Id, false)
	createdWorkOrder, err := app.CreateWorkOrder(createdAsset.Id.String(), w)
	if err != nil {
		t.Errorf("TestDisassociateTaskWithWorkOrder: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	_, err = app.AssociateWorkOrderWithTask(createdAsset.Id.String(), createdTask.Id.String(), createdWorkOrder.Id.String())
	if err != nil {
		t.Errorf("TestDisassociateTaskWithWorkOrder: failed during setup. AssociateTaskWithWorkOrder() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		workOrderID   string
		shouldSucceed bool
	}{
		{"valid disassociation", createdAsset.Id.String(), createdTask.Id.String(), createdWorkOrder.Id.String(), true},
		{"invalid work order ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil work order ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty work order ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateTaskWithWorkOrder(tc.assetID, tc.taskID, tc.workOrderID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateTaskWithWorkOrder() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateTaskWithWorkOrder() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetTask: failed during setup. CreateAsset() failed: %v", err)
	}

	ta := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), ta)
	if err != nil {
		t.Errorf("TestGetTask: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		shouldSucceed bool
	}{
		{"valid task", createdAsset.Id.String(), createdTask.Id.String(), true},
		{"invalid task ID", createdAsset.Id.String(), "invalid", false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"empty task ID", createdAsset.Id.String(), "", false},
		{"non-existent task", createdAsset.Id.String(), uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetTask(tc.assetID, tc.taskID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListTasksByAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListTasksByAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListTasksByAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	t1 := utest.SetupTask(1, createdAsset.Id, false)
	_, err = app.CreateTask(createdAsset.Id.String(), t1)
	if err != nil {
		t.Errorf("TestListTasksByAsset: failed during setup. CreateTask() failed: %v", err)
	}

	t2 := utest.SetupTask(2, createdAsset.Id, false)
	_, err = app.CreateTask(createdAsset.Id.String(), t2)
	if err != nil {
		t.Errorf("TestListTasksByAsset: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		count         int
		shouldSucceed bool
	}{
		{"valid list", createdAsset.Id.String(), 2, true},
		{"invalid asset ID", "invalid", 0, false},
		{"nil asset ID", uuid.Nil.String(), 0, false},
		{"empty asset ID", "", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := app.ListTasksByAsset(tc.assetID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListTasksByAsset() failed: %v", err)
				} else {
					if len(ts) != tc.count {
						t.Errorf("ListTasksByAsset() failed: expected %d tasks, got %d", tc.count, len(ts))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListTasksByAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUpdateTask: failed during setup. CreateAsset() failed: %v", err)
	}

	taskCount := 5
	var ids []string
	tasks := make(map[string]tp.Task)
	nilIdTasks := make(map[string]tp.Task)
	for i := 0; i < taskCount; i++ {
		ta := utest.SetupTask(i, createdAsset.Id, false)
		ct, err := app.CreateTask(createdAsset.Id.String(), ta)
		if err != nil {
			t.Errorf("TestUpdateTask: failed during setup. CreateTask() failed: %v", err)
		}

		ids = append(ids, ct.Id.String())
		tasks[ct.Id.String()] = ct
		nilIdTasks[ct.Id.String()] = ta
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		task          tp.Task
		title         string
		shouldSucceed bool
	}{
		{"valid task with matching IDs", createdAsset.Id.String(), ids[0], tasks[ids[0]], "valid title1", true},
		{"valid task with Task.Id nil", createdAsset.Id.String(), ids[1], nilIdTasks[ids[1]], "valid title2", true},
		{"mismatching task ID and Task.Id", createdAsset.Id.String(), ids[2], tasks[ids[3]], "valid title3", false},
		{"non-existent task", createdAsset.Id.String(), uuid.New().String(), tp.Task{}, "valid title3", false},

		{"invalid task ID", createdAsset.Id.String(), "invalid", tp.Task{}, "valid title3", false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), tp.Task{}, "valid title3", false},
		{"empty task ID", createdAsset.Id.String(), "", tp.Task{}, "valid title3", false},
		{"conflicting id", createdAsset.Id.String(), ids[4], tasks[ids[3]], "valid title3", false},

		{"empty title", createdAsset.Id.String(), ids[1], tasks[ids[1]], "", false},
		{"minimum length title", createdAsset.Id.String(), ids[1], tasks[ids[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", createdAsset.Id.String(), ids[1], tasks[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", createdAsset.Id.String(), ids[1], tasks[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.task.Title = tc.title
			_, err := app.UpdateTask(tc.assetID, tc.taskID, tc.task)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestValidateTask: failed during setup. CreateAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		task          tp.Task
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid task", utest.SetupTask(1, createdAsset.Id, false), uuid.New(), "valid title", true},
		{"nil id", utest.SetupTask(2, createdAsset.Id, false), uuid.Nil, "valid title", false},

		{"empty title", utest.SetupTask(3, createdAsset.Id, false), uuid.New(), "", false},
		{"minimum length title", utest.SetupTask(4, createdAsset.Id, false), uuid.New(), strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", utest.SetupTask(5, createdAsset.Id, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", utest.SetupTask(6, createdAsset.Id, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.task.Id = tc.id
			tc.task.Title = tc.title
			err := app.validateTask(tc.task)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateTask() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestTaskExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestTaskExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestTaskExists: failed during setup. CreateAsset() failed: %v", err)
	}

	ta := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), ta)
	if err != nil {
		t.Errorf("TestTaskExists: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		taskID        string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid task", createdTask.Id.String(), true, true},
		{"non-existent task", uuid.New().String(), false, true},

		{"invalid task ID", "invalid", false, false},
		{"nil task ID", uuid.Nil.String(), false, false},
		{"empty task ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.taskExists(tc.taskID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("taskExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("taskExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("taskExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
