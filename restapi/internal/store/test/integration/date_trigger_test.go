package integration

import (
	"testing"
	"time"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestDateTriggerCreate(t *testing.T) {
	var dbName = "testdatetriggercreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	dt := tp.DateTrigger{
		Date:   time.Now().AddDate(1, 0, 0),
		TaskId: assetTaskId,
	}
	createdDt, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, dt, createdDt, fieldsToExclude)
}

func TestDateTriggerDelete(t *testing.T) {
	var dbName = "testdatetriggerdelete"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	dt := tp.DateTrigger{
		Date:   time.Now().AddDate(1, 0, 0),
		TaskId: assetTaskId,
	}
	createdDt, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}

	err = store.DeleteDateTrigger(createdDt.Id)
	if err != nil {
		t.Errorf("DeleteDateTrigger() failed: %v", err)
	}

	_, err = store.GetDateTrigger(createdDt.Id)
	if err == nil {
		t.Errorf("GetDateTrigger() should have failed")
	}
}

func TestDateTriggerList(t *testing.T) {
	var dbName = "testdatetriggerlist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")

	// list
	dts, err := store.ListDateTriggers()
	if err != nil {
		t.Errorf("ListDateTriggers() failed: %v", err)
	}

	triggerMap := make(map[tp.UUID]tp.DateTrigger)
	for _, trigger := range dts {
		triggerMap[trigger.Id] = trigger
	}

	dt := tp.DateTrigger{
		Date:   time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Second),
		TaskId: assetTaskId,
	}

	createdDt, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}
	triggerMap[createdDt.Id] = createdDt

	dt.Date = time.Now().AddDate(2, 0, 0).UTC().Truncate(time.Second)
	createdDt, err = store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}
	triggerMap[createdDt.Id] = createdDt

	dts, err = store.ListDateTriggers()
	if err != nil {
		t.Errorf("ListDateTriggers() failed: %v", err)
	}

	if len(dts) != len(triggerMap) {
		t.Errorf("ListDateTriggers() failed: expected %v, got %v", len(triggerMap), len(dts))
	}

	for _, trigger := range dts {
		compEntities(t, trigger, triggerMap[trigger.Id])
	}
}

func TestDateTriggerUpdateGet(t *testing.T) {
	var dbName = "testdatetriggerupdateget"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	assetTaskId := setupTask(t, store, "1")
	dt := tp.DateTrigger{
		Date:   time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Second),
		TaskId: assetTaskId,
	}
	createdDt, err := store.CreateDateTrigger(dt)
	if err != nil {
		t.Errorf("CreateDateTrigger() failed: %v", err)
	}

	// update
	createdDt.Date = time.Now().AddDate(2, 0, 0).UTC().Truncate(time.Second)
	updatedDt, err := store.UpdateDateTrigger(createdDt.Id, createdDt)
	if err != nil {
		t.Errorf("UpdateDateTrigger() failed: %v", err)
	}

	fieldsShouldBeDifferent := convertToSet([]string{"Date"})
	compEntitiesFieldsShouldBeDifferent(t, createdDt, updatedDt, fieldsShouldBeDifferent)

	// get
	retrievedDt, err := store.GetDateTrigger(createdDt.Id)
	if err != nil {
		t.Errorf("GetDateTrigger() failed: %v", err)
	}

	compEntities(t, updatedDt, retrievedDt)
}
