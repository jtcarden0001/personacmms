package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type Equipment interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
}

type EquipmentTest interface {
	ResetSequenceEquipment(int) error
}

func (cmms *App) CreateEquipment(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	return cmms.db.CreateEquipment(title, year, make, modelNumber, description, categoryId)
}

func (cmms *App) DeleteEquipment(id int) error {
	return cmms.db.DeleteEquipment(id)
}

func (cmms *App) GetAllEquipment() ([]tp.Equipment, error) {
	return cmms.db.GetAllEquipment()
}

func (cmms *App) GetEquipment(id int) (tp.Equipment, error) {
	return cmms.db.GetEquipment(id)
}

func (cmms *App) UpdateEquipment(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	return cmms.db.UpdateEquipment(id, title, year, make, modelNumber, description, categoryId)
}

func (cmms *AppTest) ResetSequenceEquipment(id int) error {
	return cmms.db.ResetSequenceEquipment(id)
}
