package postgres_test

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupTimeTrigger(identifier int, taskId uuid.UUID) tp.TimeTrigger {
	return tp.TimeTrigger{
		Id:       uuid.New(),
		TaskId:   taskId,
		Quantity: identifier,
		TimeUnit: tp.TimeTriggerUnitDays,
	}
}

func TestTimeTriggerCreate(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTimeTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	tt := setupTimeTrigger(1, tk.Id)
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}

	utest.CompEntities(t, tt, createdTt)
}

func TestTimeTriggerDelete(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerdelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTimeTriggerDelete: failed during setup. CreateTask() failed: %v", err)
	}

	tt := setupTimeTrigger(1, tk.Id)
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("TestTimeTriggerDelete: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	err = store.DeleteTimeTrigger(createdTt.Id)
	if err != nil {
		t.Errorf("DeleteTimeTrigger() failed: %v", err)
	}

	_, err = store.GetTimeTrigger(createdTt.Id)
	if err == nil {
		t.Errorf("GetTimeTrigger() returned nil error after deletion")
	}
}

func TestTimeTriggerGet(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTimeTriggerGet: failed during setup. CreateTask() failed: %v", err)
	}

	tt := setupTimeTrigger(1, tk.Id)
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("TestTimeTriggerGet: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	getTt, err := store.GetTimeTrigger(createdTt.Id)
	if err != nil {
		t.Errorf("GetTimeTrigger() failed: %v", err)
	}

	utest.CompEntities(t, createdTt, getTt)
}

func TestTimeTriggerList(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerlist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTimeTriggerList: failed during setup. CreateTask() failed: %v", err)
	}

	tt1 := setupTimeTrigger(1, tk.Id)
	tt2 := setupTimeTrigger(2, tk.Id)
	tt3 := setupTimeTrigger(3, tk.Id)

	_, err = store.CreateTimeTrigger(tt1)
	if err != nil {
		t.Errorf("TestTimeTriggerList: failed during setup. CreateTimeTrigger() failed: %v", err)
	}
	_, err = store.CreateTimeTrigger(tt2)
	if err != nil {
		t.Errorf("TestTimeTriggerList: failed during setup. CreateTimeTrigger() failed: %v", err)
	}
	_, err = store.CreateTimeTrigger(tt3)
	if err != nil {
		t.Errorf("TestTimeTriggerList: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	tts, err := store.ListTimeTriggers()
	if err != nil {
		t.Errorf("ListTimeTriggers() failed: %v", err)
	}

	if len(tts) != 3 {
		t.Errorf("ListTimeTriggers() returned %d time triggers, expected 3", len(tts))
	}
}

func TestTimeTriggerUpdate(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerupdate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := setupTask(1)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTimeTriggerUpdate: failed during setup. CreateTask() failed: %v", err)
	}

	tt := setupTimeTrigger(1, tk.Id)
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("TestTimeTriggerUpdate: failed during setup. CreateTimeTrigger() failed: %v", err)
	}

	tt.Quantity = 60
	tt.TimeUnit = tp.TimeTriggerUnitWeeks
	updatedTt, err := store.UpdateTimeTrigger(tt)
	if err != nil {
		t.Errorf("UpdateTimeTrigger() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Quantity", "TimeUnit"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, updatedTt, createdTt, diffFields)
}
