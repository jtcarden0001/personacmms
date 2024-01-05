package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type WorkOrderStatus interface {
	CreateWorkOrderStatus(string) (int, error)
	DeleteWorkOrderStatus(int) error
	GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(int) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(int, string) error
}

func (a *App) CreateWorkOrderStatus(title string) (int, error) {
	return a.db.CreateWorkOrderStatus(title)
}

func (a *App) DeleteWorkOrderStatus(id int) error {
	return a.db.DeleteWorkOrderStatus(id)
}

func (a *App) GetAllWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	return a.db.GetAllWorkOrderStatus()
}

func (a *App) GetWorkOrderStatus(id int) (tp.WorkOrderStatus, error) {
	return a.db.GetWorkOrderStatus(id)
}

func (a *App) UpdateWorkOrderStatus(id int, title string) error {
	return a.db.UpdateWorkOrderStatus(id, title)
}
