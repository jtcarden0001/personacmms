package postgres_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupCategory(identifier int) tp.Category {
	return tp.Category{
		Id:          uuid.New(),
		Title:       fmt.Sprintf("Category %d", identifier),
		Description: utest.ToPtr(fmt.Sprintf("Category %d description", identifier)),
	}
}

func TestCategoryCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testcategorycreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := setupCategory(1)

	// test
	createdCategory, err := store.CreateCategory(c)
	if err != nil {
		t.Errorf("CreateCategory() failed: %v", err)
	}

	utest.CompEntities(t, c, createdCategory)
}

func TestCategoryDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testcategorydelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := setupCategory(1)
	createdCategory, err := store.CreateCategory(c)
	if err != nil {
		t.Errorf("TestCategoryDelete: failed during setup. CreateCategory() failed: %v", err)
	}

	// test
	err = store.DeleteCategory(createdCategory.Id)
	if err != nil {
		t.Errorf("TestCategoryDelete: DeleteCategory() failed: %v", err)
	}

	_, err = store.GetCategory(createdCategory.Id)
	if err == nil {
		t.Errorf("TestCategoryDelete: GetCategory() returned nil error after deletion")
	}
}

func TestCategoryGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testcategoryget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := setupCategory(1)
	createCategory, err := store.CreateCategory(c)
	if err != nil {
		t.Errorf("TestCategoryGet: failed during setup. CreateCategory() failed: %v", err)
	}

	// test
	getCategory, err := store.GetCategory(createCategory.Id)
	if err != nil {
		t.Errorf("GetCategory() failed: %v", err)
	}

	utest.CompEntities(t, createCategory, getCategory)
}

func TestCategoryList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testcategorylist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c1 := setupCategory(1)
	c2 := setupCategory(2)
	c3 := setupCategory(3)

	_, err := store.CreateCategory(c1)
	if err != nil {
		t.Errorf("TestCategoryList: failed during setup. CreateCategory() failed: %v", err)
	}
	_, err = store.CreateCategory(c2)
	if err != nil {
		t.Errorf("TestCategoryList: failed during setup. CreateCategory() failed: %v", err)
	}
	_, err = store.CreateCategory(c3)
	if err != nil {
		t.Errorf("TestCategoryList: failed during setup. CreateCategory() failed: %v", err)
	}

	// test
	categories, err := store.ListCategories()
	if err != nil {
		t.Errorf("ListCategories() failed: %v", err)
	}

	if len(categories) != 3 {
		t.Errorf("ListCategories() returned %d categories, expected 3", len(categories))
	}
}

func TestCategoryUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testcategoryupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	c := setupCategory(1)
	createdCategory, err := store.CreateCategory(c)
	if err != nil {
		t.Errorf("TestCategoryUpdate: failed during setup. CreateCategory() failed: %v", err)
	}

	// test
	c.Title = "Updated Name"
	c.Description = utest.ToPtr("Updated Description")
	updatedCategory, err := store.UpdateCategory(c)
	if err != nil {
		t.Errorf("UpdateCategory() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title", "Description"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdCategory, updatedCategory, differentFields)
}
