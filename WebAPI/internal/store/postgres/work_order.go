package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(int, int, int, string, *string) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByEquipmentId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, int, int, string, *string) error
}

type WorkOrderTest interface {
	ResetSequenceWorkOrder(int) error
}

func (pg *Store) CreateWorkOrder(equipmentId int, taskId int, statusId int, startDateTime string, CompleteDateTime *string) (int, error) {
	return 0, errors.New("not implemented")
}

func (pg *Store) DeleteWorkOrder(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllWorkOrder() ([]tp.WorkOrder, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllWorkOrderByEquipmentId(equipmentId int) ([]tp.WorkOrder, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetWorkOrder(id int) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, errors.New("not implemented")
}

func (pg *Store) UpdateWorkOrder(id int, equipmentId int, taskId int, statusId int, startDateTime string, CompleteDateTime *string) error {
	return errors.New("not implemented")
}
