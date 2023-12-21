package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetNotExists(t *testing.T) {
	err := initStore()
	if err != nil {
		t.Errorf("initStore() failed: %v", err)
		return
	}

	_, err = testStore.GetEquipment(0)
	errString := "sql: no rows in result set"
	if err.Error() != errString {
		t.Errorf("Get() should have failed with error: %s", errString)
	}
}

func TestDeleteNotExists(t *testing.T) {
	err := testStore.DeleteEquipment(0)
	if err != nil {
		t.Errorf("Delete() record that does not exist should produce no error")
	}
}

// TODO: parse out the initialization code so that we only wipe the db and initialize foreign keys once per test run.
func TestCreateGetDeleteGet(t *testing.T) {
	//// initialization
	err := initStore()
	if err != nil {
		t.Errorf("initStore() failed: %v", err)
		return
	}

	ecId, err := initEquipmentFKs()
	if err != nil {
		t.Errorf("initEquipmentFKs() failed: %v", err)
		return
	}

	//// Test
	// Create
	id, err := testStore.CreateEquipment("test equipment", 2023, "test make", "test model number", "test description", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get
	_, err = testStore.GetEquipment(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteEquipment(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = testStore.GetEquipment(id)
	if err == nil {
		t.Errorf("Get() should have failed")
	}
}

func TestGetAllCreateCreateGetAll(t *testing.T) {
	//// initialization
	err := initStore()
	if err != nil {
		t.Errorf("initStore() failed: %v", err)
		return
	}

	ecId, err := initEquipmentFKs()
	if err != nil {
		t.Errorf("initEquipmentFKs() failed: %v", err)
		return
	}

	//// Test
	// Get all
	e1, err := testStore.GetAllEquipment()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}
	origLength := len(e1)

	// Create
	id1, err := testStore.CreateEquipment("test equipment 1", 2023, "test make 1", "test model number 1", "test description 1", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Create
	var id2 int
	id2, err = testStore.CreateEquipment("test equipment 2", 2023, "test make 2", "test model number 2", "test description 2", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get all
	var e2 []tp.Equipment
	e2, err = testStore.GetAllEquipment()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}
	assert.Equal(t, origLength+2, len(e2))

	// Delete
	err = testStore.DeleteEquipment(id1)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteEquipment(id2)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}
}

func initEquipmentFKs() (int, error) {
	ecId, err := testStore.CreateEquipmentCategory("test equipment category")
	if err != nil {
		return 0, err
	}

	return ecId, nil
}
