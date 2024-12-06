package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestConsumableCreate(t *testing.T) {
	t.Parallel()
	dbName := "testconsumablecreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

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
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

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
	_, err = store.GetConsumable(returnedConsumable.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestConsumableList(t *testing.T) {
	t.Parallel()
	dbName := "testconsumablelist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

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
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

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
	consumable, err = store.GetConsumable(consumable.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if consumable.Title != "testconsumable2" {
		t.Errorf("Get() failed: expected testconsumable2, got %s", consumable.Title)
	}
}
