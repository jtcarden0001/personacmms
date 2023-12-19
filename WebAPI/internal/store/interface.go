package store

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

// TODO: probably will need to devise a way to break these interfaces up soon.
type Store interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategories() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error

	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
	UpdateEquipmentCategoryFK(int, int) error

	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

type StoreTest interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategories() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error
	ResetSequenceEquipmentCategory(int) error

	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
	UpdateEquipmentCategoryFK(int, int) error
	ResetSequenceEquipment(int) error

	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
	ResetSequenceTool(int) error
}
