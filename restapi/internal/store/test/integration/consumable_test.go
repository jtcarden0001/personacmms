package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestConsumableCreate(t *testing.T) {
	t.Parallel()
	dbName := "testconsumablecreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	consumable := types.Consumable{
		Title: "testconsumable1",
	}

	returnedConsumable, err := store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnedConsumable.Title != consumable.Title {
		t.Errorf("Create() failed: expected %s, got %s", consumable.Title, returnedConsumable.Title)
	}

	if returnedConsumable.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-empty ID, got empty")
	}
}

func TestConsumableDelete(t *testing.T) {
	t.Parallel()
	dbName := "testconsumabledelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	consumable := types.Consumable{
		Title: "testconsumable1",
	}

	returnedConsumable, err := store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Delete
	err = store.DeleteConsumable(returnedConsumable.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Confirm deletion
	_, err = store.GetConsumableByTitle(returnedConsumable.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestConsumableDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testconsumabledeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteConsumable("notfound")
	if err == nil {
		t.Errorf("DeleteConsumable() should have failed")
	}
}

func TestConsumableList(t *testing.T) {
	t.Parallel()
	dbName := "testconsumablelist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	consumables, err := store.ListConsumables()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(consumables) != 0 {
		t.Errorf("List() failed: expected 0, got %d", len(consumables))
	}

	// Create
	consumable := types.Consumable{
		Title: "testconsumable1",
	}

	_, err = store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Create
	consumable.Title = "testconsumable2"
	_, err = store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	consumables, err = store.ListConsumables()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	expectedConsumableCount := 2
	if len(consumables) != expectedConsumableCount {
		t.Errorf("List() failed: expected %d, got %d", expectedConsumableCount, len(consumables))
	}
}

func TestConsumableUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testconsumableupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	consumable := types.Consumable{
		Title: "testconsumable1",
	}

	returnedConsumable, err := store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	consumable.Title = "testconsumable2"
	_, err = store.UpdateConsumable(returnedConsumable.Title, consumable)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	consumable, err = store.GetConsumableByTitle(consumable.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if consumable.Title != "testconsumable2" {
		t.Errorf("Get() failed: expected testconsumable2, got %s", consumable.Title)
	}
}

func TestConsumableUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testconsumableupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	consumable := types.Consumable{
		Title: "notfound",
	}
	_, err := store.UpdateConsumable("notfound", consumable)
	if err == nil {
		t.Errorf("UpdateConsumable() should have failed")
	}
}
