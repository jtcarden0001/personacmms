package postgres_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/store"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestDateTriggerCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggercreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	tkId := setupTriggerDependencies(t, store)

	dt := utest.SetupDateTrigger(1, tkId, true)

	// test
	createdDateTrigger, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}

	utest.CompEntities(t, dt, createdDateTrigger)
}

func TestDateTriggerDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggerdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	tkId := setupTriggerDependencies(t, store)

	dt := utest.SetupDateTrigger(1, tkId, true)
	createdDateTrigger, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("TestDateTriggerDelete: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	// test
	err = store.DeleteDateTrigger(createdDateTrigger.Id)
	if err != nil {
		t.Errorf("TestDateTriggerDelete: DeleteDateTrigger() failed: %v", err)
	}

	_, err = store.GetDateTrigger(createdDateTrigger.Id)
	if err == nil {
		t.Errorf("TestDateTriggerDelete: GetDateTrigger() returned nil error after deletion")
	}
}

func TestDateTriggerGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggerget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	tkId := setupTriggerDependencies(t, store)

	dt := utest.SetupDateTrigger(1, tkId, true)
	createDateTrigger, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("TestDateTriggerGet: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	// test
	getDateTrigger, err := store.GetDateTrigger(createDateTrigger.Id)
	if err != nil {
		t.Errorf("GetDateTrigger() failed: %v", err)
	}

	utest.CompEntities(t, createDateTrigger, getDateTrigger)
}

func TestDateTriggerList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggerlist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	tkId := setupTriggerDependencies(t, store)

	dt1 := utest.SetupDateTrigger(1, tkId, true)
	dt2 := utest.SetupDateTrigger(2, tkId, true)
	dt3 := utest.SetupDateTrigger(3, tkId, true)

	_, err := store.CreateDateTrigger(dt1)
	if err != nil {
		t.Errorf("TestDateTriggerList: failed during setup. CreateDateTrigger() failed: %v", err)
	}
	_, err = store.CreateDateTrigger(dt2)
	if err != nil {
		t.Errorf("TestDateTriggerList: failed during setup. CreateDateTrigger() failed: %v", err)
	}
	_, err = store.CreateDateTrigger(dt3)
	if err != nil {
		t.Errorf("TestDateTriggerList: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	// test
	dateTriggers, err := store.ListDateTriggers()
	if err != nil {
		t.Errorf("ListDateTriggers() failed: %v", err)
	}

	if len(dateTriggers) != 3 {
		t.Errorf("ListDateTriggers() returned %d date triggers, expected 3", len(dateTriggers))
	}
}

func TestDateTriggerUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggerupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	tkId := setupTriggerDependencies(t, store)

	// test
	dt := utest.SetupDateTrigger(1, tkId, true)
	createdDateTrigger, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("TestDateTriggerUpdate: failed during setup. CreateDateTrigger() failed: %v", err)
	}

	// test
	dt.ScheduledDate = time.Now().AddDate(0, 2, 0)
	updatedDateTrigger, err := store.UpdateDateTrigger(dt)
	if err != nil {
		t.Errorf("UpdateDateTrigger() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"ScheduledDate"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdDateTrigger, updatedDateTrigger, differentFields)
}

func setupTriggerDependencies(t *testing.T, store store.Store) (taskId uuid.UUID) {
	aId := setupAssetInStore(t, store)
	tk := utest.SetupTask(1, aId, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Fatalf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	return tk.Id
}
