package postgres_test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestTimeTriggerCreate(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)
	tt := utest.SetupTimeTrigger(1, tkId, true)
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
	tkId := setupTriggerDependencies(t, store)

	tt := utest.SetupTimeTrigger(1, tkId, true)
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
	tkId := setupTriggerDependencies(t, store)
	tt := utest.SetupTimeTrigger(1, tkId, true)
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
	tkId := setupTriggerDependencies(t, store)

	tt1 := utest.SetupTimeTrigger(1, tkId, true)
	tt2 := utest.SetupTimeTrigger(2, tkId, true)
	tt3 := utest.SetupTimeTrigger(3, tkId, true)

	_, err := store.CreateTimeTrigger(tt1)
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
	tkId := setupTriggerDependencies(t, store)

	tt := utest.SetupTimeTrigger(1, tkId, true)
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
