package test

import (
	"testing"
)

func TestCategoryCRUD(t *testing.T) {
	// Create
	cat1, err := testStore.CreateCategory("test category", "test description")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// create a second category
	cat2, err := testStore.CreateCategory("test category 2", "test description 2")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	cats, err := testStore.ListCategories()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}
	if len(cats) != 2 {
		t.Errorf("ListCategory() failed: expected 2, got %d", len(cats))
	}

	// Get
	getCat, err := testStore.GetCategory(cat1.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if (getCat.Id != cat1.Id) || (getCat.Title != cat1.Title) {
		t.Errorf("GetCategory() failed: expected %d, got %d", cat1.Id, getCat.Id)
	}

	// Update
	upCat, err := testStore.UpdateCategory(cat1.Title, "test category 3", "test description 3")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	getCat, err = testStore.GetCategory(upCat.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}
	if (upCat.Id != getCat.Id) || (upCat.Title != getCat.Title) {
		t.Errorf("GetCategory() failed: expected %d, got %d", getCat.Id, upCat.Id)
	}

	// Delete
	err = testStore.DeleteCategory(upCat.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = testStore.DeleteCategory(cat2.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// err = teardownTable("asset_category")
	// if err != nil {
	// 	t.Errorf("teardownTable(asset_category) failed: %v", err)
	// }
}
