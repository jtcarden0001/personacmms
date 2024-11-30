package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTaskConsumableCreateDelete(t *testing.T) {
	dbname := "testtaskconsumabledelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.TaskConsumable{
		TaskId:       setupTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	returnedTaskConsumable, err := store.CreateTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateTaskConsumable() failed %v", err)
	}

	compEntities(t, at, returnedTaskConsumable)

	// Delete
	err = store.DeleteTaskConsumable(at.TaskId, at.ConsumableId)
	if err != nil {
		t.Errorf("DeleteTaskConsumable() failed %v", err)
	}

	_, err = store.GetTaskConsumable(at.TaskId, at.ConsumableId)
	if err == nil {
		t.Errorf("GetTaskConsumable() failed: expected error")
	}
}

func TestTaskConsumableList(t *testing.T) {
	dbname := "testtaskconsumablelist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	atcs, err := store.ListTaskConsumables()
	if err != nil {
		t.Errorf("ListTaskConsumables() failed %v", err)
	}

	// Create
	at := tp.TaskConsumable{
		TaskId:       setupTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	_, err = store.CreateTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateTaskConsumable() failed %v", err)
	}

	// Create
	at2 := tp.TaskConsumable{
		TaskId:       setupTask(t, store, "2"),
		ConsumableId: setupConsumable(t, store, "2"),
		QuantityNote: "2",
	}

	_, err = store.CreateTaskConsumable(at2)
	if err != nil {
		t.Errorf("CreateTaskConsumable() failed %v", err)
	}

	// List
	ratcs, err := store.ListTaskConsumables()
	if err != nil {
		t.Errorf("ListTaskConsumables() failed %v", err)
	}

	if len(ratcs) != len(atcs)+2 {
		t.Errorf("ListTaskConsumables() failed: expected 1, got %d", len(atcs))
	}
}

func TestTaskConsumableUpdateGet(t *testing.T) {
	dbname := "testtaskconsumableupdateget"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	at := tp.TaskConsumable{
		TaskId:       setupTask(t, store, "1"),
		ConsumableId: setupConsumable(t, store, "1"),
		QuantityNote: "1",
	}
	_, err := store.CreateTaskConsumable(at)
	if err != nil {
		t.Errorf("CreateTaskConsumable() failed %v", err)
	}

	// Update
	at.QuantityNote = "2"
	_, err = store.UpdateTaskConsumable(at)
	if err != nil {
		t.Errorf("UpdateTaskConsumable() failed %v", err)
	}

	// Get
	rat, err := store.GetTaskConsumable(at.TaskId, at.ConsumableId)
	if err != nil {
		t.Errorf("GetTaskConsumable() failed %v", err)
	}

	compEntities(t, at, rat)
}
