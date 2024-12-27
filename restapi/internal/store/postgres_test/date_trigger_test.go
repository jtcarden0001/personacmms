package postgres_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupDateTrigger(identifier int, taskId uuid.UUID) tp.DateTrigger {
	return tp.DateTrigger{
		Id:            uuid.New(),
		ScheduledDate: time.Now().AddDate(0, identifier, 0),
		TaskId:        taskId,
	}
}

func TestDateTriggerCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testdatetriggercreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	dt := setupDateTrigger(1, tk.Id)

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

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	dt := setupDateTrigger(1, tk.Id)
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

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	dt := setupDateTrigger(1, tk.Id)
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

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	dt1 := setupDateTrigger(1, tk.Id)
	dt2 := setupDateTrigger(2, tk.Id)
	dt3 := setupDateTrigger(3, tk.Id)

	_, err = store.CreateDateTrigger(dt1)
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

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestDateTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	dt := setupDateTrigger(1, tk.Id)
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
