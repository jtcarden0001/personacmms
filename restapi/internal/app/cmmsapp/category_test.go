package cmmsapp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingCategory := setupApiCategoryRequest(1)
	_, err = app.CreateCategory(conflictingCategory)
	if err != nil {
		t.Errorf("TestCreateCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	emptyTitleCategory := setupApiCategoryRequest(2)
	emptyTitleCategory.Title = ""

	testCases := []struct {
		name          string
		category      apitp.CategoryRequest
		shouldSucceed bool
	}{
		{"valid category", setupApiCategoryRequest(3), true},
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

	c := setupApiCategoryRequest(1)
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

	c := setupApiCategoryRequest(1)
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

	c := setupApiCategoryRequest(1)
	_, err = app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListCategories: failed during setup. CreateCategory() failed: %v", err)
	}

	c = setupApiCategoryRequest(2)
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

func TestListCategoriesByAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListCategoriesByAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	c := setupApiCategoryRequest(1)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListCategoriesByAsset: failed during setup. CreateCategory() failed: %v", err)
	}

	a := setupApiAssetRequest(1)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListCategoriesByAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	_, err = app.AssociateAssetWithCategory(createdAsset.Id.String(), createdCategory.Id.String())
	if err != nil {
		t.Errorf("TestListCategoriesByAsset: failed during setup. AddCategoryToAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetId       string
		count         int
		shouldSucceed bool
	}{
		{"valid category", createdAsset.Id.String(), 1, true},
		{"invalid asset ID", "invalid", 0, false},
		{"nil asset ID", uuid.Nil.String(), 0, false},
		{"empty asset ID", "", 0, false},
		{"non-existent asset", uuid.New().String(), 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs, err := app.ListCategoriesByAsset(tc.assetId)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListCategoriesByAsset() failed: %v", err)
				} else {
					if len(cs) != tc.count {
						t.Errorf("ListCategoriesByAsset() failed: expected %d categories, got %d", tc.count, len(cs))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListCategoriesByAsset() should have failed with %s", tc.name)
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

	categoryCount := 2
	apiCatRequests := []apitp.CategoryRequest{}
	apiCatResponses := []apitp.CategoryResponse{}
	for i := 0; i < categoryCount; i++ {
		c := setupApiCategoryRequest(i)
		cc, err := app.CreateCategory(c)
		if err != nil {
			t.Errorf("TestUpdateCategory: failed during setup. CreateCategory() failed: %v", err)
		}

		apiCatRequests = append(apiCatRequests, c)
		apiCatResponses = append(apiCatResponses, cc)
	}

	testCases := []struct {
		name          string
		categoryId    string
		category      apitp.CategoryRequest
		title         string
		shouldSucceed bool
	}{
		{"valid category", apiCatResponses[1].Id.String(), apiCatRequests[1], "valid title3", true},
		{"non-existent category", uuid.New().String(), apitp.CategoryRequest{}, "valid title3", false},

		{"invalid category ID", "invalid", apitp.CategoryRequest{}, "valid title3", false},
		{"nil category ID", uuid.Nil.String(), apitp.CategoryRequest{}, "valid title3", false},
		{"empty category ID", "", apitp.CategoryRequest{}, "valid title3", false},

		{"empty title", apiCatResponses[1].Id.String(), apiCatRequests[1], "", false},
		{"minimum length title", apiCatResponses[1].Id.String(), apiCatRequests[1], strings.Repeat("a", apitp.MinEntityTitleLength), true},
		{"maximum length title", apiCatResponses[1].Id.String(), apiCatRequests[1], strings.Repeat("a", apitp.MaxEntityTitleLength), true},
		{"too long title", apiCatResponses[1].Id.String(), apiCatRequests[1], strings.Repeat("a", apitp.MaxEntityTitleLength+1), false},
		{"conflicting title", apiCatResponses[2].Id.String(), apiCatRequests[2], apiCatResponses[1].Title, false},
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
		category      apitp.CategoryRequest
		title         string
		shouldSucceed bool
	}{
		{"valid category", setupApiCategoryRequest(1), "valid title", true},
		{"nil id", setupApiCategoryRequest(2), "valid title", false},

		{"empty title", setupApiCategoryRequest(3), "", false},
		{"minimum length title", setupApiCategoryRequest(4), strings.Repeat("a", apitp.MinEntityTitleLength), true},
		{"maximum length title", setupApiCategoryRequest(5), strings.Repeat("a", apitp.MaxEntityTitleLength), true},
		{"too long title", setupApiCategoryRequest(6), strings.Repeat("a", apitp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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

	c := setupApiCategoryRequest(1)
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

func setupApiCategoryRequest(identifier int) apitp.CategoryRequest {
	return apitp.CategoryRequest{
		Title:       fmt.Sprintf("Category %d", identifier),
		Description: utest.ToPtr(fmt.Sprintf("Category %d description", identifier)),
	}
}
