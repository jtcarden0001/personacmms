package postgres_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupAsset(identifier int) tp.Asset {
	year := 2000 + (identifier % 22) // Random year based on identifier
	return tp.Asset{
		Id:           uuid.New(),
		Title:        fmt.Sprintf("Asset %d", identifier),
		Year:         utest.ToPtr(year),
		Manufacturer: utest.ToPtr(fmt.Sprintf("Asset %d manufacturer", identifier)),
		Make:         utest.ToPtr(fmt.Sprintf("Asset %d make", identifier)),
		ModelNumber:  utest.ToPtr(fmt.Sprintf("Asset %d model number", identifier)),
		SerialNumber: utest.ToPtr(fmt.Sprintf("Asset %d serial number", identifier)),
		Description:  utest.ToPtr(fmt.Sprintf("Asset %d description", identifier)),
	}
}

func TestAssetCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := setupAsset(1)

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

	a := setupAsset(1)
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

func TestAssetGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := setupAsset(1)
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

	a1 := setupAsset(1)
	a2 := setupAsset(2)
	a3 := setupAsset(3)

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

func TestAssetUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testassetupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	a := setupAsset(1)
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
