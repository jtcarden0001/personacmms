package cmmsapp

import (
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// Create a work order status
func (a *App) CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	return a.db.CreateWorkOrderStatus(wos)
}

// Delete a work order status
func (a *App) DeleteWorkOrderStatus(title string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetWorkOrderStatus(title); err != nil {
		return err
	}

	return a.db.DeleteWorkOrderStatus(title)
}

// List all work order statuses
func (a *App) ListWorkOrderStatus() ([]tp.WorkOrderStatus, error) {
	return a.db.ListWorkOrderStatuses()
}

// Get a work order status
func (a *App) GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error) {
	return a.db.GetWorkOrderStatus(title)
}

// Update a work order status
func (a *App) UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error) {
	if wos.Title == "" {
		return tp.WorkOrderStatus{}, ae.ErrWorkOrderStatusTitleRequired
	}

	return a.db.UpdateWorkOrderStatus(title, wos)
}
