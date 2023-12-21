package test

import (
	"testing"
)

// CreateConsumable(string) (int, error)
// DeleteConsumable(int) error
// GetAllConsumable() ([]tp.EquipmentCategory, error)
// GetConsumable(int) (tp.EquipmentCategory, error)
// UpdateConsumable(int, string) error

func TestConsumableCreateGetUpdateGetDeleteGet(t *testing.T) {
	// Create
	coTitle := "test consumable"
	id, err := testStore.CreateConsumable(coTitle)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get
	co, err := testStore.GetConsumable(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if co.Title != coTitle {
		t.Errorf("Get() returned an unexpected value: %s, expected: %s", co.Title, coTitle)
	}

	// Update
	newCoTitle := "new test consumable"
	err = testStore.UpdateConsumable(id, newCoTitle)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	co, err = testStore.GetConsumable(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if co.Title != newCoTitle {
		t.Errorf("Get() returned an unexpected value: %s, expected: %s", co.Title, coTitle)
	}

	// Delete
	err = testStore.DeleteConsumable(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = testStore.GetConsumable(id)
	if err == nil {
		t.Errorf("Get() should have failed")
	}

	err = teardownTable("consumable", id)
	if err != nil {
		t.Errorf("teardownTable() failed: %v", err)
	}
}

func TestConsumableCreateCreateGetAll(t *testing.T) {
	// Create
	coTitle1 := "test consumable 1"
	id1, err := testStore.CreateConsumable(coTitle1)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Create
	coTitle2 := "test consumable 2"
	_, err = testStore.CreateConsumable(coTitle2)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get all
	co, err := testStore.GetAllConsumable()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}

	if len(co) != 2 {
		t.Errorf("GetAll() returned an unexpected value: %d, expected: %d", len(co), 2)
	}

	// data validation
	for _, c := range co {
		if c.Id == id1 {
			if c.Title != coTitle1 {
				t.Errorf("GetAll() returned an unexpected value: %s, expected: %s", c.Title, coTitle1)
			}
		} else {
			if c.Title != coTitle2 {
				t.Errorf("GetAll() returned an unexpected value: %s, expected: %s", c.Title, coTitle2)
			}
		}
	}

	err = teardownTable("consumable", id1)
	if err != nil {
		t.Errorf("teardownTable() failed: %v", err)
	}
}
