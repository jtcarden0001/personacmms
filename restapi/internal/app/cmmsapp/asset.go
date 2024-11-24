package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateAsset(groupTitle string, asset tp.Asset) (tp.Asset, error) {
	// TODO: validate group

	return a.db.CreateAsset(asset)
}

func (a *App) DeleteAsset(groupTitle string, assetTitle string) error {
	// TODO: validate group

	return a.db.DeleteAsset(groupTitle, assetTitle)
}

func (a *App) ListAssets(groupTitle string) ([]tp.Asset, error) {
	// TODO: validate group

	assets, err := a.db.ListAssets()
	if err != nil {
		return []tp.Asset{}, err
	}

	// filter assets by group title
	var groupAssets []tp.Asset
	for _, asset := range assets {
		if asset.GroupTitle == groupTitle {
			groupAssets = append(groupAssets, asset)
		}
	}

	return groupAssets, nil
}

func (a *App) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	return a.db.GetAsset(groupTitle, assetTitle)
}

func (a *App) UpdateAsset(groupTitle string, assetTitle string, asset tp.Asset) (tp.Asset, error) {
	return a.db.UpdateAsset(groupTitle, assetTitle, asset)
}
