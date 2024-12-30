package cmmsapp

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssociateAssetWithCategory(t *testing.T) {
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
