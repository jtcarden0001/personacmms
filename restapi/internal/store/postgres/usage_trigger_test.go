package postgres_test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestUsageTriggerCreate(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)

	ut := utest.SetupUsageTrigger(1, tkId, true)
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	utest.CompEntities(t, ut, createdUt)
}

func TestUsageTriggerDelete(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerdelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)

	ut := utest.SetupUsageTrigger(1, tkId, true)
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("TestUsageTriggerDelete: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	err = store.DeleteUsageTrigger(createdUt.Id)
	if err != nil {
		t.Errorf("DeleteUsageTrigger() failed: %v", err)
	}

	_, err = store.GetUsageTrigger(createdUt.Id)
	if err == nil {
		t.Errorf("GetUsageTrigger() returned nil error after deletion")
	}
}

// TODO: TestUsageTriggerDeleteFromTask

func TestUsageTriggerGet(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)

	ut := utest.SetupUsageTrigger(1, tkId, true)
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("TestUsageTriggerGet: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	getUt, err := store.GetUsageTrigger(createdUt.Id)
	if err != nil {
		t.Errorf("GetUsageTrigger() failed: %v", err)
	}

	utest.CompEntities(t, createdUt, getUt)
}

func TestUsageTriggerList(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerlist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)

	ut1 := utest.SetupUsageTrigger(1, tkId, true)
	ut2 := utest.SetupUsageTrigger(2, tkId, true)
	ut3 := utest.SetupUsageTrigger(3, tkId, true)

	_, err := store.CreateUsageTrigger(ut1)
	if err != nil {
		t.Errorf("TestUsageTriggerList: failed during setup. CreateUsageTrigger() failed: %v", err)
	}
	_, err = store.CreateUsageTrigger(ut2)
	if err != nil {
		t.Errorf("TestUsageTriggerList: failed during setup. CreateUsageTrigger() failed: %v", err)
	}
	_, err = store.CreateUsageTrigger(ut3)
	if err != nil {
		t.Errorf("TestUsageTriggerList: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	uts, err := store.ListUsageTriggers()
	if err != nil {
		t.Errorf("ListUsageTriggers() failed: %v", err)
	}

	if len(uts) != 3 {
		t.Errorf("ListUsageTriggers() returned %d usage triggers, expected 3", len(uts))
	}
}

// TODO: TestUsageTriggerListByTask

func TestUsageTriggerUpdate(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerupdate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)
	tkId := setupTriggerDependencies(t, store)

	ut := utest.SetupUsageTrigger(1, tkId, true)
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("TestUsageTriggerUpdate: failed during setup. CreateUsageTrigger() failed: %v", err)
	}

	ut.Quantity = 60
	ut.UsageUnit = tp.UsageTriggerUnitHours
	updatedUt, err := store.UpdateUsageTrigger(ut)
	if err != nil {
		t.Errorf("UpdateUsageTrigger() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Quantity", "UsageUnit"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdUt, updatedUt, diffFields)
}
