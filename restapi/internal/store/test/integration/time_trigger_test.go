package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTimeTriggerCreate(t *testing.T) {
	dbName := "testtimetriggercreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	tt := tp.TimeTrigger{
		TaskId:   assetTaskId,
		Quantity: 30,
		TimeUnit: tp.TimeUnitDays,
	}
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, tt, createdTt, fieldsToExclude)
}

func TestTimeTriggerDelete(t *testing.T) {
	dbName := "testtimetriggerdelete"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	tt := tp.TimeTrigger{
		TaskId:   assetTaskId,
		Quantity: 30,
		TimeUnit: tp.TimeUnitDays,
	}
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}

	err = store.DeleteTimeTrigger(createdTt.Id)
	if err != nil {
		t.Errorf("DeleteTimeTrigger() failed: %v", err)
	}

	_, err = store.GetTimeTrigger(createdTt.Id)
	if err == nil {
		t.Errorf("GetTimeTrigger() should have failed")
	}
}

func TestTimeTriggerList(t *testing.T) {
	dbName := "testtimetriggerlist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// List
	tts, err := store.ListTimeTriggers()
	if err != nil {
		t.Errorf("ListTimeTriggers() failed: %v", err)
	}

	// make a map ttId -> tt
	ttMap := make(map[tp.UUID]tp.TimeTrigger)
	for _, tt := range tts {
		ttMap[tt.Id] = tt
	}

	// setup
	assetTaskId := setupTask(t, store, "1")
	tt := tp.TimeTrigger{
		TaskId:   assetTaskId,
		Quantity: 30,
		TimeUnit: tp.TimeUnitDays,
	}
	tt, err = store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}
	ttMap[tt.Id] = tt

	tt2 := tp.TimeTrigger{
		TaskId:   assetTaskId,
		Quantity: 60,
		TimeUnit: tp.TimeUnitWeeks,
	}
	tt2, err = store.CreateTimeTrigger(tt2)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}
	ttMap[tt2.Id] = tt2

	tts, err = store.ListTimeTriggers()
	if err != nil {
		t.Errorf("ListTimeTriggers() failed: %v", err)
	}

	if len(tts) != len(ttMap) {
		t.Errorf("ListTimeTriggers() failed: expected %d, got %d", len(ttMap), len(tts))
	}

	// compare
	for _, tt := range tts {
		compEntities(t, tt, ttMap[tt.Id])
	}
}

func TestTimeTriggerUpdateGet(t *testing.T) {
	dbName := "testtimetriggerupdateget"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	tt := tp.TimeTrigger{
		TaskId:   assetTaskId,
		Quantity: 30,
		TimeUnit: tp.TimeUnitDays,
	}
	createdTt, err := store.CreateTimeTrigger(tt)
	if err != nil {
		t.Errorf("CreateTimeTrigger() failed: %v", err)
	}

	// update
	createdTt.Quantity = 60
	createdTt.TimeUnit = tp.TimeUnitWeeks
	updatedTt, err := store.UpdateTimeTrigger(createdTt.Id, createdTt)
	if err != nil {
		t.Errorf("UpdateTimeTrigger() failed: %v", err)
	}

	compEntities(t, createdTt, updatedTt)

	// get
	tt, err = store.GetTimeTrigger(updatedTt.Id)
	if err != nil {
		t.Errorf("GetTimeTrigger() failed: %v", err)
	}

	compEntities(t, updatedTt, tt)
}
