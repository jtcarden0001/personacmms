package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetTaskConsumableCreateDelete(t *testing.T) {
	dbname := "testassettaskconsumabledelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTaskConsumable{
		AssetTaskId:  setupAssetTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	returnedAssetTaskConsumable, err := store.CreateAssetTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateAssetTaskConsumable() failed %v", err)
	}

	compEntities(t, at, returnedAssetTaskConsumable)

	// Delete
	err = store.DeleteAssetTaskConsumable(at.AssetTaskId, at.ConsumableId)
	if err != nil {
		t.Errorf("DeleteAssetTaskConsumable() failed %v", err)
	}

	_, err = store.GetAssetTaskConsumable(at.AssetTaskId, at.ConsumableId)
	if err == nil {
		t.Errorf("GetAssetTaskConsumable() failed: expected error")
	}
}

func TestAssetTaskConsumableList(t *testing.T) {
	dbname := "testassettaskconsumablelist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	atcs, err := store.ListAssetTaskConsumables()
	if err != nil {
		t.Errorf("ListAssetTaskConsumables() failed %v", err)
	}

	// Create
	at := tp.AssetTaskConsumable{
		AssetTaskId:  setupAssetTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	_, err = store.CreateAssetTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateAssetTaskConsumable() failed %v", err)
	}

	// Create
	at2 := tp.AssetTaskConsumable{
		AssetTaskId:  setupAssetTask(t, store, "2"),
		ConsumableId: setupConsumable(t, store, "2"),
		QuantityNote: "2",
	}

	_, err = store.CreateAssetTaskConsumable(at2)
	if err != nil {
		t.Errorf("CreateAssetTaskConsumable() failed %v", err)
	}

	// List
	ratcs, err := store.ListAssetTaskConsumables()
	if err != nil {
		t.Errorf("ListAssetTaskConsumables() failed %v", err)
	}

	if len(ratcs) != len(atcs)+2 {
		t.Errorf("ListAssetTaskConsumables() failed: expected 1, got %d", len(atcs))
	}
}

func TestAssetTaskConsumableUpdateGet(t *testing.T) {
	dbname := "testassettaskconsumableupdateget"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.AssetTaskConsumable{
		AssetTaskId:  setupAssetTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	_, err := store.CreateAssetTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateAssetTaskConsumable() failed %v", err)
	}

	// Update
	at.QuantityNote = "2"
	_, err = store.UpdateAssetTaskConsumable(at)
	if err != nil {
		t.Errorf("UpdateAssetTaskConsumable() failed %v", err)
	}

	// Get
	rat, err := store.GetAssetTaskConsumable(at.AssetTaskId, at.ConsumableId)
	if err != nil {
		t.Errorf("GetAssetTaskConsumable() failed %v", err)
	}

	compEntities(t, at, rat)
}
