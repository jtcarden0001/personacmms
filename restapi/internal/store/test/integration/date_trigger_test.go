package integration

// import (
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
// )

// func TestDateTriggerCreate(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggercreate"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// setup
// 	assetTaskId := setupTask(t, store, "1")
// 	dt := tp.DateTrigger{
// 		Date:   time.Now().AddDate(1, 0, 0),
// 		TaskId: assetTaskId,
// 	}
// 	createdDt, err := store.CreateDateTrigger(dt)
// 	if err != nil {
// 		t.Errorf("CreateDateTrigger() failed: %v", err)
// 	}

// 	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
// 	utest.CompEntitiesExcludeFields(t, dt, createdDt, fieldsToExclude)
// }

// func TestDateTriggerDelete(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggerdelete"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// setup
// 	assetTaskId := setupTask(t, store, "1")
// 	dt := tp.DateTrigger{
// 		Date:   time.Now().AddDate(1, 0, 0),
// 		TaskId: assetTaskId,
// 	}
// 	createdDt, err := store.CreateDateTrigger(dt)
// 	if err != nil {
// 		t.Errorf("CreateDateTrigger() failed: %v", err)
// 	}

// 	err = store.DeleteDateTrigger(createdDt.Id)
// 	if err != nil {
// 		t.Errorf("DeleteDateTrigger() failed: %v", err)
// 	}

// 	_, err = store.GetDateTrigger(createdDt.Id)
// 	if err == nil {
// 		t.Errorf("GetDateTrigger() should have failed")
// 	}
// }

// func TestDateTriggerDeleteNotFound(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggerdeletenotfound"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	err := store.DeleteDateTrigger(uuid.UUID{})
// 	if err == nil {
// 		t.Errorf("DeleteDateTrigger() should have failed")
// 	}
// }

// func TestDateTriggerList(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggerlist"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// setup
// 	assetTaskId := setupTask(t, store, "1")

// 	// list
// 	dts, err := store.ListDateTriggers()
// 	if err != nil {
// 		t.Errorf("ListDateTriggers() failed: %v", err)
// 	}

// 	triggerMap := make(map[uuid.UUID]tp.DateTrigger)
// 	for _, trigger := range dts {
// 		triggerMap[trigger.Id] = trigger
// 	}

// 	dt := tp.DateTrigger{
// 		Date:   time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Second),
// 		TaskId: assetTaskId,
// 	}

// 	createdDt, err := store.CreateDateTrigger(dt)
// 	if err != nil {
// 		t.Errorf("CreateDateTrigger() failed: %v", err)
// 	}
// 	triggerMap[createdDt.Id] = createdDt

// 	dt.Date = time.Now().AddDate(2, 0, 0).UTC().Truncate(time.Second)
// 	createdDt, err = store.CreateDateTrigger(dt)
// 	if err != nil {
// 		t.Errorf("CreateDateTrigger() failed: %v", err)
// 	}
// 	triggerMap[createdDt.Id] = createdDt

// 	dts, err = store.ListDateTriggers()
// 	if err != nil {
// 		t.Errorf("ListDateTriggers() failed: %v", err)
// 	}

// 	if len(dts) != len(triggerMap) {
// 		t.Errorf("ListDateTriggers() failed: expected %v, got %v", len(triggerMap), len(dts))
// 	}

// 	for _, trigger := range dts {
// 		utest.CompEntities(t, trigger, triggerMap[trigger.Id])
// 	}
// }

// func TestDateTriggerUpdateGet(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggerupdateget"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// setup
// 	assetTaskId := setupTask(t, store, "1")
// 	dt := tp.DateTrigger{
// 		Date:   time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Second),
// 		TaskId: assetTaskId,
// 	}
// 	createdDt, err := store.CreateDateTrigger(dt)
// 	if err != nil {
// 		t.Errorf("CreateDateTrigger() failed: %v", err)
// 	}

// 	// update
// 	createdDt.Date = time.Now().AddDate(2, 0, 0).UTC().Truncate(time.Second)
// 	updatedDt, err := store.UpdateDateTrigger(createdDt.Id, createdDt)
// 	if err != nil {
// 		t.Errorf("UpdateDateTrigger() failed: %v", err)
// 	}

// 	fieldsShouldBeDifferent := utest.ConvertToSet([]string{"Date"})
// 	utest.CompEntitiesFieldsShouldBeDifferent(t, createdDt, updatedDt, fieldsShouldBeDifferent)

// 	// get
// 	retrievedDt, err := store.GetDateTrigger(createdDt.Id)
// 	if err != nil {
// 		t.Errorf("GetDateTrigger() failed: %v", err)
// 	}

// 	utest.CompEntities(t, updatedDt, retrievedDt)
// }

// func TestDateTriggerUpdateNotFound(t *testing.T) {
// 	t.Parallel()
// 	var dbName = "testdatetriggerupdatenotfound"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	dt := tp.DateTrigger{
// 		Id:     uuid.UUID{},
// 		Date:   time.Now().AddDate(1, 0, 0),
// 		TaskId: uuid.UUID{},
// 	}
// 	_, err := store.UpdateDateTrigger(dt.Id, dt)
// 	if err == nil {
// 		t.Errorf("UpdateDateTrigger() should have failed")
// 	}
// }
