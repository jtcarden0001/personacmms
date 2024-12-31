package cmmsapp

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssociateAssetWithCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateAssetWithCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateAssetWithCategory: failed during setup. CreateAsset() failed: %v", err)
	}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestAssociateAssetWithCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		categoryID    string
		shouldSucceed bool
	}{
		{"valid association", createdAsset.Id.String(), createdCategory.Id.String(), true},
		{"invalid category ID", createdAsset.Id.String(), "invalid", false},
		{"invalid asset ID", "invalid", createdCategory.Id.String(), false},
		{"invalid asset and category ID", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdCategory.Id.String(), false},
		{"nil category ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"nil asset and category ID", uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdCategory.Id.String(), false},
		{"empty category ID", createdAsset.Id.String(), "", false},
		{"empty asset and category ID", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.AssociateAssetWithCategory(tc.assetID, tc.categoryID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("AssociateAssetWithCategory() failed: %v", err)
			}
			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateAssetWithCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestAssociateAssetWithGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssociateAssetWithGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssociateAssetWithGroup: failed during setup. CreateAsset() failed: %v", err)
	}

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestAssociateAssetWithGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		groupID       string
		shouldSucceed bool
	}{
		{"valid association", createdAsset.Id.String(), createdGroup.Id.String(), true},
		{"invalid group ID", createdAsset.Id.String(), "invalid", false},
		{"invalid asset ID", "invalid", createdGroup.Id.String(), false},
		{"invalid asset and group ID", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdGroup.Id.String(), false},
		{"nil group ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"nil asset and group ID", uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdGroup.Id.String(), false},
		{"empty group ID", createdAsset.Id.String(), "", false},
		{"empty asset and group ID", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.AssociateAssetWithGroup(tc.assetID, tc.groupID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("AssociateAssetWithGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("AssociateAssetWithGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestCreateAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestCreateAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	conflictingAsset := utest.SetupAsset(1, false)
	_, err = app.CreateAsset(conflictingAsset)
	if err != nil {
		t.Errorf("TestCreateAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	emtyTitleAsset := utest.SetupAsset(2, false)
	emtyTitleAsset.Title = ""

	testCases := []struct {
		name          string
		asset         tp.Asset
		shouldSucceed bool
	}{
		{"valid asset", utest.SetupAsset(3, false), true},
		{"non nil id", utest.SetupAsset(4, true), false},
		{"empty title", emtyTitleAsset, false},
		{"conflicting title", conflictingAsset, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.CreateAsset(tc.asset)
			if tc.shouldSucceed && err != nil {
				t.Errorf("CreateAsset() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("CreateAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDeleteAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDeleteAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDeleteAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetId       string
		shouldSucceed bool
	}{
		{"valid asset deletion", createdAsset.Id.String(), true},
		{"invalid asset ID", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), false},
		{"empty asset ID", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DeleteAsset(tc.assetId)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DeleteAsset() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DeleteAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateAssetWithCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateAssetWithCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithCategory: failed during setup. CreateAsset() failed: %v", err)
	}

	c := utest.SetupCategory(1, false)
	createdCategory1, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	c = utest.SetupCategory(2, false)
	createdCategory2, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	_, err = app.AssociateAssetWithCategory(createdAsset.Id.String(), createdCategory1.Id.String())
	if err != nil {
		t.Errorf("TestDisassociateAssetWithCategory: failed during setup. AssociateAssetWithCategory() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		categoryID    string
		shouldSucceed bool
	}{
		{"valid disassociation", createdAsset.Id.String(), createdCategory1.Id.String(), true},
		{"invalid category ID", createdAsset.Id.String(), "invalid", false},
		{"invalid asset ID", "invalid", createdCategory1.Id.String(), false},
		{"invalid asset and category ID", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdCategory1.Id.String(), false},
		{"nil category ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"nil asset and category ID", uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdCategory1.Id.String(), false},
		{"empty category ID", createdAsset.Id.String(), "", false},
		{"empty asset and category ID", "", "", false},
		{"non-existent asset", uuid.New().String(), createdCategory1.Id.String(), false},
		{"non-existent category", createdAsset.Id.String(), uuid.New().String(), false},
		{"non-existent asset and category", uuid.New().String(), uuid.New().String(), false},
		{"disassociate call with no association", createdAsset.Id.String(), createdCategory2.Id.String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateAssetWithCategory(tc.assetID, tc.categoryID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateAssetWithCategory() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateAssetWithCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestDisassociateAssetWithGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestDisassociateAssetWithGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithGroup: failed during setup. CreateAsset() failed: %v", err)
	}

	g := utest.SetupGroup(1, false)
	createdGroup1, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	g = utest.SetupGroup(2, false)
	createdGroup2, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestDisassociateAssetWithGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	_, err = app.AssociateAssetWithGroup(createdAsset.Id.String(), createdGroup1.Id.String())
	if err != nil {
		t.Errorf("TestDisassociateAssetWithGroup: failed during setup. AssociateAssetWithGroup() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		groupID       string
		shouldSucceed bool
	}{
		{"valid disassociation", createdAsset.Id.String(), createdGroup1.Id.String(), true},
		{"invalid group ID", createdAsset.Id.String(), "invalid", false},
		{"invalid asset ID", "invalid", createdGroup1.Id.String(), false},
		{"invalid asset and group ID", "invalid", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), createdGroup1.Id.String(), false},
		{"nil group ID", createdAsset.Id.String(), uuid.Nil.String(), false},
		{"nil asset and group ID", uuid.Nil.String(), uuid.Nil.String(), false},
		{"empty asset ID", "", createdGroup1.Id.String(), false},
		{"empty group ID", createdAsset.Id.String(), "", false},
		{"empty asset and group ID", "", "", false},
		{"non-existent asset", uuid.New().String(), createdGroup1.Id.String(), false},
		{"non-existent group", createdAsset.Id.String(), uuid.New().String(), false},
		{"non-existent asset and group", uuid.New().String(), uuid.New().String(), false},
		{"disassociate call with no association", createdAsset.Id.String(), createdGroup2.Id.String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.DisassociateAssetWithGroup(tc.assetID, tc.groupID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("DisassociateAssetWithGroup() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("DisassociateAssetWithGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestGetAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestGetAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestGetAsset: failed during setup. CreateAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		shouldSucceed bool
	}{
		{"valid asset", createdAsset.Id.String(), true},
		{"invalid asset ID", "invalid", false},
		{"nil asset ID", uuid.Nil.String(), false},
		{"empty asset ID", "", false},
		{"non-existent asset", uuid.New().String(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := app.GetAsset(tc.assetID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("GetAsset() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("GetAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListAssets(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListAssets")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	_, err = app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListAssets: failed during setup. CreateAsset() failed: %v", err)
	}

	a = utest.SetupAsset(2, false)
	_, err = app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestListAssets: failed during setup. CreateAsset() failed: %v", err)
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
			as, err := app.ListAssets()
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListAssets() failed: %v", err)
				} else {
					if len(as) != tc.count {
						t.Errorf("ListAssets() failed: expected %d assets, got %d", tc.count, len(as))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListAssets() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListAssetByCategory(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListAssetByCategory")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	caCount := 5
	associatedCaCount := 3
	var createdAssets []tp.Asset
	for i := 0; i < caCount; i++ {
		a := utest.SetupAsset(i, false)
		createdAsset, err := app.CreateAsset(a)
		if err != nil {
			t.Errorf("TestListAssetByCategory: failed during setup. CreateAsset() failed: %v", err)
		}
		createdAssets = append(createdAssets, createdAsset)
	}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListAssetByCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	for i := 0; i < associatedCaCount; i++ {
		_, err := app.AssociateAssetWithCategory(createdAssets[i].Id.String(), createdCategory.Id.String())
		if err != nil {
			t.Errorf("TestListAssetByCategory: failed during setup. AssociateAssetWithCategory() failed: %v", err)
		}
	}

	testCases := []struct {
		name          string
		categoryID    string
		count         int
		shouldSucceed bool
	}{
		{"valid list", createdCategory.Id.String(), associatedCaCount, true},
		{"invalid category ID", "invalid", 0, false},
		{"nil category ID", uuid.Nil.String(), 0, false},
		{"empty category ID", "", 0, false},
		{"non-existent category", uuid.New().String(), 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			as, err := app.ListAssetsByCategory(tc.categoryID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListAssetsByCategory() failed: %v", err)
				} else {
					if len(as) != tc.count {
						t.Errorf("ListAssetsByCategory() failed: expected %d assets, got %d", tc.count, len(as))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListAssetsByCategory() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListAssetByCategoryAndGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListAssetByCategoryAndGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	aCount := 5
	associatedAcCount := 4
	associatedAgCount := 3 // limiting factor
	var createdAssets []tp.Asset
	for i := 0; i < aCount; i++ {
		a := utest.SetupAsset(i, false)
		createdAsset, err := app.CreateAsset(a)
		if err != nil {
			t.Errorf("TestListAssetByCategoryAndGroup: failed during setup. CreateAsset() failed: %v", err)
		}
		createdAssets = append(createdAssets, createdAsset)
	}

	c := utest.SetupCategory(1, false)
	createdCategory, err := app.CreateCategory(c)
	if err != nil {
		t.Errorf("TestListAssetByCategoryAndGroup: failed during setup. CreateCategory() failed: %v", err)
	}

	for i := 0; i < associatedAcCount; i++ {
		_, err := app.AssociateAssetWithCategory(createdAssets[i].Id.String(), createdCategory.Id.String())
		if err != nil {
			t.Errorf("TestListAssetByCategoryAndGroup: failed during setup. AssociateAssetWithCategory() failed: %v", err)
		}
	}

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestListAssetByCategoryAndGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	for i := 0; i < associatedAgCount; i++ {
		_, err := app.AssociateAssetWithGroup(createdAssets[i].Id.String(), createdGroup.Id.String())
		if err != nil {
			t.Errorf("TestListAssetByCategoryAndGroup: failed during setup. AssociateAssetWithGroup() failed: %v", err)
		}
	}

	testCases := []struct {
		name          string
		categoryID    string
		groupID       string
		count         int
		shouldSucceed bool
	}{
		{"valid list", createdCategory.Id.String(), createdGroup.Id.String(), associatedAgCount, true},
		{"invalid category ID", "invalid", createdGroup.Id.String(), 0, false},
		{"invalid group ID", createdCategory.Id.String(), "invalid", 0, false},
		{"invalid category and group ID", "invalid", "invalid", 0, false},
		{"nil category ID", uuid.Nil.String(), createdGroup.Id.String(), 0, false},
		{"nil group ID", createdCategory.Id.String(), uuid.Nil.String(), 0, false},
		{"nil category and group ID", uuid.Nil.String(), uuid.Nil.String(), 0, false},
		{"empty category ID", "", createdGroup.Id.String(), 0, false},
		{"empty group ID", createdCategory.Id.String(), "", 0, false},
		{"empty category and group ID", "", "", 0, false},
		{"non-existent category", uuid.New().String(), createdGroup.Id.String(), 0, false},
		{"non-existent group", createdCategory.Id.String(), uuid.New().String(), 0, false},
		{"non-existent category and group", uuid.New().String(), uuid.New().String(), 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			as, err := app.ListAssetsByCategoryAndGroup(tc.categoryID, tc.groupID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListAssetsByCategoryAndGroup() failed: %v", err)
				} else {
					if len(as) != tc.count {
						t.Errorf("ListAssetsByCategoryAndGroup() failed: expected %d assets, got %d", tc.count, len(as))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListAssetsByCategoryAndGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestListAssetsByGroup(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestListAssetsByGroup")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	aCount := 5
	associatedACount := 3
	var createdAssets []tp.Asset
	for i := 0; i < aCount; i++ {
		a := utest.SetupAsset(i, false)
		createdAsset, err := app.CreateAsset(a)
		if err != nil {
			t.Errorf("TestListAssetsByGroup: failed during setup. CreateAsset() failed: %v", err)
		}
		createdAssets = append(createdAssets, createdAsset)
	}

	g := utest.SetupGroup(1, false)
	createdGroup, err := app.CreateGroup(g)
	if err != nil {
		t.Errorf("TestListAssetsByGroup: failed during setup. CreateGroup() failed: %v", err)
	}

	for i := 0; i < associatedACount; i++ {
		_, err := app.AssociateAssetWithGroup(createdAssets[i].Id.String(), createdGroup.Id.String())
		if err != nil {
			t.Errorf("TestListAssetsByGroup: failed during setup. AssociateAssetWithGroup() failed: %v", err)
		}
	}

	testCases := []struct {
		name          string
		groupID       string
		count         int
		shouldSucceed bool
	}{
		{"valid list", createdGroup.Id.String(), associatedACount, true},
		{"invalid group ID", "invalid", 0, false},
		{"nil group ID", uuid.Nil.String(), 0, false},
		{"empty group ID", "", 0, false},
		{"non-existent group", uuid.New().String(), 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			as, err := app.ListAssetsByGroup(tc.groupID)
			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("ListAssetsByGroup() failed: %v", err)
				} else {
					if len(as) != tc.count {
						t.Errorf("ListAssetsByGroup() failed: expected %d assets, got %d", tc.count, len(as))
					}
				}
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("ListAssetsByGroup() should have failed with %s", tc.name)
			}
		})
	}
}

func TestUpdateAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestUpdateAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	assetCount := 5
	var ids []string
	assets := make(map[string]tp.Asset)
	nilIdAssets := make(map[string]tp.Asset)
	for i := 0; i < assetCount; i++ {
		a := utest.SetupAsset(i, false)
		ca, err := app.CreateAsset(a)
		if err != nil {
			t.Errorf("TestUpdateAsset: failed during setup. CreateAsset() failed: %v", err)
		}

		ids = append(ids, ca.Id.String())
		assets[ca.Id.String()] = ca
		nilIdAssets[ca.Id.String()] = a
	}

	testCases := []struct {
		name          string
		assetId       string
		asset         tp.Asset
		title         string
		shouldSucceed bool
	}{
		{"valid asset with matching IDs", ids[0], assets[ids[0]], "valid title1", true},
		{"valid asset with Asset.Id nil", ids[1], nilIdAssets[ids[1]], "valid title2", true},
		{"mismatching asset ID and Asset.iD", ids[2], assets[ids[3]], "valid title3", false},
		{"non-existent asset", uuid.New().String(), tp.Asset{}, "valid title3", false},

		{"invalid asset ID", "invalid", tp.Asset{}, "valid title3", false},
		{"nil asset ID", uuid.Nil.String(), tp.Asset{}, "valid title3", false},
		{"empty asset ID", "", tp.Asset{}, "valid title3", false},
		{"conflicting id", ids[4], assets[ids[3]], "valid title3", false},

		{"empty title", ids[1], assets[ids[1]], "", false},
		{"minimum length title", ids[1], assets[ids[1]], strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", ids[1], assets[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", ids[1], assets[ids[1]], strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
		{"conflicting title", ids[2], assets[ids[2]], assets[ids[3]].Title, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.asset.Title = tc.title
			_, err := app.UpdateAsset(tc.assetId, tc.asset)
			if tc.shouldSucceed && err != nil {
				t.Errorf("UpdateAsset() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("UpdateAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestValidateAsset(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestValidateAsset")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	testCases := []struct {
		name          string
		asset         tp.Asset
		id            uuid.UUID
		title         string
		shouldSucceed bool
	}{
		{"valid asset", utest.SetupAsset(1, false), uuid.New(), "valid title", true},
		{"nil id", utest.SetupAsset(2, false), uuid.Nil, "valid title", false},

		{"empty title", utest.SetupAsset(3, false), uuid.New(), "", false},
		{"minimum length title", utest.SetupAsset(4, false), uuid.New(), strings.Repeat("a", tp.MinEntityTitleLength), true},
		{"maximum length title", utest.SetupAsset(5, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength), true},
		{"too long title", utest.SetupAsset(6, false), uuid.New(), strings.Repeat("a", tp.MaxEntityTitleLength+1), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.asset.Id = tc.id
			tc.asset.Title = tc.title
			err := app.validateAsset(tc.asset)
			if tc.shouldSucceed && err != nil {
				t.Errorf("validateAsset() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("validateAsset() should have failed with %s", tc.name)
			}
		})
	}
}

func TestAssetExists(t *testing.T) {
	t.Parallel()
	app, cleanup, err := initializeAppTest(t, "TestAssetExists")
	if err != nil {
		t.Fatalf("Could not initialize app: %s", err)
	}
	defer cleanup()

	a := utest.SetupAsset(1, false)
	createdAsset, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetExists: failed during setup. CreateAsset() failed: %v", err)
	}

	testCases := []struct {
		name          string
		assetID       string
		shouldExist   bool
		shouldSucceed bool
	}{
		{"valid asset", createdAsset.Id.String(), true, true},
		{"non-existent asset", uuid.New().String(), false, true},

		{"invalid asset ID", "invalid", false, false},
		{"nil asset ID", uuid.Nil.String(), false, false},
		{"empty asset ID", "", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, exists, err := app.assetExists(tc.assetID)
			if tc.shouldSucceed && err != nil {
				t.Errorf("assetExists() failed: %v", err)
			}

			if !tc.shouldSucceed && err == nil {
				t.Errorf("assetExists() should have failed with %s", tc.name)
			}

			if exists != tc.shouldExist {
				t.Errorf("assetExists() failed: expected %t, got %t", tc.shouldExist, exists)
			}
		})
	}
}
