package cmmsapp

import (
	"testing"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateDateTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateDateTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestCreateDateTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestCreateDateTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	dateTrigger := tp.DateTrigger{
		ScheduledDate: time.Now().Add(24 * time.Hour),
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		dateTrigger   tp.DateTrigger
		shouldSucceed bool
	}{
		{"valid date trigger", createdAsset.Id.String(), createdTask.Id.String(), dateTrigger, true},
		{"invalid asset ID", "invalid", createdTask.Id.String(), dateTrigger, false},
		{"invalid task ID", createdAsset.Id.String(), "invalid", dateTrigger, false},
		{"invalid asset and task ID", "invalid", "invalid", dateTrigger, false},
		{"nil asset ID", uuid.Nil.String(), createdTask.Id.String(), dateTrigger, false},
		{"nil task ID", createdAsset.Id.String(), uuid.Nil.String(), dateTrigger, false},
		{"nil asset and task ID", uuid.Nil.String(), uuid.Nil.String(), dateTrigger, false},
		{"empty asset ID", "", createdTask.Id.String(), dateTrigger, false},
		{"empty task ID", createdAsset.Id.String(), "", dateTrigger, false},
		{"empty asset and task ID", "", "", dateTrigger, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateDateTrigger(tc.assetID, tc.taskID, tc.dateTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateDateTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateDateTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteDateTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteDateTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteDateTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDeleteDateTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	dt := utest.SetupDateTrigger(1, createdTask.Id, false)
	createdDateTrigger, err := app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt)
	if err != nil {
		t.Errorf("TestDeleteDateTrigger: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		dateTriggerID string
		shouldSucceed bool
	}{
		{"valid date trigger deletion", createdAsset.Id.String(), createdTask.Id.String(), createdDateTrigger.Id.String(), true},
		{"invalid date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteDateTrigger(tc.assetID, tc.taskID, tc.dateTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteDateTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteDateTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetDateTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetDateTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetDateTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestGetDateTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	dt := utest.SetupDateTrigger(1, createdTask.Id, false)
	createdDateTrigger, err := app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt)
	if err != nil {
		t.Errorf("TestGetDateTrigger: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		dateTriggerID string
		shouldSucceed bool
	}{
		{"valid date trigger", createdAsset.Id.String(), createdTask.Id.String(), createdDateTrigger.Id.String(), true},
		{"invalid date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", false},
		{"nil date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), false},
		{"empty date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", false},
		{"non-existent date trigger", createdAsset.Id.String(), createdTask.Id.String(), uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetDateTrigger(tc.assetID, tc.taskID, tc.dateTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetDateTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetDateTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListDateTriggersByAssetAndTask(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListDateTriggersByAssetAndTask")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListDateTriggersByAssetAndTask: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestListDateTriggersByAssetAndTask: failed during setup. CreateTask() failed: %v", err)
	}

	dt1 := utest.SetupDateTrigger(1, createdTask.Id, false)
	_, err = app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt1)
	if err != nil {
		t.Errorf("TestListDateTriggersByAssetAndTask: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	dt2 := utest.SetupDateTrigger(1, createdTask.Id, false)
	_, err = app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt2)
	if err != nil {
		t.Errorf("TestListDateTriggersByAssetAndTask: failed during setup. CreateDateTrigger() failed: %v", err)
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
			dts, err := app.ListDateTriggersByAssetAndTask(tc.assetID, tc.taskID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListDateTriggersByAssetAndTask() failed: %v", err)
				} else {
					if len(dts) != tc.count {
						t.Errorf("ListDateTriggersByAssetAndTask() failed: expected %d date triggers, got %d", tc.count, len(dts))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListDateTriggersByAssetAndTask() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateDateTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateDateTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestUpdateDateTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestUpdateDateTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	dt := utest.SetupDateTrigger(1, createdTask.Id, false)
	createdDateTrigger, err := app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt)
	if err != nil {
		t.Errorf("TestUpdateDateTrigger: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		taskID        string
		dateTriggerID string
		dateTrigger   tp.DateTrigger
		shouldSucceed bool
	}{
		{"valid date trigger update", createdAsset.Id.String(), createdTask.Id.String(), createdDateTrigger.Id.String(), createdDateTrigger, true},
		{"invalid date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "invalid", createdDateTrigger, false},
		{"nil date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), uuid.Nil.String(), createdDateTrigger, false},
		{"empty date trigger ID", createdAsset.Id.String(), createdTask.Id.String(), "", createdDateTrigger, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.UpdateDateTrigger(tc.assetID, tc.taskID, tc.dateTriggerID, tc.dateTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateDateTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateDateTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateDateTrigger(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateDateTrigger")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestValidateDateTrigger: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestValidateDateTrigger: failed during setup. CreateTask() failed: %v", err)
	}

	testCases := []struct {
		name          string
		dateTrigger   tp.DateTrigger
		id            uuid.UUID
		scheduledDate time.Time
		shouldSucceed bool
	}{
		{"valid date trigger", utest.SetupDateTrigger(1, createdTask.Id, false), uuid.New(), time.Now().Add(24 * time.Hour), true},
		{"nil id", utest.SetupDateTrigger(2, uuid.New(), false), uuid.Nil, time.Now().Add(24 * time.Hour), false},

		{"nil task ID", utest.SetupDateTrigger(2, uuid.New(), false), uuid.New(), time.Now().Add(24 * time.Hour), false},
		{"invalid task ID", utest.SetupDateTrigger(2, uuid.New(), false), uuid.New(), time.Now().Add(24 * time.Hour), false},
		{"non-existent task ID", utest.SetupDateTrigger(2, uuid.New(), false), uuid.New(), time.Now().Add(24 * time.Hour), false},

		{"nil scheduled date", utest.SetupDateTrigger(3, createdTask.Id, false), uuid.New(), time.Time{}, false},
		{"past scheduled date", utest.SetupDateTrigger(3, uuid.New(), false), uuid.New(), time.Now().Add(-24 * time.Hour), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.dateTrigger.Id = tc.id
			tc.dateTrigger.ScheduledDate = tc.scheduledDate
			err := app.validateDateTrigger(tc.dateTrigger)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateDateTrigger() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateDateTrigger() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDateTriggerExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDateTriggerExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDateTriggerExists: failed during setup. CreateAsset() failed: %v", err)
	}

	tk := utest.SetupTask(1, createdAsset.Id, false)
	createdTask, err := app.CreateTask(createdAsset.Id.String(), tk)
	if err != nil {
		t.Errorf("TestDateTriggerExists: failed during setup. CreateTask() failed: %v", err)
	}

	dt := utest.SetupDateTrigger(1, createdTask.Id, false)
	createdDateTrigger, err := app.CreateDateTrigger(createdAsset.Id.String(), createdTask.Id.String(), dt)
	if err != nil {
		t.Errorf("TestDateTriggerExists: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	testCases := []struct {
		name          string
		dateTriggerID string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid date trigger", createdDateTrigger.Id.String(), true, true},
		{"non-existent date trigger", uuid.New().String(), false, true},
		{"invalid date trigger ID", "invalid", false, false},
		{"nil date trigger ID", uuid.Nil.String(), false, false},
		{"empty date trigger ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.dateTriggerExists(tc.dateTriggerID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("dateTriggerExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("dateTriggerExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("dateTriggerExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
