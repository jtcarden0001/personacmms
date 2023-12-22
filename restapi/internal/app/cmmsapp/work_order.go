package cmmsapp

import (
	tm "time"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(int, int, int, tm.Time, *tm.Time) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByEquipmentId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, int, int, tm.Time, *tm.Time) error
}

func (a *App) CreateWorkOrder(taskId int, statusId int, startDateTime tm.Time, CompleteDateTime *tm.Time) (int, error) {
	return a.db.CreateWorkOrder(taskId, statusId, startDateTime, CompleteDateTime)
}

func (a *App) DeleteWorkOrder(id int) error {
	return a.db.DeleteWorkOrder(id)
}

func (a *App) GetAllWorkOrder() ([]tp.WorkOrder, error) {
	return a.db.GetAllWorkOrder()
}

func (a *App) GetAllWorkOrderByEquipmentId(id int) ([]tp.WorkOrder, error) {
	return a.db.GetAllWorkOrderByEquipmentId(id)
}

func (a *App) GetWorkOrder(id int) (tp.WorkOrder, error) {
	return a.db.GetWorkOrder(id)
}

func (a *App) UpdateWorkOrder(id int, taskId int, statusId int, startDateTime tm.Time, CompleteDateTime *tm.Time) error {
	return a.db.UpdateWorkOrder(id, taskId, statusId, startDateTime, CompleteDateTime)
}
