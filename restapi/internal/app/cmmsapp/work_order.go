package cmmsapp

import (
	tm "time"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(int, int, tm.Time, *tm.Time) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByAssetId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, tm.Time, *tm.Time) error
}

func (a *App) CreateWorkOrder(preventativeTaskId int, statusId int, startDateTime tm.Time, CompleteDateTime *tm.Time) (int, error) {
	return a.db.CreateWorkOrder(preventativeTaskId, statusId, startDateTime, CompleteDateTime)
}

func (a *App) DeleteWorkOrder(id int) error {
	return a.db.DeleteWorkOrder(id)
}

func (a *App) GetAllWorkOrder() ([]tp.WorkOrder, error) {
	return a.db.GetAllWorkOrder()
}

func (a *App) GetAllWorkOrderByAssetId(id int) ([]tp.WorkOrder, error) {
	return a.db.GetAllWorkOrderByAssetId(id)
}

func (a *App) GetWorkOrder(id int) (tp.WorkOrder, error) {
	return a.db.GetWorkOrder(id)
}

func (a *App) UpdateWorkOrder(workOrderId int, statusId int, startDateTime tm.Time, CompleteDateTime *tm.Time) error {
	return a.db.UpdateWorkOrder(workOrderId, statusId, startDateTime, CompleteDateTime)
}
