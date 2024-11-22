package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type WorkOrderStatus interface {
	CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(title string) error
	ListWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
}

func (a *App) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	return a.db.CreateWorkOrderStatus(wos)
}

func (a *App) DeleteWorkOrderStatus(title string) error {
	return a.db.DeleteWorkOrderStatus(title)
}

func (a *App) ListWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	return a.db.ListWorkOrderStatus()
}

func (a *App) GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error) {
	return a.db.GetWorkOrderStatus(title)
}

func (a *App) UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	return a.db.UpdateWorkOrderStatus(title, wos)
}
