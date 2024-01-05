package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type EquipmentCategory interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategory() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error
}

func (a *App) CreateEquipmentCategory(title string) (int, error) {
	return a.db.CreateEquipmentCategory(title)
}

func (a *App) DeleteEquipmentCategory(id int) error {
	return a.db.DeleteEquipmentCategory(id)
}

func (a *App) GetAllEquipmentCategory() ([]tp.EquipmentCategory, error) {
	return a.db.GetAllEquipmentCategory()
}

func (a *App) GetEquipmentCategory(id int) (tp.EquipmentCategory, error) {
	return a.db.GetEquipmentCategory(id)
}

func (a *App) UpdateEquipmentCategory(id int, title string) error {
	return a.db.UpdateEquipmentCategory(id, title)
}
