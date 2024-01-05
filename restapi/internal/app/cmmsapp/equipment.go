package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Equipment interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
}

func (a *App) CreateEquipment(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	return a.db.CreateEquipment(title, year, make, modelNumber, description, categoryId)
}

func (a *App) DeleteEquipment(id int) error {
	return a.db.DeleteEquipment(id)
}

func (a *App) GetAllEquipment() ([]tp.Equipment, error) {
	return a.db.GetAllEquipment()
}

func (a *App) GetEquipment(id int) (tp.Equipment, error) {
	return a.db.GetEquipment(id)
}

func (a *App) UpdateEquipment(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	return a.db.UpdateEquipment(id, title, year, make, modelNumber, description, categoryId)
}
