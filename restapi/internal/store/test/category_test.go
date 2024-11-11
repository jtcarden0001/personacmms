package test

import (
	"testing"
)

func TestCategoryCRUD(t *testing.T) {
	// Create
	id, err := testStore.CreateCategory("test category")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// create a second category
	id2, err := testStore.CreateCategory("test category 2")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	cats, err := testStore.ListCategory()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}
	if len(cats) != 2 {
		t.Errorf("ListCategory() failed: expected 2, got %d", len(cats))
	}

	// Get
	cat, err := testStore.GetCategory(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if (cat.Id != id) || (cat.Title != "test category") {
		t.Errorf("GetCategory() failed: expected %d, got %d", id, cat.Id)
	}

	// Update
	err = testStore.UpdateCategory(id, "test category 3")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	cat, err = testStore.GetCategory(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if (cat.Id != id) || (cat.Title != "test category 3") {
		t.Errorf("GetCategory() failed: expected %d, got %d", id, cat.Id)
	}

	// Delete
	err = testStore.DeleteCategory(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = testStore.DeleteCategory(id2)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// err = teardownTable("asset_category")
	// if err != nil {
	// 	t.Errorf("teardownTable(asset_category) failed: %v", err)
	// }
}
