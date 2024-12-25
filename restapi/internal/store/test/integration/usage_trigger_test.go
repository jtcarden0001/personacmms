package integration

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestUsageTriggerCreate(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	ut := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, ut, createdUt, fieldsToExclude)
}

func TestUsageTriggerDelete(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerdelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	ut := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	err = store.DeleteUsageTrigger(createdUt.Id)
	if err != nil {
		t.Errorf("DeleteUsageTrigger() failed: %v", err)
	}

	_, err = store.GetUsageTrigger(createdUt.Id)
	if err == nil {
		t.Errorf("GetUsageTrigger() should have failed")
	}
}

func TestUsageTriggerDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerdeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteUsageTrigger(uuid.New())
	if err == nil {
		t.Errorf("DeleteUsageTrigger() failed: expected error, got nil")
	}
}

func TestUsageTriggerList(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerlist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	uts, err := store.ListUsageTriggers()
	if err != nil {
		t.Errorf("ListUsageTriggers() failed: %v", err)
	}

	// make a map utId -> ut
	utMap := make(map[uuid.UUID]tp.UsageTrigger)
	for _, ut := range uts {
		utMap[ut.Id] = ut
	}

	// setup
	assetTaskId := setupTask(t, store, "1")
	ut := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	ut, err = store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}
	utMap[ut.Id] = ut

	ut2 := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  60,
		UsageUnit: tp.UsageUnitMiles,
	}
	ut2, err = store.CreateUsageTrigger(ut2)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}
	utMap[ut2.Id] = ut2

	uts, err = store.ListUsageTriggers()
	if err != nil {
		t.Errorf("ListUsageTriggers() failed: %v", err)
	}

	if len(uts) != len(utMap) {
		t.Errorf("ListUsageTriggers() failed: expected %d, got %d", len(utMap), len(uts))
	}

	// compare
	for _, ut := range uts {
		utest.CompEntities(t, ut, utMap[ut.Id])
	}
}

// TODO: add content checks
func TestUsageTriggerListByTaskId(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerlistbytaskid"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// create 2 usage triggers for task 1
	assetTaskId := setupTask(t, store, "1")
	ut := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	_, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	ut2 := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  60,
		UsageUnit: tp.UsageUnitMiles,
	}
	_, err = store.CreateUsageTrigger(ut2)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	// create 1 usage trigger for task 2
	assetTaskId2 := setupTask(t, store, "2")
	ut3 := tp.UsageTrigger{
		TaskId:    assetTaskId2,
		Quantity:  90,
		UsageUnit: tp.UsageUnitHours,
	}
	_, err = store.CreateUsageTrigger(ut3)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	uts, err := store.ListUsageTriggersByTaskId(assetTaskId)
	if err != nil {
		t.Errorf("ListUsageTriggersByTaskId() failed: %v", err)
	}

	if len(uts) != 2 {
		t.Errorf("ListUsageTriggersByTaskId() failed: expected 2, got %d", len(uts))
	}

	// list all usage triggers for task 2
	uts, err = store.ListUsageTriggersByTaskId(assetTaskId2)
	if err != nil {
		t.Errorf("ListUsageTriggersByTaskId() failed: %v", err)
	}

	if len(uts) != 1 {
		t.Errorf("ListUsageTriggersByTaskId() failed: expected 1, got %d", len(uts))
	}

	// list all usage triggers
	uts, err = store.ListUsageTriggers()
	if err != nil {
		t.Errorf("ListUsageTriggers() failed: %v", err)
	}

	if len(uts) != 3 {
		t.Errorf("ListUsageTriggers() failed: expected 3, got %d", len(uts))
	}
}

func TestUsageTriggerUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	ut := tp.UsageTrigger{
		TaskId:    assetTaskId,
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	// update
	createdUt.Quantity = 60
	createdUt.UsageUnit = tp.UsageUnitHours
	updatedUt, err := store.UpdateUsageTrigger(createdUt.Id, createdUt)
	if err != nil {
		t.Errorf("UpdateUsageTrigger() failed: %v", err)
	}

	utest.CompEntities(t, createdUt, updatedUt)

	// get
	ut, err = store.GetUsageTrigger(updatedUt.Id)
	if err != nil {
		t.Errorf("GetUsageTrigger() failed: %v", err)
	}

	utest.CompEntities(t, updatedUt, ut)
}

func TestUsageTriggerUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testusagetriggerupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	ut := tp.UsageTrigger{
		TaskId:    uuid.New(),
		Quantity:  30,
		UsageUnit: tp.UsageUnitDays,
	}
	_, err := store.UpdateUsageTrigger(uuid.New(), ut)
	if err == nil {
		t.Errorf("UpdateUsageTrigger() failed: expected error, got nil")
	}
}
