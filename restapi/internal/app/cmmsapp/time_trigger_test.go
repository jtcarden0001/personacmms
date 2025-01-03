package cmmsapp

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateTimeTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateTimeTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestCreateTimeTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestCreateTimeTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	timeTrigger := tp.TimeTrigger{
		Quantity: 1,
		TimeUnit: "day",
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		timeTrigger   tp.TimeTrigger
		shouldSucceed bool
	}{
		{"valid time trigger", createdAsset.Id.String(), createdTask.Id.String(), timeTrigger, true},
		{"invalid asset ID", "invalid", createdTask.Id.String(), timeTrigger, false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", timeTrigger, false},
		{"invalid asset and task ID", "invalid", "invalid", timeTrigger, false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), timeTrigger, false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), timeTrigger, false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), timeTrigger, false},
		{"empty asset ID", "", createdTask.Id.String(), timeTrigger, false},
		{"empty task ID", createdAsset.Id.String(), "", timeTrigger, false},
		{"empty asset and task ID", "", "", timeTrigger, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateTimeTrigger(tc.assetID, tc.taskID, tc.timeTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateTimeTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateTimeTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteTimeTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteTimeTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteTimeTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDeleteTimeTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	tt := utest.SetupTimeTrigger(1, createdTask.Id, false)
	createdTimeTrigger, err := app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt)
	if err != nil {
		t.Errorf("TestDeleteTimeTrigger: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		timeTriggerID string
		shouldSucceed bool
	}{
		{"valid time trigger deletion", createdAsset.Id.String(), createdTask.Id.String(), createdTimeTrigger.Id.String(), true},
		{"invalid time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteTimeTrigger(tc.assetID, tc.taskID, tc.timeTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteTimeTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteTimeTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetTimeTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetTimeTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetTimeTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestGetTimeTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	tt := utest.SetupTimeTrigger(1, createdTask.Id, false)
	createdTimeTrigger, err := app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt)
	if err != nil {
		t.Errorf("TestGetTimeTrigger: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		timeTriggerID string
		shouldSucceed bool
	}{
		{"valid time trigger", createdAsset.Id.String(), createdTask.Id.String(), createdTimeTrigger.Id.String(), true},
		{"invalid time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
		{"non-existent time trigger", createdAsset.Id.String(), createdTask.Id.String(), uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetTimeTrigger(tc.assetID, tc.taskID, tc.timeTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetTimeTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetTimeTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListTimeTriggersByAssetAndTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListTimeTriggersByAssetAndTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListTimeTriggersByAssetAndTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestListTimeTriggersByAssetAndTask: failed during setup. CreateTask() failed: %v", err)
	}

	tt1 := utest.SetupTimeTrigger(1, createdTask.Id, false)
	_, err = app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt1)
	if err != nil {
		t.Errorf("TestListTimeTriggersByAssetAndTask: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	tt2 := utest.SetupTimeTrigger(2, createdTask.Id, false)
	_, err = app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt2)
	if err != nil {
		t.Errorf("TestListTimeTriggersByAssetAndTask: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		count         int
		shouldSucceed bool
	}{
		{"valid list", createdAsset.Id.String(), createdTask.Id.String(), 2, true},
		{"invalid asset ID", "invalid", createdTask.Id.String(), 0, false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", 0, false},
		{"invalid asset and task ID", "invalid", "invalid", 0, false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), 0, false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), 0, false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), 0, false},
		{"empty asset ID", "", createdTask.Id.String(), 0, false},
		{"empty task ID", createdAsset.Id.String(), "", 0, false},
		{"empty asset and task ID", "", "", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tts, err := app.ListTimeTriggersByAssetAndTask(tc.assetID, tc.taskID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListTimeTriggersByAssetAndTask() failed: %v", err)
				} else {
					if len(tts) != tc.count {
						t.Errorf("ListTimeTriggersByAssetAndTask() failed: expected %d time triggers, got %d", tc.count, len(tts))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListTimeTriggersByAssetAndTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateTimeTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateTimeTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUpdateTimeTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestUpdateTimeTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	tt := utest.SetupTimeTrigger(1, createdTask.Id, false)
	createdTimeTrigger, err := app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt)
	if err != nil {
		t.Errorf("TestUpdateTimeTrigger: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		timeTriggerID string
		timeTrigger   tp.TimeTrigger
		shouldSucceed bool
	}{
		{"valid time trigger update", createdAsset.Id.String(), createdTask.Id.String(), createdTimeTrigger.Id.String(), createdTimeTrigger, true},
		{"invalid time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", createdTimeTrigger, false},
		{"nil time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), createdTimeTrigger, false},
		{"empty time trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", createdTimeTrigger, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.UpdateTimeTrigger(tc.assetID, tc.taskID, tc.timeTriggerID, tc.timeTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateTimeTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateTimeTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateTimeTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateTimeTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestValidateTimeTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestValidateTimeTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		timeTrigger   tp.TimeTrigger
		id            uuid.UUID
		quantity      int
		timeUnit      string
		shouldSucceed bool
	}{
		{"valid time trigger", utest.SetupTimeTrigger(1, createdTask.Id, false), uuid.New(), 1, "day", true},
		{"nil id", utest.SetupTimeTrigger(2, createdTask.Id, false), uuid.Nil, 1, "day", false},
		{"invalid quantity", utest.SetupTimeTrigger(3, createdTask.Id, false), uuid.New(), 0, "day", false},
		{"invalid time unit", utest.SetupTimeTrigger(4, createdTask.Id, false), uuid.New(), 1, "invalid", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.timeTrigger.Id = tc.id
			tc.timeTrigger.Quantity = tc.quantity
			tc.timeTrigger.TimeUnit = tc.timeUnit
			err := app.validateTimeTrigger(tc.timeTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateTimeTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateTimeTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestTimeTriggerExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestTimeTriggerExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestTimeTriggerExists: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestTimeTriggerExists: failed during setup. CreateTask() failed: %v", err)
	}

	tt := utest.SetupTimeTrigger(1, createdTask.Id, false)
	createdTimeTrigger, err := app.CreateTimeTrigger(createdAsset.Id.String(), createdTask.Id.String(), tt)
	if err != nil {
		t.Errorf("TestTimeTriggerExists: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		timeTriggerID string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid time trigger", createdTimeTrigger.Id.String(), true, true},
		{"non-existent time trigger", uuid.New().String(), false, true},
		{"invalid time trigger ID", "invalid", false, false},
		{"nil time trigger ID", uuid.Nil.String(), false, false},
		{"empty time trigger ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.timeTriggerExists(tc.timeTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("timeTriggerExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("timeTriggerExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("timeTriggerExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
