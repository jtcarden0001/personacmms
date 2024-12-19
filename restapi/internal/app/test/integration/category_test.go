package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	testapp := newApp("testcreatecategory")
	category := tp.Category{Title: "New Category " + utest.UniqueSuffix()}

	createdCategory, err := testapp.CreateCategory(category)
	assert.NoError(t, err)
	assert.Equal(t, category.Title, createdCategory.Title)
}

func TestDeleteCategory(t *testing.T) {
	testapp := newApp("testdeletecategory")
	category := tp.Category{Title: "Category to Delete " + utest.UniqueSuffix()}
	_, _ = testapp.CreateCategory(category)

	err := testapp.DeleteCategory(category.Title)
	assert.NoError(t, err)

	_, err = testapp.GetCategory(category.Title)
	assert.Error(t, err)
}

func TestListCategories(t *testing.T) {
	testapp := newApp("testlistcategories")
	category1 := tp.Category{Title: "Category 1 " + utest.UniqueSuffix()}
	category2 := tp.Category{Title: "Category 2 " + utest.UniqueSuffix()}
	_, _ = testapp.CreateCategory(category1)
	_, _ = testapp.CreateCategory(category2)

	categories, err := testapp.ListCategories()
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestGetCategory(t *testing.T) {
	testapp := newApp("testgetcategory")
	category := tp.Category{Title: "Existing Category " + utest.UniqueSuffix()}
	_, _ = testapp.CreateCategory(category)

	retrievedCategory, err := testapp.GetCategory(category.Title)
	assert.NoError(t, err)
	assert.Equal(t, category.Title, retrievedCategory.Title)
}

func TestUpdateCategory(t *testing.T) {
	testapp := newApp("testupdatecategory")
	oldCategory := tp.Category{Title: "Old Category " + utest.UniqueSuffix()}
	_, _ = testapp.CreateCategory(oldCategory)

	newCategory := tp.Category{Title: "Updated Category " + utest.UniqueSuffix()}
	updatedCategory, err := testapp.UpdateCategory(oldCategory.Title, newCategory)
	assert.NoError(t, err)
	assert.Equal(t, newCategory.Title, updatedCategory.Title)
}
