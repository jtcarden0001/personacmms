package test

import (
	"testing"
)

func TestAssetCategoryCreateUpdateDelete(t *testing.T) {
	// Create
	id, err := testStore.CreateAssetCategory("test category")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	err = testStore.UpdateAssetCategory(id, "test category 2")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteAssetCategory(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = teardownTable("asset_category", &id)
	if err != nil {
		t.Errorf("teardownTable(asset_category) failed: %v", err)
	}
}
