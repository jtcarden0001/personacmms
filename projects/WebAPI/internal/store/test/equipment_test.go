package test

import (
	"fmt"
	"testing"

	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetNotExists(t *testing.T) {
	_, err := testStore.GetEquipment(0)
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

func TestCreateGetDeleteGet(t *testing.T) {
	// Create
	id, err := testStore.CreateEquipment("test equipment", "test description")
	fmt.Println("test equipment id:", id)
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

	err = testStore.ResetSequenceEquipment(id)
	if err != nil {
		t.Errorf("ResetSequence() failed: %v", err)
	}
}

func TestGetAllCreateCreateGetAll(t *testing.T) {
	// Get all
	e1, err := testStore.GetAllEquipment()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}
	origLength := len(e1)

	// Create
	id1, err := testStore.CreateEquipment("test equipment 1", "test description 1")
	fmt.Println("test equipment id:", id1)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Create
	var id2 int
	id2, err = testStore.CreateEquipment("test equipment 2", "test description 2")
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

	err = testStore.ResetSequenceEquipment(id1)
	if err != nil {
		t.Errorf("ResetSequence() failed: %v", err)
	}
}
