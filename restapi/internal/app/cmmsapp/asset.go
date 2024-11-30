package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateAsset(groupTitle string, asset tp.Asset) (tp.Asset, error) {
	gp, err := a.GetGroup(groupTitle)
	if err != nil {
		return tp.Asset{}, err
	}

	asset.GroupTitle = gp.Title
	return a.db.CreateAsset(asset)
}

func (a *App) DeleteAsset(groupTitle string, assetTitle string) error {
	_, err := a.GetGroup(groupTitle)
	if err != nil {
		return err
	}

	_, err = a.GetAsset(groupTitle, assetTitle)
	if err != nil {
		return err
	}

	return a.db.DeleteAsset(groupTitle, assetTitle)
}

func (a *App) ListAssets(groupTitle string) ([]tp.Asset, error) {
	_, err := a.GetGroup(groupTitle)
	if err != nil {
		return []tp.Asset{}, err
	}

	assets, err := a.db.ListAssetsByGroup(groupTitle)
	if err != nil {
		return []tp.Asset{}, err
	}

	return assets, nil
}

func (a *App) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	_, err := a.GetGroup(groupTitle)
	if err != nil {
		return tp.Asset{}, err
	}

	return a.db.GetAsset(groupTitle, assetTitle)
}

func (a *App) UpdateAsset(groupTitle string, assetTitle string, asset tp.Asset) (tp.Asset, error) {
	_, err := a.GetGroup(groupTitle)
	if err != nil {
		return tp.Asset{}, err
	}

	return a.db.UpdateAsset(groupTitle, assetTitle, asset)
}
