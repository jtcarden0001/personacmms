package cmmsapp

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "asset1", GroupTitle: "group1"}
	createdAsset, err := app.CreateAsset("group1", asset)
	assert.NoError(t, err)
	assert.Equal(t, asset.Title, createdAsset.Title)
	assert.Equal(t, asset.GroupTitle, createdAsset.GroupTitle)
}

func TestDeleteAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "asset1", GroupTitle: "group1"}
	db.CreateAsset(asset)

	err := app.DeleteAsset("group1", "asset1")
	assert.NoError(t, err)

	_, err = app.GetAsset("group1", "asset1")
	assert.Error(t, err)
}

func TestListAssets(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset1 := tp.Asset{Title: "asset1", GroupTitle: "group1"}
	asset2 := tp.Asset{Title: "asset2", GroupTitle: "group1"}
	db.CreateAsset(asset1)
	db.CreateAsset(asset2)

	assets, err := app.ListAssets("group1")
	assert.NoError(t, err)
	assert.Len(t, assets, 2)
}

func TestGetAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "asset1", GroupTitle: "group1"}
	db.CreateAsset(asset)

	retrievedAsset, err := app.GetAsset("group1", "asset1")
	assert.NoError(t, err)
	assert.Equal(t, asset.Title, retrievedAsset.Title)
	assert.Equal(t, asset.GroupTitle, retrievedAsset.GroupTitle)
}

func TestUpdateAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "asset1", GroupTitle: "group1"}
	db.CreateAsset(asset)

	updatedAsset := tp.Asset{Title: "asset1_updated", GroupTitle: "group1"}
	_, err := app.UpdateAsset("group1", "asset1", updatedAsset)
	assert.NoError(t, err)

	retrievedAsset, err := app.GetAsset("group1", "asset1_updated")
	assert.NoError(t, err)
	assert.Equal(t, updatedAsset.Title, retrievedAsset.Title)
	assert.Equal(t, updatedAsset.GroupTitle, retrievedAsset.GroupTitle)
}

func TestCreateAssetWithEmptyTitle(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "", GroupTitle: "group1"}
	_, err := app.CreateAsset("group1", asset)
	assert.Error(t, err)
}

func TestCreateAssetWithNonExistentGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	asset := tp.Asset{Title: "asset1", GroupTitle: "nonexistentgroup"}
	_, err := app.CreateAsset("nonexistentgroup", asset)
	assert.Error(t, err)
}

func TestDeleteNonExistentAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	err := app.DeleteAsset("group1", "nonexistentasset")
	assert.Error(t, err)
}

func TestGetNonExistentAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	_, err := app.GetAsset("group1", "nonexistentasset")
	assert.Error(t, err)
}

func TestUpdateNonExistentAsset(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	updatedAsset := tp.Asset{Title: "asset1_updated", GroupTitle: "group1"}
	_, err := app.UpdateAsset("group1", "nonexistentasset", updatedAsset)
	assert.Error(t, err)
}

func TestCreateAssetWithMismatchedGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Id: uuid.New(), Title: "group1"}
	db.CreateGroup(group)

	asset := tp.Asset{Title: "asset1", GroupTitle: "group2"}
	_, err := app.CreateAsset("group1", asset)
	assert.Error(t, err)
}
