package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCategoryCreate(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupCategory(1, false)
	creC, err := app.CreateCategory(c)
	if err != nil {
		t.Fatalf("CreateCategory() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, c, creC, diffFields)
}

func TestCategoryDelete(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Fatalf("TestCategoryDelete: failed during setup. CreateCategory() failed: %v", err)
	}

	err = app.DeleteCategory(createdCategory.Id.String())
	if err != nil {
		t.Errorf("DeleteCategory() failed: %v", err)
	}

	_, err = app.GetCategory(createdCategory.Id.String())
	if err == nil {
		t.Errorf("GetCategory() returned nil error after deletion")
	}
}

func TestCategoryGet(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Fatalf("TestCategoryGet: failed during setup. CreateCategory() failed: %v", err)
	}

	gotCategory, err := app.GetCategory(createdCategory.Id.String())
	if err != nil {
		t.Errorf("GetCategory() failed: %v", err)
	}

	utest.CompEntities(t, createdCategory, gotCategory)
}

func TestCategoryUpdate(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Fatalf("TestCategoryUpdate: failed during setup. CreateCategory() failed: %v", err)
	}

	c.Description = utest.ToPtr("updated description")
	updatedCategory, err := app.UpdateCategory(createdCategory.Id.String(), c)
	if err != nil {
		t.Errorf("UpdateCategory() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Description"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdCategory, updatedCategory, diffFields)
}
