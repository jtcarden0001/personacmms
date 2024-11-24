package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Asset interface {
	CreateAsset(string, tp.Asset) (tp.Asset, error)
	DeleteAsset(string, string) error
	ListAsset(string) ([]tp.Asset, error)
	GetAsset(string, string) (tp.Asset, error)
	UpdateAsset(string, string, tp.Asset) (tp.Asset, error)
}

func (a *App) CreateAsset(groupTitle string, asset tp.Asset) (tp.Asset, error) {
	return a.db.CreateAsset(groupTitle, asset)
}

func (a *App) DeleteAsset(groupTitle string, assetTitle string) error {
	return a.db.DeleteAsset(groupTitle, assetTitle)
}

func (a *App) ListAsset(groupTitle string) ([]tp.Asset, error) {
	return a.db.ListAsset(groupTitle)
}

func (a *App) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	return a.db.GetAsset(groupTitle, assetTitle)
}

func (a *App) UpdateAsset(groupTitle string, assetTitle string, asset tp.Asset) (tp.Asset, error) {
	return a.db.UpdateAsset(groupTitle, assetTitle, asset)
}
