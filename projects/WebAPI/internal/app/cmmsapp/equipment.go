package cmmsapp

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

func (cmms *App) CreateEquipment(title string, description string) (int, error) {
	return cmms.store.CreateEquipment(title, description)
}

func (cmms *App) DeleteEquipment(id int) error {
	return cmms.store.DeleteEquipment(id)
}

func (cmms *App) GetAllEquipment() ([]tp.Equipment, error) {
	return cmms.store.GetAllEquipment()
}

func (cmms *App) GetEquipment(id int) (tp.Equipment, error) {
	return cmms.store.GetEquipment(id)
}

func (cmms *App) UpdateEquipment(id int, title string, description string) error {
	return cmms.store.UpdateEquipment(id, title, description)
}

func (cmms *AppTest) ResetSequenceEquipment(id int) error {
	return cmms.store.ResetSequenceEquipment(id)
}
