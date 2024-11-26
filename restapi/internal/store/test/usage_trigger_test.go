package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestUsageTriggerCreate(t *testing.T) {
	store := initializeStore("testusagetriggercreate")

	// setup
	assetTaskId := setupAssetTask(t, store, "1")
	ut := tp.UsageTrigger{
		AssetTaskId: assetTaskId,
		Quantity:    30,
		UsageUnit:   tp.UsageUnitDays,
	}
	createdUt, err := store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, ut, createdUt, fieldsToExclude)
}

func TestUsageTriggerDelete(t *testing.T) {
	store := initializeStore("testusagetriggerdelete")

	// setup
	assetTaskId := setupAssetTask(t, store, "1")
	ut := tp.UsageTrigger{
		AssetTaskId: assetTaskId,
		Quantity:    30,
		UsageUnit:   tp.UsageUnitDays,
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

func TestUsageTriggerList(t *testing.T) {
	store := initializeStore("testusagetriggerlist")

	// List
	uts, err := store.ListUsageTriggers()
	if err != nil {
		t.Errorf("ListUsageTriggers() failed: %v", err)
	}

	// make a map utId -> ut
	utMap := make(map[tp.UUID]tp.UsageTrigger)
	for _, ut := range uts {
		utMap[ut.Id] = ut
	}

	// setup
	assetTaskId := setupAssetTask(t, store, "1")
	ut := tp.UsageTrigger{
		AssetTaskId: assetTaskId,
		Quantity:    30,
		UsageUnit:   tp.UsageUnitDays,
	}
	ut, err = store.CreateUsageTrigger(ut)
	if err != nil {
		t.Errorf("CreateUsageTrigger() failed: %v", err)
	}
	utMap[ut.Id] = ut

	ut2 := tp.UsageTrigger{
		AssetTaskId: assetTaskId,
		Quantity:    60,
		UsageUnit:   tp.UsageUnitMiles,
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
		compEntities(t, ut, utMap[ut.Id])
	}
}

func TestUsageTriggerUpdateGet(t *testing.T) {
	store := initializeStore("testusagetriggerupdateget")

	// setup
	assetTaskId := setupAssetTask(t, store, "1")
	ut := tp.UsageTrigger{
		AssetTaskId: assetTaskId,
		Quantity:    30,
		UsageUnit:   tp.UsageUnitDays,
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

	compEntities(t, createdUt, updatedUt)

	// get
	ut, err = store.GetUsageTrigger(updatedUt.Id)
	if err != nil {
		t.Errorf("GetUsageTrigger() failed: %v", err)
	}

	compEntities(t, updatedUt, ut)
}
