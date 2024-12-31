package cmmsapp

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingCategory := utest.SetupCategory(1, false)
	_, err = app.CreateCategory(conflictingCategory)
	if err != nil {
		t.Errorf("TestCreateCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	emptyTitleCategory := utest.SetupCategory(2, false)
	emptyTitleCategory.Title = ""

	testCases := []struct {
		name          string
		category      tp.Category
		shouldSucceed bool
	}{
		{"valid category", utest.SetupCategory(3, false), true},
		{"non nil id", utest.SetupCategory(4, true), false},
		{"empty title", emptyTitleCategory, false},
		{"conflicting title", conflictingCategory, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateCategory(tc.category)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestDeleteCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		categoryId    string
		shouldSucceed bool
	}{
		{"valid category deletion", createdCategory.Id.String(), true},
		{"invalid category ID", "invalid", false},
		{"nil category ID", uuid.Nil.String(), false},
		{"empty category ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteCategory(tc.categoryId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestGetCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		categoryId    string
		shouldSucceed bool
	}{
		{"valid category", createdCategory.Id.String(), true},
		{"invalid category ID", "invalid", false},
		{"nil category ID", uuid.Nil.String(), false},
		{"empty category ID", "", false},
		{"non-existent category", uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetCategory(tc.categoryId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListCategories(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListCategories")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupCategory(1, false)
	_, err = app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListCategories: failed during setup. CreateCategory() failed: %v", err)
	}

	c = utest.SetupCategory(2, false)
	_, err = app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListCategories: failed during setup. CreateCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		count         int
		shouldSucceed bool
	}{
		{"valid list", 2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs, err := app.ListCategories()
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListCategories() failed: %v", err)
				} else {
					if len(cs) != tc.count {
						t.Errorf("ListCategories() failed: expected %d categories, got %d", tc.count, len(cs))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListCategories() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	categoryCount := 5
	var ids []string
	categories := make(map[string]tp.Category)
	nilIdCategories := make(map[string]tp.Category)
	for i := 0; i < categoryCount; i++ {
		c := utest.SetupCategory(i, false)
		cc, err := app.CreateCategory(c)
		if err != nil {
			t.Errorf("TestUpdateCategory: failed during setup. CreateCategory() failed: %v", err)
		}

		ids = append(ids, cc.Id.String())
		categories[cc.Id.String()] = cc
		nilIdCategories[cc.Id.String()] = c
	}

	testCases := []struct {
		name          string
		categoryId    string
		category      tp.Category
		title         string
		shouldSucceed bool
	}{
		{"valid category with matching IDs", ids[0], categories[ids[0]], "valid title1", true},
		{"valid category with Category.Id nil", ids[1], nilIdCategories[ids[1]], "valid title2", true},
		{"mismatching category ID and Category.Id", ids[2], categories[ids[3]], "valid title3", false},
		{"non-existent category", uuid.New().String(), tp.Category{}, "valid title3", false},

		{"invalid category ID", "invalid", tp.Category{}, "valid title3", false},
		{"nil category ID", uuid.Nil.String(), tp.Category{}, "valid title3", false},
		{"empty category ID", "", tp.Category{}, "valid title3", false},
		{"conflicting id", ids[4], categories[ids[3]], "valid title3", false},

		{"empty title", ids[1], categories[ids[1]], "", false},
		{"minimum length title", ids[1], categories[ids[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", ids[1], categories[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", ids[1], categories[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
		{"conflicting title", ids[2], categories[ids[2]], categories[ids[3]].Title, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.category.Title = tc.title
			_, err := app.UpdateCategory(tc.categoryId, tc.category)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name          string
		category      tp.Category
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid category", utest.SetupCategory(1, false), uuid.New(), "valid title", true},
		{"nil id", utest.SetupCategory(2, false), uuid.Nil, "valid title", false},

		{"empty title", utest.SetupCategory(3, false), uuid.New(), "", false},
		{"minimum length title", utest.SetupCategory(4, false), uuid.New(), strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", utest.SetupCategory(5, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", utest.SetupCategory(6, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.category.Id = tc.id
			tc.category.Title = tc.title
			err := app.validateCategory(tc.category)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestCategoryExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCategoryExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestCategoryExists: failed during setup. CreateCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		categoryId    string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid category", createdCategory.Id.String(), true, true},
		{"non-existent category", uuid.New().String(), false, true},

		{"invalid category ID", "invalid", false, false},
		{"nil category ID", uuid.Nil.String(), false, false},
		{"empty category ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.categoryExists(tc.categoryId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("categoryExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("categoryExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("categoryExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
