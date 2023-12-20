package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type WorkOrderStatus interface {
	CreateWorkOrderStatus(string) (int, error)
	DeleteWorkOrderStatus(int) error
	GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(int) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(int, string) error
}

func (pg *Store) CreateWorkOrderStatus(name string) (int, error) {
	return 0, errors.New("not implemented")
}

func (pg *Store) DeleteWorkOrderStatus(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetWorkOrderStatus(id int) (tp.WorkOrderStatus, error) {
	return tp.WorkOrderStatus{}, errors.New("not implemented")
}

func (pg *Store) UpdateWorkOrderStatus(id int, name string) error {
	return errors.New("not implemented")
}
