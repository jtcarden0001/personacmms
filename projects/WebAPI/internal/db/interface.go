package db

import tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"

type Store interface {
	CreateEquipment(string, string) (int, error)
	DeleteEquipment(int) error
	GetEquipment(int) (tp.Equipment, error)
	GetAllEquipment() ([]tp.Equipment, error)
	UpdateEquipment(int, string, string) error

	CreateTool(string) (int, error)
	DeleteTool(int) error
	GetTool(int) (tp.Tool, error)
	GetAllTool() ([]tp.Tool, error)
	UpdateTool(int, string) error
}
