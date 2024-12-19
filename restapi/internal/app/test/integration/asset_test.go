package integration

// func TestCreateAsset(t *testing.T) {
// 	app := newApp("testcreateasset")
// 	asset := tp.Asset{Title: "Test Asset", GroupTitle: "Test Group"}
// 	createdAsset, err := app.CreateAsset("Test Group", asset)
// 	assert.NoError(t, err)
// 	assert.Equal(t, asset.Title, createdAsset.Title)
// }

// func TestDeleteAsset(t *testing.T) {
// 	app := newApp("testdeleteasset")
// 	asset := tp.Asset{Title: "Test Asset", GroupTitle: "Test Group"}
// 	_, _ = app.CreateAsset("Test Group", asset)
// 	err := app.DeleteAsset("Test Group", asset.Title)
// 	assert.NoError(t, err)
// }

// func TestListAssets(t *testing.T) {
// 	app := newApp("testlistassets")
// 	asset := tp.Asset{Title: "Test Asset", GroupTitle: "Test Group"}
// 	_, _ = app.CreateAsset("Test Group", asset)
// 	assets, err := app.ListAssets("Test Group")
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, assets)
// }

// func TestGetAsset(t *testing.T) {
// 	app := newApp("testgetasset")
// 	asset := tp.Asset{Title: "Test Asset", GroupTitle: "Test Group"}
// 	_, _ = app.CreateAsset("Test Group", asset)
// 	retrievedAsset, err := app.GetAsset("Test Group", asset.Title)
// 	assert.NoError(t, err)
// 	assert.Equal(t, asset.Title, retrievedAsset.Title)
// }

// func TestUpdateAsset(t *testing.T) {
// 	app := newApp("testupdateasset")
// 	asset := tp.Asset{Title: "Test Asset", GroupTitle: "Test Group"}
// 	_, _ = app.CreateAsset("Test Group", asset)
// 	updatedAsset := tp.Asset{Title: "Updated Asset", GroupTitle: "Test Group"}
// 	_, err := app.UpdateAsset("Test Group", asset.Title, updatedAsset)
// 	assert.NoError(t, err)
// 	retrievedAsset, _ := app.GetAsset("Test Group", updatedAsset.Title)
// 	assert.Equal(t, updatedAsset.Title, retrievedAsset.Title)
// }
