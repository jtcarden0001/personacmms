package store

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

type Store interface {
	CreateEquipment(string, int, string, string, string) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string) error

	CreateTool(string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string) error
}

type StoreTest interface {
	CreateEquipment(string, int, string, string, string) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string) error
	ResetSequenceEquipment(int) error

	CreateTool(string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string) error
	ResetSequenceTool(int) error
}
