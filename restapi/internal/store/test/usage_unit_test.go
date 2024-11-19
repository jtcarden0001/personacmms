package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestUsageUnitCreate(t *testing.T) {
	store := InitializeStore("testusageunitcreate")

	// Create
	usageUnit := types.UsageUnit{
		Title: "testusageunit1",
	}

	returnedUnit, err := store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnedUnit.Title != usageUnit.Title {
		t.Errorf("Create() failed: expected %s, got %s", usageUnit.Title, returnedUnit.Title)
	}

	if returnedUnit.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-empty ID, got empty")
	}
}

func TestUsageUnitDelete(t *testing.T) {
	store := InitializeStore("testusageunitdelete")

	// Create
	usageUnit := types.UsageUnit{
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
	store := InitializeStore("testusageunitlist")

	// List
	units, err := store.ListUsageUnits()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(units) != 0 {
		t.Errorf("List() failed: expected 0, got %d", len(units))
	}

	// Create
	usageUnit := types.UsageUnit{
		Title: "testusageunit1",
	}

	_, err = store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	usageUnit.Title = "testusageunit2"
	_, err = store.CreateUsageUnit(usageUnit)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	units, err = store.ListUsageUnits()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	expected := 2
	if len(units) != expected {
		t.Errorf("List() failed: expected %d units, got %d", expected, len(units))
	}
}

func TestUsageUnitUpdateGet(t *testing.T) {
	store := InitializeStore("testusageunitupdateget")

	// Create
	usageUnit := types.UsageUnit{
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
