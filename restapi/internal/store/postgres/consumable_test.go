package postgres_test

import (
	"testing"

	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

// TODO: TestConsumableAssociateWithTask

// TODO: TestConsumableAssociateWithWorkOrder

func TestConsumableCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testconsumablecreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := utest.SetupConsumable(1, true)

	// test
	createdConsumable, err := store.CreateConsumable(c)
	if err != nil {
		t.Errorf("CreateConsumable() failed: %v", err)
	}

	utest.CompEntities(t, c, createdConsumable)
}

func TestConsumableDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testconsumabledelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := utest.SetupConsumable(1, true)
	createdConsumable, err := store.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestConsumableDelete: failed during setup. CreateConsumable() failed: %v", err)
	}

	// test
	err = store.DeleteConsumable(createdConsumable.Id)
	if err != nil {
		t.Errorf("TestConsumableDelete: DeleteConsumable() failed: %v", err)
	}

	_, err = store.GetConsumable(createdConsumable.Id)
	if err == nil {
		t.Errorf("TestConsumableDelete: GetConsumable() returned nil error after deletion")
	}
}

// TODO: TestConsumableDisassociateWithTask

// TODO: TestConsumableDisassociateWithWorkOrder

func TestConsumableGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testconsumableget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := utest.SetupConsumable(1, true)
	createConsumable, err := store.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestConsumableGet: failed during setup. CreateConsumable() failed: %v", err)
	}

	// test
	getConsumable, err := store.GetConsumable(createConsumable.Id)
	if err != nil {
		t.Errorf("GetConsumable() failed: %v", err)
	}

	utest.CompEntities(t, createConsumable, getConsumable)
}

func TestConsumableList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testconsumablelist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c1 := utest.SetupConsumable(1, true)
	c2 := utest.SetupConsumable(2, true)
	c3 := utest.SetupConsumable(3, true)

	_, err := store.CreateConsumable(c1)
	if err != nil {
		t.Errorf("TestConsumableList: failed during setup. CreateConsumable() failed: %v", err)
	}
	_, err = store.CreateConsumable(c2)
	if err != nil {
		t.Errorf("TestConsumableList: failed during setup. CreateConsumable() failed: %v", err)
	}
	_, err = store.CreateConsumable(c3)
	if err != nil {
		t.Errorf("TestConsumableList: failed during setup. CreateConsumable() failed: %v", err)
	}

	// test
	consumables, err := store.ListConsumables()
	if err != nil {
		t.Errorf("ListConsumables() failed: %v", err)
	}

	if len(consumables) != 3 {
		t.Errorf("ListConsumables() returned %d consumables, expected 3", len(consumables))
	}
}

// TODO: TestConsumableListByTask

// TODO: TestConsumableListByWorkOrder

func TestConsumableUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testconsumableupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := utest.SetupConsumable(1, true)
	createdConsumable, err := store.CreateConsumable(c)
	if err != nil {
		t.Errorf("TestConsumableUpdate: failed during setup. CreateConsumable() failed: %v", err)
	}

	// test
	c.Title = "Updated Title"
	updatedConsumable, err := store.UpdateConsumable(c)
	if err != nil {
		t.Errorf("UpdateConsumable() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdConsumable, updatedConsumable, differentFields)
}
