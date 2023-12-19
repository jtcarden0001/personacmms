package app

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

type App interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error

	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

type AppTest interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
	ResetSequenceEquipment(int) error

	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
	ResetSequenceTool(int) error
}
