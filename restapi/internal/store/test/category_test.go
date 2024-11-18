package test

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestCategoryCreate(t *testing.T) {
	store := InitializeStore("testcategorycreate")

	// Create
	cat := types.Category{
		Title:       "testcategory1",
		Description: "test description",
	}

	returnCat, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnCat.Title != cat.Title {
		t.Errorf("Create() failed: expected %s, got %s", cat.Title, returnCat.Title)
	}

	if returnCat.Description != cat.Description {
		t.Errorf("Create() failed: expected %s, got %s", cat.Description, returnCat.Description)
	}
}

func TestCategoryDelete(t *testing.T) {
	store := InitializeStore("testcategorydelete")

	// Delete
	cat := types.Category{
		Title:       "testcategory1",
		Description: "test description",
	}
	_, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeleteCategory(cat.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetCategory(cat.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestCategoryList(t *testing.T) {
	store := InitializeStore("testcategorylist")

	// List
	cats, err := store.ListCategories()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(cats) != 0 {
		t.Errorf("ListCategory() failed: expected 0, got %d", len(cats))
	}

	// Create
	cat := types.Category{
		Title:       "testcategory1",
		Description: "test description",
	}
	_, err = store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	cat.Title = "testcategory2"
	_, err = store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	cats, err = store.ListCategories()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(cats) != 2 {
		t.Errorf("ListCategory() failed: expected 2, got %d", len(cats))
	}
}

func TestCategoryGet(t *testing.T) {
	store := InitializeStore("testcategoryget")

	// Get
	cat := types.Category{
		Title:       "testcategory1",
		Description: "test description",
	}
	_, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	returnCat, err := store.GetCategory(cat.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returnCat.Title != cat.Title {
		t.Errorf("Get() failed: expected %s, got %s", cat.Title, returnCat.Title)
	}

	if returnCat.Description != cat.Description {
		t.Errorf("Get() failed: expected %s, got %s", cat.Description, returnCat.Description)
	}
}

func TestCategoryUpdate(t *testing.T) {
	store := InitializeStore("testcategoryupdate")

	// Update
	cat := types.Category{
		Title:       "testcategory1",
		Description: "test description",
	}
	_, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	cat.Description = "new description"
	returnCat, err := store.UpdateCategory(cat.Title, cat)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	if returnCat.Title != cat.Title {
		t.Errorf("Update() failed: expected %s, got %s", cat.Title, returnCat.Title)
	}

	if returnCat.Description != cat.Description {
		t.Errorf("Update() failed: expected %s, got %s", cat.Description, returnCat.Description)
	}
}
