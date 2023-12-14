package cmmsapp

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

func (cmms *App) CreateEquipment(title string, year int, make, modelNumber, description string) (int, error) {
	return cmms.db.CreateEquipment(title, year, make, modelNumber, description)
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

func (cmms *App) UpdateEquipment(id int, title string, year int, make, modelNumber, description string) error {
	return cmms.db.UpdateEquipment(id, title, year, make, modelNumber, description)
}

func (cmms *AppTest) ResetSequenceEquipment(id int) error {
	return cmms.db.ResetSequenceEquipment(id)
}
