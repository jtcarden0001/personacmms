package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestUsageUnitCreate(t *testing.T) {
	dbName := "testusageunitcreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	usageUnit := tp.UsageUnit{
		Title: "testusageunit1",
	}

	returnedUnit, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, usageUnit, returnedUnit, fieldsToExclude)
}

func TestUsageUnitDelete(t *testing.T) {
	dbName := "testusageunitdelete"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	usageUnit := tp.UsageUnit{
		Title: "testusageunit1",
	}

	returnedUnit, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Delete
	err = store.DeleteUsageUnit(returnedUnit.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Confirm deletion
	_, err = store.GetUsageUnit(returnedUnit.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestUsageUnitList(t *testing.T) {
	dbName := "testusageunitlist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// List
	units, err := store.ListUsageUnits()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	// convert output to a map
	unitMap := make(map[string]tp.UsageUnit)
	for _, unit := range units {
		unitMap[unit.Title] = unit
	}

	// Create
	usageUnit := tp.UsageUnit{
		Title: "testusageunit1",
	}
	u1, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}
	unitMap[u1.Title] = u1

	// Create
	usageUnit.Title = "testusageunit2"
	u2, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}
	unitMap[u2.Title] = u2

	// List
	units, err = store.ListUsageUnits()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(units) != len(unitMap) {
		t.Errorf("List() failed: expected %d, got %d", len(unitMap), len(units))
	}

	// compare
	for _, unit := range units {
		compEntities(t, unit, unitMap[unit.Title])
	}
}

func TestUsageUnitUpdateGet(t *testing.T) {
	dbName := "testusageunitupdateget"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	usageUnit := tp.UsageUnit{
		Title: "testusageunit1",
	}

	returnedUnit, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	usageUnit.Title = "testusageunit2"
	updatedUnit, err := store.UpdateUsageUnit(returnedUnit.Title, usageUnit)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	if updatedUnit.Title != usageUnit.Title {
		t.Errorf("Update() failed: expected %s, got %s", usageUnit.Title, updatedUnit.Title)
	}

	// Get
	unit, err := store.GetUsageUnit(updatedUnit.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if unit.Title != updatedUnit.Title {
		t.Errorf("Get() failed: expected %s, got %s", updatedUnit.Title, unit.Title)
	}
}
