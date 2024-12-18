package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// Create an Asset
func (a *App) CreateAsset(groupTitle string, asset tp.Asset) (tp.Asset, error) {
	if _, err := a.GetGroup(groupTitle); err != nil {
		return tp.Asset{}, err
	}

	if asset.GroupTitle != groupTitle {
		return tp.Asset{}, ae.ErrGroupTitleMismatch
	}

	return a.db.CreateAsset(asset)
}

// Delete an Asset
func (a *App) DeleteAsset(groupTitle string, assetTitle string) error {
	// Get Asset will validate the group as well, Get before delete so we can return a not found error.
	if _, err := a.GetAsset(groupTitle, assetTitle); err != nil {
		return err
	}

	return a.db.DeleteAsset(groupTitle, assetTitle)
}

// List all Assets in a Group
func (a *App) ListAssets(groupTitle string) ([]tp.Asset, error) {
	if _, err := a.GetGroup(groupTitle); err != nil {
		return []tp.Asset{}, err
	}

	assets, err := a.db.ListAssetsByGroup(groupTitle)
	if err != nil {
		return []tp.Asset{}, err
	}

	return assets, nil
}

// Get an Asset
func (a *App) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	if _, err := a.GetGroup(groupTitle); err != nil {
		return tp.Asset{}, err
	}

	return a.db.GetAsset(groupTitle, assetTitle)
}

// Update an Asset
func (a *App) UpdateAsset(oldGroupTitle string, oldAssetTitle string, asset tp.Asset) (tp.Asset, error) {
	if _, err := a.GetGroup(oldGroupTitle); err != nil {
		return tp.Asset{}, err
	}

	if asset.GroupTitle == "" {
		asset.GroupTitle = oldGroupTitle
	}

	if asset.Title == "" {
		asset.Title = oldAssetTitle
	}

	return a.db.UpdateAsset(oldGroupTitle, oldAssetTitle, asset)
}
