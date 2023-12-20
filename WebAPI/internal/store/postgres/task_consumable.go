package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type TaskConsumable interface {
	CreateTaskConsumable(int, int, string) error
	DeleteTaskConsumable(int, int) error
	GetAllTaskConsumable() ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByEquipmentId(int) ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByTaskId(int) ([]tp.TaskConsumable, error)
	// GetTaskConsumable(int, int) (tp.TaskConsumable, error)
	UpdateTaskConsumable(int, int, string) error
}

func (pg *Store) CreateTaskConsumable(taskId int, consumableId int, quantity string) error {
	return errors.New("not implemented")
}

func (pg *Store) DeleteTaskConsumable(taskId int, consumableId int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllTaskConsumable() ([]tp.TaskConsumable, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllTaskConsumableByEquipmentId(equipmentId int) ([]tp.TaskConsumable, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllTaskConsumableByTaskId(taskId int) ([]tp.TaskConsumable, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) UpdateTaskConsumable(taskId int, consumableId int, quantity string) error {
	return errors.New("not implemented")
}
