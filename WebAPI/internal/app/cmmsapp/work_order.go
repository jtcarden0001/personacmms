package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type WorkOrder interface {
	CreateWorkOrder(int, int, int, string, *string) (int, error)
	DeleteWorkOrder(int) error
	GetAllWorkOrder() ([]tp.WorkOrder, error)
	GetAllWorkOrderByEquipmentId(int) ([]tp.WorkOrder, error)
	GetWorkOrder(int) (tp.WorkOrder, error)
	UpdateWorkOrder(int, int, int, int, string, *string) error
}

func (a *App) CreateWorkOrder(taskId int, statusId int, startDateTime string, CompleteDateTime *string) (int, error) {
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

func (a *App) UpdateWorkOrder(id int, taskId int, statusId int, startDateTime string, CompleteDateTime *string) error {
	return a.db.UpdateWorkOrder(id, taskId, statusId, startDateTime, CompleteDateTime)
}
