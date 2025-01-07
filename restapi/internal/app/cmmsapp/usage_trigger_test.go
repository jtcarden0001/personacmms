package cmmsapp

import (
	"testing"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
)

func TestCreateUsageTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateUsageTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestCreateUsageTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestCreateUsageTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	usageTrigger := apitp.UsageTriggerRequest{
		Quantity:  1,
		UsageUnit: "hour",
	}

	usageTriggerWithId := apitp.UsageTriggerRequest{
		Id:        uuid.New(),
		Quantity:  1,
		UsageUnit: "hour",
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		usageTrigger  apitp.UsageTriggerRequest
		shouldSucceed bool
	}{
		{"valid usage trigger", createdAsset.Id.String(), createdTask.Id.String(), usageTrigger, true},
		{"valid usage trigger but with Id", createdAsset.Id.String(), createdTask.Id.String(), usageTriggerWithId, false},
		{"invalid asset ID", "invalid", createdTask.Id.String(), usageTrigger, false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", usageTrigger, false},
		{"invalid asset and task ID", "invalid", "invalid", usageTrigger, false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), usageTrigger, false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), usageTrigger, false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), usageTrigger, false},
		{"empty asset ID", "", createdTask.Id.String(), usageTrigger, false},
		{"empty task ID", createdAsset.Id.String(), "", usageTrigger, false},
		{"empty asset and task ID", "", "", usageTrigger, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateUsageTrigger(tc.assetID, tc.taskID, tc.usageTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateUsageTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateUsageTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteUsageTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteUsageTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteUsageTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDeleteUsageTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	ut := setupApiUsageTriggerRequest(1, createdTask.Id)
	createdUsageTrigger, err := app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut)
	if err != nil {
		t.Errorf("TestDeleteUsageTrigger: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	testCases := []struct {
		name           string
		assetID        string
		taskID         string
		usageTriggerID string
		shouldSucceed  bool
	}{
		{"valid usage trigger deletion", createdAsset.Id.String(), createdTask.Id.String(), createdUsageTrigger.Id.String(), true},
		{"invalid usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteUsageTrigger(tc.assetID, tc.taskID, tc.usageTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteUsageTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteUsageTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetUsageTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetUsageTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetUsageTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestGetUsageTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	ut := setupApiUsageTriggerRequest(1, createdTask.Id)
	createdUsageTrigger, err := app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut)
	if err != nil {
		t.Errorf("TestGetUsageTrigger: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	testCases := []struct {
		name           string
		assetID        string
		taskID         string
		usageTriggerID string
		shouldSucceed  bool
	}{
		{"valid usage trigger", createdAsset.Id.String(), createdTask.Id.String(), createdUsageTrigger.Id.String(), true},
		{"invalid usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
		{"non-existent usage trigger", createdAsset.Id.String(), createdTask.Id.String(), uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetUsageTrigger(tc.assetID, tc.taskID, tc.usageTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetUsageTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetUsageTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListUsageTriggersByAssetAndTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListUsageTriggersByAssetAndTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListUsageTriggersByAssetAndTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestListUsageTriggersByAssetAndTask: failed during setup. CreateTask() failed: %v", err)
	}

	ut1 := setupApiUsageTriggerRequest(1, createdTask.Id)
	_, err = app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut1)
	if err != nil {
		t.Errorf("TestListUsageTriggersByAssetAndTask: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	ut2 := setupApiUsageTriggerRequest(2, createdTask.Id)
	_, err = app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut2)
	if err != nil {
		t.Errorf("TestListUsageTriggersByAssetAndTask: failed during setup. CreateUsageTrigger() failed: %v", err)
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
			uts, err := app.ListUsageTriggersByAssetAndTask(tc.assetID, tc.taskID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListUsageTriggersByAssetAndTask() failed: %v", err)
				} else {
					if len(uts) != tc.count {
						t.Errorf("ListUsageTriggersByAssetAndTask() failed: expected %d usage triggers, got %d", tc.count, len(uts))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListUsageTriggersByAssetAndTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateUsageTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateUsageTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUpdateUsageTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestUpdateUsageTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	ut := setupApiUsageTriggerRequest(1, createdTask.Id)
	createdUsageTrigger, err := app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut)
	if err != nil {
		t.Errorf("TestUpdateUsageTrigger: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	testCases := []struct {
		name           string
		assetID        string
		taskID         string
		usageTriggerID string
		usageTrigger   apitp.UsageTriggerRequest
		shouldSucceed  bool
	}{
		{"valid usage trigger update", createdAsset.Id.String(), createdTask.Id.String(), createdUsageTrigger.Id.String(), ut, true},
		{"invalid usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", ut, false},
		{"nil usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), ut, false},
		{"empty usage trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", ut, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.UpdateUsageTrigger(tc.assetID, tc.taskID, tc.usageTriggerID, tc.usageTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateUsageTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateUsageTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateUsageTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateUsageTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestValidateUsageTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestValidateUsageTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		usageTrigger  apitp.UsageTriggerRequest
		id            uuid.UUID
		quantity      int
		usageUnit     string
		shouldSucceed bool
	}{
		{"valid usage trigger", setupApiUsageTriggerRequest(1, createdTask.Id), uuid.New(), 1, "hour", true},
		{"nil id", setupApiUsageTriggerRequest(2, createdTask.Id), uuid.Nil, 1, "hour", false},
		{"invalid quantity", setupApiUsageTriggerRequest(3, createdTask.Id), uuid.New(), 0, "hour", false},
		{"invalid usage unit", setupApiUsageTriggerRequest(4, createdTask.Id), uuid.New(), 1, "invalid", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.usageTrigger.Id = tc.id
			tc.usageTrigger.Quantity = tc.quantity
			tc.usageTrigger.UsageUnit = tc.usageUnit
			err := app.validateUsageTrigger(tc.usageTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateUsageTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateUsageTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListUsageTriggerUnits(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListUsageTriggerUnits")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	units, err := app.ListUsageTriggerUnits()
	if err != nil {
		t.Errorf("ListUsageTriggerUnits() failed: %v", err)
	}

	if len(units) != len(apitp.ValidUsageTriggerUnits) {
		t.Errorf("ListUsageTriggerUnits() failed: expected %d units, got %d", len(tp.ValidUsageTriggerUnits), len(units))
	}
}

func TestUsageTriggerExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUsageTriggerExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUsageTriggerExists: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := setupApiTaskRequest(1, createdAsset.Id)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestUsageTriggerExists: failed during setup. CreateTask() failed: %v", err)
	}

	ut := setupApiUsageTriggerRequest(1, createdTask.Id)
	createdUsageTrigger, err := app.CreateUsageTrigger(createdAsset.Id.String(), createdTask.Id.String(), ut)
	if err != nil {
		t.Errorf("TestUsageTriggerExists: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	testCases := []struct {
		name           string
		usageTriggerID string
		shouldExist    bool
		shouldSucceed  bool
	}{
		{"valid usage trigger", createdUsageTrigger.Id.String(), true, true},
		{"non-existent usage trigger", uuid.New().String(), false, true},
		{"invalid usage trigger ID", "invalid", false, false},
		{"nil usage trigger ID", uuid.Nil.String(), false, false},
		{"empty usage trigger ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.usageTriggerExists(tc.usageTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("usageTriggerExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("usageTriggerExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("usageTriggerExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}

func setupApiUsageTriggerRequest(identifier int, taskId uuid.UUID) apitp.UsageTriggerRequest {
	return apitp.UsageTriggerRequest{
		TaskId:    taskId,
		Quantity:  identifier,
		UsageUnit: tp.UsageTriggerUnitDays,
	}
}
