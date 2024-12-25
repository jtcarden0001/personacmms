package integration

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestTimeTriggerCreate(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

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

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, tt, createdTt, fieldsToExclude)
}

func TestTimeTriggerDelete(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerdelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

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

func TestTimeTriggerDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerdeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteTimeTrigger(uuid.UUID{})
	if err == nil {
		t.Errorf("DeleteTimeTrigger() should have failed")
	}
}

func TestTimeTriggerList(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerlist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	tts, err := store.ListTimeTriggers()
	if err != nil {
		t.Errorf("ListTimeTriggers() failed: %v", err)
	}

	// make a map ttId -> tt
	ttMap := make(map[uuid.UUID]tp.TimeTrigger)
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
		utest.CompEntities(t, tt, ttMap[tt.Id])
	}
}

func TestTimeTriggerListByTaskId(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerlistbytaskid"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	tts, err := store.ListTimeTriggersByTaskId(uuid.UUID{})
	if err != nil {
		t.Errorf("ListTimeTriggersByTaskId() failed: %v", err)
	}

	// make a map ttId -> tt
	ttMap := make(map[uuid.UUID]tp.TimeTrigger)
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

	tts, err = store.ListTimeTriggersByTaskId(assetTaskId)
	if err != nil {
		t.Errorf("ListTimeTriggersByTaskId() failed: %v", err)
	}

	if len(tts) != len(ttMap) {
		t.Errorf("ListTimeTriggersByTaskId() failed: expected %d, got %d", len(ttMap), len(tts))
	}

	// compare
	for _, tt := range tts {
		utest.CompEntities(t, tt, ttMap[tt.Id])
	}
}

func TestTimeTriggerUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

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

	utest.CompEntities(t, createdTt, updatedTt)

	// get
	tt, err = store.GetTimeTrigger(updatedTt.Id)
	if err != nil {
		t.Errorf("GetTimeTrigger() failed: %v", err)
	}

	utest.CompEntities(t, updatedTt, tt)
}

func TestTimeTriggerUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testtimetriggerupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	tt := tp.TimeTrigger{
		Id:       uuid.UUID{},
		TaskId:   uuid.UUID{},
		Quantity: 30,
		TimeUnit: tp.TimeUnitDays,
	}
	_, err := store.UpdateTimeTrigger(tt.Id, tt)
	if err == nil {
		t.Errorf("UpdateTimeTrigger() should have failed")
	}
}
