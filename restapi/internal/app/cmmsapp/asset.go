package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateAsset(groupTitle string, asset tp.Asset) (tp.Asset, error) {
	// interpolate args into target of creation
	gp, err := a.GetGroup(groupTitle)
	if err != nil {
		return tp.Asset{}, err
	}

	asset.GroupTitle = gp.Title
	return a.db.CreateAsset(asset)
}

func (a *App) DeleteAsset(groupTitle string, assetTitle string) error {
	// validate
	if _, err := a.GetGroup(groupTitle); err != nil {
		return err
	}

	if _, err := a.GetAsset(groupTitle, assetTitle); err != nil {
		return err
	}

	return a.db.DeleteAsset(groupTitle, assetTitle)
}

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

func (a *App) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	if _, err := a.GetGroup(groupTitle); err != nil {
		return tp.Asset{}, err
	}

	return a.db.GetAsset(groupTitle, assetTitle)
}

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
