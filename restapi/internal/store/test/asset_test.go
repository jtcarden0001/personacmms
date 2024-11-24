package test

import (
	"fmt"
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetCreate(t *testing.T) {
	store := InitializeStore("testassetcreate")

	// setup
	groupTitle, categoryTitle := setupAssetDependencies(t, store, "1")

	// Create
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	returnedAsset, err := store.CreateAsset(groupTitle, asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, asset, returnedAsset, fieldsToExclude)
}

func TestAssetDelete(t *testing.T) {
	store := InitializeStore("testassetdelete")

	// setup
	groupTitle, categoryTitle := setupAssetDependencies(t, store, "1")

	// Delete
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	_, err := store.CreateAsset(groupTitle, asset)
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
	store := InitializeStore("testassetlist")

	// setup
	groupTitle, categoryTitle := setupAssetDependencies(t, store, "1")

	// original List
	assets, err := store.ListAsset(groupTitle)
	if err != nil {
		t.Errorf("ListAsset() failed: %v", err)
	}

	// create a map of the assets title: tp.asset
	assetMap := make(map[string]tp.Asset)
	for _, asset := range assets {
		assetMap[asset.Title] = asset
	}

	// add 2 assets to the map and create them
	assetMap["testasset1"] = getTestAsset(groupTitle, categoryTitle, "1")
	assetMap["testasset2"] = getTestAsset(groupTitle, categoryTitle, "2")

	_, err = store.CreateAsset(groupTitle, assetMap["testasset1"])
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	_, err = store.CreateAsset(groupTitle, assetMap["testasset2"])
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	// List, compare the length to the original list+2 and compare the new listed assets with the expected assets
	assets, err = store.ListAsset(groupTitle)
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

func TestAssetUpdateGet(t *testing.T) {
	store := InitializeStore("testassetupdateget")

	// setup
	groupTitle, categoryTitle := setupAssetDependencies(t, store, "1")

	// Create
	asset := getTestAsset(groupTitle, categoryTitle, "1")
	createAsset, err := store.CreateAsset(groupTitle, asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	// Update
	asset.Title = "updated title"
	asset.Year = 2022
	updatedAsset, err := store.UpdateAsset(groupTitle, createAsset.Title, asset)
	if err != nil {
		t.Errorf("UpdateAsset() failed: %v", err)
	}

	differentFields := convertToSet([]string{"Title", "Year"})
	compEntitiesFieldsShouldBeDifferent(t, createAsset, updatedAsset, differentFields)

	// Get
	returnedAsset, err := store.GetAsset(groupTitle, updatedAsset.Title)
	if err != nil {
		t.Errorf("GetAsset() failed: %v", err)
	}

	compEntities(t, updatedAsset, returnedAsset)
}

func setupAssetDependencies(t *testing.T, store store.Store, suffix string) (string, string) {
	group := tp.Group{
		Title: fmt.Sprintf("testgroup%s", suffix),
	}
	_, err := store.CreateGroup(group)
	if err != nil {
		t.Errorf("CreateGroup() failed: %v", err)
	}

	category := tp.Category{
		Title:       fmt.Sprintf("testcategory%s", suffix),
		Description: fmt.Sprintf("test description %s", suffix),
	}
	_, err = store.CreateCategory(category)
	if err != nil {
		t.Errorf("CreateCategory() failed: %v", err)
	}

	return group.Title, category.Title
}

func getTestAsset(groupTitle string, categoryTitle string, suffix string) tp.Asset {
	return tp.Asset{
		Title:         fmt.Sprintf("testasset%s", suffix),
		GroupTitle:    groupTitle,
		Year:          2021,
		Description:   fmt.Sprintf("test description %s", suffix),
		Make:          fmt.Sprintf("test make %s", suffix),
		ModelNumber:   fmt.Sprintf("test model number %s", suffix),
		SerialNumber:  fmt.Sprintf("test serial number %s", suffix),
		CategoryTitle: categoryTitle,
	}
}
