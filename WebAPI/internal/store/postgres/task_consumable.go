package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type TaskConsumable interface {
	CreateTaskConsumable(int, int, string) error
	DeleteTaskConsumable(int, int) error
	GetAllTaskConsumable() ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByEquipmentId(int) ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByTaskId(int) ([]tp.TaskConsumable, error)
	// GetTaskConsumable(int, int) (tp.TaskConsumable, error)
	UpdateTaskConsumable(int, int, string) error
}
