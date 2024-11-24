package test

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestAssetTaskCreate(t *testing.T) {
	store := InitializeStore("testassettaskcreate")

	// setup
	groupTitle, categoryTitle := setupAssetDependencies(t, store, "1")
	assetId := setupAssetTaskDependencies(t, store, groupTitle, categoryTitle, "1")

	// Create
	at := tp.AssetTask{
		Title:              "testassettask1",
		AssetId:            assetId,
		UniqueInstructions: "test instructions",
	}
	returnedAssetTask, err := store.CreateAssetTask(groupTitle, categoryTitle, at)
	if err != nil {
		t.Errorf("CreateAssetTask() failed: %v", err)
	}

	compareEntitiesExcludingId(t, at, returnedAssetTask)

}

func setupAssetTaskDependencies(t *testing.T, store store.Store, groupTitle string, categoryTitle string, assetTitle string) tp.UUID {
	// create an asset
	asset := getTestAsset(groupTitle, categoryTitle, assetTitle)
	asset, err := store.CreateAsset(groupTitle, asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}

	return asset.Id
}
