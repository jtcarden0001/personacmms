package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type EquipmentCategory interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategory() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error
}
