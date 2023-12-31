package test

import (
	"testing"
)

func TestEquipmentCategoryCreateUpdateDelete(t *testing.T) {
	// Create
	id, err := testStore.CreateEquipmentCategory("test category")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	err = testStore.UpdateEquipmentCategory(id, "test category 2")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteEquipmentCategory(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = teardownTable("equipment_category", &id)
	if err != nil {
		t.Errorf("teardownTable(equipment_category) failed: %v", err)
	}
}
