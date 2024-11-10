package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type AssetCategory interface {
	CreateAssetCategory(string) (int, error)
	DeleteAssetCategory(int) error
	GetAllAssetCategory() ([]tp.AssetCategory, error)
	GetAssetCategory(int) (tp.AssetCategory, error)
	UpdateAssetCategory(int, string) error
}

func (a *App) CreateAssetCategory(title string) (int, error) {
	return a.db.CreateAssetCategory(title)
}

func (a *App) DeleteAssetCategory(id int) error {
	return a.db.DeleteAssetCategory(id)
}

func (a *App) GetAllAssetCategory() ([]tp.AssetCategory, error) {
	return a.db.GetAllAssetCategory()
}

func (a *App) GetAssetCategory(id int) (tp.AssetCategory, error) {
	return a.db.GetAssetCategory(id)
}

func (a *App) UpdateAssetCategory(id int, title string) error {
	return a.db.UpdateAssetCategory(id, title)
}
