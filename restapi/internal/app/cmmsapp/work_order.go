package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) AssociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "AssociateWorkOrderWithTask not implemented")
}

func (a *App) CreateWorkOrder(assetId string, taskId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "CreateWorkOrder not implemented")
}

func (a *App) DeleteWorkOrder(assetId string, tId string, woId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteWorkOrder not implemented")
}

func (a *App) DisassociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateWorkOrderWithTask not implemented")
}

func (a *App) GetWorkOrder(assetId string, atId string, woId string) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "GetWorkOrder not implemented")
}

func (a *App) ListWorkOrders(assetId string, atId string) ([]tp.WorkOrder, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListWorkOrders not implemented")
}

func (a *App) UpdateWorkOrder(assetId string, tId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "UpdateWorkOrder not implemented")
}
