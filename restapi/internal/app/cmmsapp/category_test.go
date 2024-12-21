package cmmsapp

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category := tp.Category{Id: uuid.New(), Title: "category1"}
	createdCategory, err := app.CreateCategory(category)
	assert.NoError(t, err)
	assert.Equal(t, category.Title, createdCategory.Title)
}

func TestDeleteCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category := tp.Category{Title: "category1"}
	db.CreateCategory(category)

	err := app.DeleteCategory("category1")
	assert.NoError(t, err)

	_, err = app.GetCategory("category1")
	assert.Error(t, err)
}

func TestListCategories(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category1 := tp.Category{Id: uuid.New(), Title: "category1"}
	category2 := tp.Category{Id: uuid.New(), Title: "category2"}
	db.CreateCategory(category1)
	db.CreateCategory(category2)

	categories, err := app.ListCategories()
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestGetCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category := tp.Category{Id: uuid.New(), Title: "category1"}
	db.CreateCategory(category)

	retrievedCategory, err := app.GetCategory("category1")
	assert.NoError(t, err)
	assert.Equal(t, category.Title, retrievedCategory.Title)
}

func TestUpdateCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category := tp.Category{Id: uuid.New(), Title: "category1"}
	db.CreateCategory(category)

	updatedCategory := tp.Category{Title: "category1_updated"}
	_, err := app.UpdateCategory("category1", updatedCategory)
	assert.NoError(t, err)

	retrievedCategory, err := app.GetCategory("category1_updated")
	assert.NoError(t, err)
	assert.Equal(t, updatedCategory.Title, retrievedCategory.Title)
}

func TestCreateCategoryWithEmptyTitle(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	category := tp.Category{Id: uuid.New(), Title: ""}
	_, err := app.CreateCategory(category)
	assert.Error(t, err)
}

func TestDeleteNonExistentCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	err := app.DeleteCategory("nonexistentcategory")
	assert.Error(t, err)
}

func TestGetNonExistentCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	_, err := app.GetCategory("nonexistentcategory")
	assert.Error(t, err)
}

func TestUpdateNonExistentCategory(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	updatedCategory := tp.Category{Title: "category1_updated"}
	_, err := app.UpdateCategory("nonexistentcategory", updatedCategory)
	assert.Error(t, err)
}
