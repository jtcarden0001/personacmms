package integration

// import (
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
// )

// var preseedTimeUnitCount = 4

// func TestTimeUnitCreate(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitcreate"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// Create
// 	timeUnit := types.TimeUnit{
// 		Title: "testtimeunit1",
// 	}

// 	returnedUnit, err := store.CreateTimeUnit(timeUnit)
// 	if err != nil {
// 		t.Errorf("Create() failed: %v", err)
// 	}

// 	if returnedUnit.Title != timeUnit.Title {
// 		t.Errorf("Create() failed: expected %s, got %s", timeUnit.Title, returnedUnit.Title)
// 	}

// 	if returnedUnit.Id == uuid.Nil {
// 		t.Errorf("Create() failed: expected non-empty ID, got empty")
// 	}
// }

// func TestTimeUnitDelete(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitdelete"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// Create
// 	timeUnit := types.TimeUnit{
// 		Title: "testtimeunit1",
// 	}

// 	returnedUnit, err := store.CreateTimeUnit(timeUnit)
// 	if err != nil {
// 		t.Errorf("Create() failed: %v", err)
// 	}

// 	// Delete
// 	err = store.DeleteTimeUnit(returnedUnit.Title)
// 	if err != nil {
// 		t.Errorf("Delete() failed: %v", err)
// 	}

// 	// Confirm deletion
// 	_, err = store.GetTimeUnit(returnedUnit.Title)
// 	if err == nil {
// 		t.Errorf("Get() failed: expected error, got nil")
// 	}
// }

// func TestTimeUnitDeleteNotFound(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitdeletenotfound"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	err := store.DeleteTimeUnit("nonexistent-title")
// 	if err == nil {
// 		t.Errorf("DeleteTimeUnit() failed: expected error, got nil")
// 	}
// }

// func TestTimeUnitList(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitlist"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// List
// 	timeUnits, err := store.ListTimeUnits()
// 	if err != nil {
// 		t.Errorf("List() failed: %v", err)
// 	}

// 	if len(timeUnits) != preseedTimeUnitCount {
// 		t.Errorf("List() failed: expected %d, got %d", preseedTimeUnitCount, len(timeUnits))
// 	}

// 	// Create
// 	timeUnit := types.TimeUnit{
// 		Title: "testtimeunit1",
// 	}

// 	_, err = store.CreateTimeUnit(timeUnit)
// 	if err != nil {
// 		t.Errorf("Create() failed: %v", err)
// 	}

// 	// List
// 	timeUnits, err = store.ListTimeUnits()
// 	if err != nil {
// 		t.Errorf("List() failed: %v", err)
// 	}

// 	if len(timeUnits) != 1+preseedTimeUnitCount {
// 		t.Errorf("List() failed: expected 1, got %d", len(timeUnits))
// 	}
// }

// func TestTimeUnitUpdateGet(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitupdateget"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	// Create
// 	timeUnit := types.TimeUnit{
// 		Title: "testtimeunit1",
// 	}

// 	returnedUnit, err := store.CreateTimeUnit(timeUnit)
// 	if err != nil {
// 		t.Errorf("Create() failed: %v", err)
// 	}

// 	// Update
// 	returnedUnit.Title = "testtimeunit2"
// 	_, err = store.UpdateTimeUnit(timeUnit.Title, returnedUnit)
// 	if err != nil {
// 		t.Errorf("Update() failed: %v", err)
// 	}

// 	// Get
// 	updatedUnit, err := store.GetTimeUnit(returnedUnit.Title)
// 	if err != nil {
// 		t.Errorf("Get() failed: %v", err)
// 	}

// 	if updatedUnit.Title != returnedUnit.Title {
// 		t.Errorf("Get() failed: expected %s, got %s", returnedUnit.Title, updatedUnit.Title)
// 	}
// }

// func TestTimeUnitUpdateNotFound(t *testing.T) {
// 	t.Parallel()
// 	dbName := "testtimeunitupdatenotfound"
// 	store := utest.InitializeStore(dbName)
// 	defer utest.CloseStore(store, dbName)

// 	timeUnit := types.TimeUnit{
// 		Title: "nonexistent-title",
// 	}
// 	_, err := store.UpdateTimeUnit("nonexistent-title", timeUnit)
// 	if err == nil {
// 		t.Errorf("UpdateTimeUnit() failed: expected error, got nil")
// 	}
// }
