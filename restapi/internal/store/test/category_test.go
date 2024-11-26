package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestCategoryCreate(t *testing.T) {
	store := initializeStore("testcategorycreate")

	// Create
	cat := tp.Category{
		Title:       "testcategory1",
		Description: "test description",
	}

	returnCat, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, cat, returnCat, fieldsToExclude)
}

func TestCategoryDelete(t *testing.T) {
	store := initializeStore("testcategorydelete")

	// Delete
	cat := tp.Category{
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
	store := initializeStore("testcategorylist")

	// List
	cats, err := store.ListCategories()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	// create a map of the categories title: tp.category
	catMap := make(map[string]tp.Category)
	for _, cat := range cats {
		catMap[cat.Title] = cat
	}

	catMap["testcategory1"] = tp.Category{
		Title:       "testcategory1",
		Description: "test description",
	}

	catMap["testcategory2"] = tp.Category{
		Title:       "testcategory2",
		Description: "test description",
	}

	// create the 2 new categories
	_, err = store.CreateCategory(catMap["testcategory1"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	_, err = store.CreateCategory(catMap["testcategory2"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	newCats, err := store.ListCategories()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(newCats) != len(cats)+2 {
		t.Errorf("List() failed: expected %d, got %d", len(cats)+2, len(newCats))
	}

	newCatMap := make(map[string]tp.Category)
	for _, cat := range newCats {
		newCatMap[cat.Title] = cat
	}

	// compare the two maps
	for key, cat := range catMap {
		fieldsToExclude := convertToSet([]string{"Id"})
		compEntitiesExcludeFields(t, cat, newCatMap[key], fieldsToExclude)
	}
}

func TestCategoryUpdateGet(t *testing.T) {
	store := initializeStore("testcategoryupdate")

	// Update
	cat := tp.Category{
		Title:       "testcategory1",
		Description: "test description",
	}
	createCat, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	cat.Description = "new description"
	updateCat, err := store.UpdateCategory(cat.Title, cat)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	differentFields := convertToSet([]string{"Description"})
	compEntitiesFieldsShouldBeDifferent(t, createCat, updateCat, differentFields)

	getCat, err := store.GetCategory(cat.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	// exclude no fields
	compEntities(t, updateCat, getCat)
}
