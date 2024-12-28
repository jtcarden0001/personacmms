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

	tk := utest.SetupTask(1, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestUsageTriggerCreate: failed during setup. CreateTask() failed: %v", err)
	}

	ut := utest.SetupUsageTrigger(1, tk.Id, true)
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

	tk := utest.SetupTask(1, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestUsageTriggerDelete: failed during setup. CreateTask() failed: %v", err)
	}

	ut := utest.SetupUsageTrigger(1, tk.Id, true)
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

func TestUsageTriggerGet(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := utest.SetupTask(1, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestUsageTriggerGet: failed during setup. CreateTask() failed: %v", err)
	}

	ut := utest.SetupUsageTrigger(1, tk.Id, true)
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

	tk := utest.SetupTask(1, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestUsageTriggerList: failed during setup. CreateTask() failed: %v", err)
	}

	ut1 := utest.SetupUsageTrigger(1, tk.Id, true)
	ut2 := utest.SetupUsageTrigger(2, tk.Id, true)
	ut3 := utest.SetupUsageTrigger(3, tk.Id, true)

	_, err = store.CreateUsageTrigger(ut1)
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

func TestUsageTriggerUpdate(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerupdate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tk := utest.SetupTask(1, true)
	_, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestUsageTriggerUpdate: failed during setup. CreateTask() failed: %v", err)
	}

	ut := utest.SetupUsageTrigger(1, tk.Id, true)
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
