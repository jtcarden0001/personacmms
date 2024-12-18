package integration

import (
	"errors"
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCategoryCreate(t *testing.T) {
	t.Parallel()
	dbName := "testcategorycreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	cat := tp.Category{
		Title:       "testcategory1",
		Description: utest.ToPtr("test description"),
	}

	returnCat, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, cat, returnCat, fieldsToExclude)
}

func TestCategoryDelete(t *testing.T) {
	t.Parallel()
	dbName := "testcategorydelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Delete something that doesn't exist
	err := store.DeleteCategory("notfound")
	if err == nil {
		t.Errorf("Delete() failed: expected error, got nil")
	}

	cat := tp.Category{
		Title:       "testcategory1",
		Description: utest.ToPtr("test description"),
	}
	_, err = store.CreateCategory(cat)
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
	t.Parallel()
	dbName := "testcategorylist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

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
		Description: utest.ToPtr("test description"),
	}

	catMap["testcategory2"] = tp.Category{
		Title:       "testcategory2",
		Description: utest.ToPtr("test description"),
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
		fieldsToExclude := utest.ConvertToSet([]string{"Id"})
		utest.CompEntitiesExcludeFields(t, cat, newCatMap[key], fieldsToExclude)
	}
}

func TestCategoryUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testcategoryupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Update
	cat := tp.Category{
		Title:       "testcategory1",
		Description: utest.ToPtr("test description"),
	}
	createCat, err := store.CreateCategory(cat)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	cat.Description = utest.ToPtr("new description")
	updateCat, err := store.UpdateCategory(cat.Title, cat)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	differentFields := utest.ConvertToSet([]string{"Description"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createCat, updateCat, differentFields)

	getCat, err := store.GetCategory(cat.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	// exclude no fields
	utest.CompEntities(t, updateCat, getCat)
}

func TestCategoryNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testcategorynotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	_, err := store.GetCategory("notfound")
	var appErr ae.AppError
	if !errors.As(err, &appErr) {
		t.Errorf("Get() failed: expected AppError, got %v", err)
	}

	if appErr.Code != ae.CodeNotFound {
		t.Errorf("Get() failed: expected CodeNotFound, got %v", appErr.Code)
	}
}

func TestCategoryDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testcategorydeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteCategory("notfound")
	if err == nil {
		t.Errorf("DeleteCategory() should have failed")
	}
}

func TestCategoryUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testcategoryupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	cat := tp.Category{
		Title:       "notfound",
		Description: utest.ToPtr("test description"),
	}
	_, err := store.UpdateCategory(cat.Title, cat)
	if err == nil {
		t.Errorf("UpdateCategory() should have failed")
	}
}
