package integration

import (
	"fmt"
	"strconv"
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetCreate(t *testing.T) {
	t.Parallel()
	dbName := "testassetcreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	groupTitle := setupGroup(t, store, "1")
	categoryTitle := setupCategory(t, store, "1")

	// Create
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	returnedAsset, err := store.CreateAsset(asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, asset, returnedAsset, fieldsToExclude)
}

func TestAssetDelete(t *testing.T) {
	t.Parallel()
	dbName := "testassetdelete"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	groupTitle := setupGroup(t, store, "1")
	categoryTitle := setupCategory(t, store, "1")

	// Delete
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	_, err := store.CreateAsset(asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	err = store.DeleteAsset(groupTitle, asset.Title)
	if err != nil {
		t.Errorf("DeleteAsset() failed: %v", err)
	}

	// Get
	_, err = store.GetAsset(groupTitle, asset.Title)
	if err == nil {
		t.Errorf("GetAsset() failed: expected error, got nil")
	}
}

func TestAssetList(t *testing.T) {
	t.Parallel()
	dbName := "testassetlist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	groupTitle := setupGroup(t, store, "1")
	categoryTitle := setupCategory(t, store, "1")

	// original List
	assets, err := store.ListAssets()
	if err != nil {
		t.Errorf("ListAsset() failed: %v", err)
	}

	// create a map of the assets title: tp.asset
	assetMap := make(map[string]tp.Asset)
	for _, asset := range assets {
		assetMap[asset.Title] = asset
	}

	// add 2 assets to the map and create them
	asset1 := getTestAsset(groupTitle, categoryTitle, "1")
	createAsset1, err := store.CreateAsset(asset1)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}
	assetMap[createAsset1.Title] = createAsset1

	asset2 := getTestAsset(groupTitle, categoryTitle, "2")
	createAsset2, err := store.CreateAsset(asset2)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}
	assetMap[createAsset2.Title] = createAsset2

	// List, compare the length to the original list+2 and compare the new listed assets with the expected assets
	assets, err = store.ListAssets()
	if err != nil {
		t.Errorf("ListAsset() failed: %v", err)
	}

	if len(assets) != len(assetMap) {
		t.Errorf("ListAsset() failed: expected %d assets, got %d", len(assetMap), len(assets))
	}

	newAssetMap := make(map[string]tp.Asset)
	for _, asset := range assets {
		newAssetMap[asset.Title] = asset
	}

	for title, asset := range assetMap {
		fieldsToExclude := convertToSet([]string{"Id"})
		compEntitiesExcludeFields(t, asset, newAssetMap[title], fieldsToExclude)
	}
}

func TestAssetListByGroup(t *testing.T) {
	t.Parallel()
	dbName := "testassetlistbygroup"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	groupTitle1 := setupGroup(t, store, "1")
	groupTitle2 := setupGroup(t, store, "2")

	// create 2 assets in group 1 and 2 assets in group 2
	categoryTitle := setupCategory(t, store, "1")
	asset1 := getTestAsset(groupTitle1, categoryTitle, "1")
	_, err := store.CreateAsset(asset1)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	asset2 := getTestAsset(groupTitle1, categoryTitle, "2")
	_, err = store.CreateAsset(asset2)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	asset3 := getTestAsset(groupTitle2, categoryTitle, "1")
	_, err = store.CreateAsset(asset3)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	asset4 := getTestAsset(groupTitle2, categoryTitle, "2")
	_, err = store.CreateAsset(asset4)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	// List by group 1
	assets, err := store.ListAssetsByGroup(groupTitle1)
	if err != nil {
		t.Errorf("ListAssetsByGroup() failed: %v", err)
	}

	if len(assets) != 2 {
		t.Errorf("ListAssetsByGroup() failed: expected 2 assets, got %d", len(assets))
	}

	// List by group 2
	assets, err = store.ListAssetsByGroup(groupTitle2)
	if err != nil {
		t.Errorf("ListAssetsByGroup() failed: %v", err)
	}

	if len(assets) != 2 {
		t.Errorf("ListAssetsByGroup() failed: expected 2 assets, got %d", len(assets))
	}

	// List all
	assets, err = store.ListAssets()
	if err != nil {
		t.Errorf("ListAssets() failed: %v", err)
	}

	if len(assets) != 4 {
		t.Errorf("ListAssets() failed: expected 4 assets, got %d", len(assets))
	}
}

func TestAssetUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testassetupdateget"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// setup
	groupTitle := setupGroup(t, store, "1")
	categoryTitle := setupCategory(t, store, "1")

	// Create
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	createAsset, err := store.CreateAsset(asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	// Update
	asset = getTestAsset(groupTitle, categoryTitle, "2")
	asset.Id = createAsset.Id
	updatedAsset, err := store.UpdateAsset(groupTitle, createAsset.Title, asset)
	if err != nil {
		t.Errorf("UpdateAsset() failed: %v", err)
	}

	differentFields := convertToSet([]string{"Title", "Year", "Make", "ModelNumber", "SerialNumber", "Description"})
	compEntitiesFieldsShouldBeDifferent(t, createAsset, updatedAsset, differentFields)

	// Get
	returnedAsset, err := store.GetAsset(groupTitle, updatedAsset.Title)
	if err != nil {
		t.Errorf("GetAsset() failed: %v", err)
	}

	compEntities(t, updatedAsset, returnedAsset)
}

func getTestAsset(groupTitle string, categoryTitle string, suffix string) tp.Asset {
	return tp.Asset{
		GroupTitle:    groupTitle,
		Title:         fmt.Sprintf("testasset%s", suffix),
		Year:          func() *int { year, _ := strconv.Atoi(fmt.Sprintf("202%s", suffix)); return &year }(),
		Description:   toPtr(fmt.Sprintf("test description %s", suffix)),
		Make:          toPtr(fmt.Sprintf("test make %s", suffix)),
		ModelNumber:   toPtr(fmt.Sprintf("test model number %s", suffix)),
		SerialNumber:  toPtr(fmt.Sprintf("test serial number %s", suffix)),
		CategoryTitle: &categoryTitle,
	}
}
