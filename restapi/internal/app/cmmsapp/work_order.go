package cmmsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "AssociateWorkOrderWithTask not implemented")
}

func (a *App) CreateWorkOrder(assetId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	if wo.Id != uuid.Nil {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "work order id must be nil on create, we will create an id for you")
	}
	wo.Id = uuid.New()

	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "CreateWorkOrder not implemented")
}

func (a *App) DeleteWorkOrder(assetId string, woId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteWorkOrder not implemented")
}

func (a *App) DisassociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateWorkOrderWithTask not implemented")
}

func (a *App) GetWorkOrder(assetId string, woId string) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "GetWorkOrder not implemented")
}

func (a *App) ListWorkOrders(assetId string) ([]tp.WorkOrder, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListWorkOrders not implemented")
}

func (a *App) ListWorkOrderStatus() ([]string, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListWorkOrderStatus not implemented")
}

func (a *App) UpdateWorkOrder(assetId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	return tp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "UpdateWorkOrder not implemented")
}

func (a *App) validateWorkOrder(wo tp.WorkOrder) error {
	if wo.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "work order id is required")
	}

	if len(wo.Title) < tp.MinEntityTitleLength || len(wo.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("work order title must be between [%d] and [%d] characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	if wo.CreatedDate.After(time.Now()) {
		return ae.New(ae.CodeInvalid, "work order created date cannot be in the future")
	}

	if !tp.ValidWorkOrderStatuses[wo.Status] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("work order status must be one of [%s]", tp.PrintValidWorkOrderStatuses()))
	}

	return nil
}

func (a *App) workOrderExists(id uuid.UUID) (bool, error) {
	_, err := a.db.GetWorkOrder(id)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
