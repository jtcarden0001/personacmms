package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Asset interface {
	CreateAsset(string, int, string, string, string, int) (int, error)
	DeleteAsset(int) error
	GetAllAsset() ([]tp.Asset, error)
	GetAsset(int) (tp.Asset, error)
	UpdateAsset(int, string, int, string, string, string, int) error
}

func (a *App) CreateAsset(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	return a.db.CreateAsset(title, year, make, modelNumber, description, categoryId)
}

func (a *App) DeleteAsset(id int) error {
	return a.db.DeleteAsset(id)
}

func (a *App) GetAllAsset() ([]tp.Asset, error) {
	return a.db.GetAllAsset()
}

func (a *App) GetAsset(id int) (tp.Asset, error) {
	return a.db.GetAsset(id)
}

func (a *App) UpdateAsset(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	return a.db.UpdateAsset(id, title, year, make, modelNumber, description, categoryId)
}
