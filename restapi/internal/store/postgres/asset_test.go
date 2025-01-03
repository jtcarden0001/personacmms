package postgres_test

import (
	"testing"

	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestAssetAssociateWithCategory(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetassociatewithcategory"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := utest.SetupAsset(1, true)
	createdAsset, err := store.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetAssociateWithCategory: failed during setup. CreateAsset() failed: %v", err)
	}

	c := utest.SetupCategory(1, true)
	createdCategory, err := store.CreateCategory(c)
	if err != nil {
		t.Errorf("TestAssetAssociateWithCategory: failed during setup. CreateCategory() failed: %v", err)
	}

	// test
	_, err = store.AssociateAssetWithCategory(createdAsset.Id, createdCategory.Id)
	if err != nil {
		t.Errorf("AssociateAssetWithCategory() failed: %v", err)
	}
}

// TODO: TestAssetAssociateWithGroup

func TestAssetCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := utest.SetupAsset(1, true)

	// test
	createdAsset, err := store.CreateAsset(a)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	utest.CompEntities(t, a, createdAsset)
}

func TestAssetDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := utest.SetupAsset(1, true)
	createdAsset, err := store.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetDelete: failed during setup. CreateAsset() failed: %v", err)
	}

	// test
	err = store.DeleteAsset(createdAsset.Id)
	if err != nil {
		t.Errorf("TestAssetDelete: DeleteAsset() failed: %v", err)
	}

	_, err = store.GetAsset(createdAsset.Id)
	if err == nil {
		t.Errorf("TestAssetDelete: GetAsset() returned nil error after deletion")
	}
}

// TODO: TestAssetDisassociateWithCategory

// TODO: TestAssetDisassociateWithGroup

func TestAssetGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := utest.SetupAsset(1, true)
	createAsset, err := store.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetGet: failed during setup. CreateAsset() failed: %v", err)
	}

	// test
	getAsset, err := store.GetAsset(createAsset.Id)
	if err != nil {
		t.Errorf("GetAsset() failed: %v", err)
	}

	utest.CompEntities(t, createAsset, getAsset)
}

func TestAssetList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetlist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a1 := utest.SetupAsset(1, true)
	a2 := utest.SetupAsset(2, true)
	a3 := utest.SetupAsset(3, true)

	_, err := store.CreateAsset(a1)
	if err != nil {
		t.Errorf("TestAssetList: failed during setup. CreateAsset() failed: %v", err)
	}
	_, err = store.CreateAsset(a2)
	if err != nil {
		t.Errorf("TestAssetList: failed during setup. CreateAsset() failed: %v", err)
	}
	_, err = store.CreateAsset(a3)
	if err != nil {
		t.Errorf("TestAssetList: failed during setup. CreateAsset() failed: %v", err)
	}

	// test
	assets, err := store.ListAssets()
	if err != nil {
		t.Errorf("ListAssets() failed: %v", err)
	}

	if len(assets) != 3 {
		t.Errorf("ListAssets() returned %d assets, expected 3", len(assets))
	}
}

// TODO: TestAssetListByCategory

// TODO: TestAssetListByCategoryAndGroup

// TODO: TestAssetListByGroup

func TestAssetUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := utest.SetupAsset(1, true)
	createdAsset, err := store.CreateAsset(a)
	if err != nil {
		t.Errorf("TestAssetUpdate: failed during setup. CreateAsset() failed: %v", err)
	}

	// test
	a.Title = "Updated Title"
	a.Description = utest.ToPtr("Updated Description")
	updatedAsset, err := store.UpdateAsset(a)
	if err != nil {
		t.Errorf("UpdateAsset() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title", "Description"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdAsset, updatedAsset, differentFields)
}
