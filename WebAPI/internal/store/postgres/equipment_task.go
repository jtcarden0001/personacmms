package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type EquipmentTask interface {
	CreateEquipmentTask(int, int) error
	DeleteEquipmentTask(int, int) error
	GetAllEquipmentTask() ([]tp.EquipmentTask, error)
	GetAllEquipmentTaskByEquipmentId(int) ([]tp.EquipmentTask, error)
	GetAllEquipmentTaskByTaskId(int) ([]tp.EquipmentTask, error)
	// GetEquipmentTask(int, int) (tp.EquipmentTask, error)
	// UpdateEquipmentTask(int, int) error
}

func (pg *Store) CreateEquipmentTask(equipmentId int, taskId int) error {
	return errors.New("not implemented")
}

func (pg *Store) DeleteEquipmentTask(equipmentId int, taskId int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllEquipmentTask() ([]tp.EquipmentTask, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllEquipmentTaskByEquipmentId(equipmentId int) ([]tp.EquipmentTask, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllEquipmentTaskByTaskId(taskId int) ([]tp.EquipmentTask, error) {
	return nil, errors.New("not implemented")
}
