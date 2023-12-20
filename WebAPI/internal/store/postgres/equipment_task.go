package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

// these interfaces are platform agnostic but I like having them in the same file as the implementation,
// will move to a common location if I field the idea of another implemntation
type EquipmentTask interface {
	CreateEquipmentTask(int, int) error
	DeleteEquipmentTask(int, int) error
	GetAllEquipmentTask() ([]tp.EquipmentTask, error)
	GetAllEquipmentTaskByEquipmentId(int) (tp.EquipmentTask, error)
	GetAllEquipmentTaskByTaskId(int) (tp.EquipmentTask, error)
	// GetEquipmentTask(int, int) (tp.EquipmentTask, error)
	// UpdateEquipmentTask(int, int) error
}
