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
	wUid, wFound, err := a.workOrderExists(workOrderId)
	if err != nil {
		return tp.WorkOrder{}, errors.Wrapf(err, "error checking work order exists")
	}

	if !wFound {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, fmt.Sprintf("work order with id [%s] not found", workOrderId))
	}

	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.AssociateWorkOrderWithTask(task.Id, wUid)
}

func (a *App) CreateWorkOrder(assetId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	if wo.Id != uuid.Nil {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "work order id must be nil on create, we will create an id for you")
	}
	wo.Id = uuid.New()

	aUid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	if wo.AssetId != uuid.Nil && wo.AssetId != aUid {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset id mismatch [%s] does not match [%s]", wo.AssetId, assetId))
	}

	wo.AssetId = aUid
	err = a.validateWorkOrder(wo)
	if err != nil {
		return tp.WorkOrder{}, errors.Wrapf(err, "CreateWorkOrder validation failed")
	}

	return a.db.CreateWorkOrder(wo)
}

func (a *App) DeleteWorkOrder(assetId string, woId string) error {
	wUid, err := uuid.Parse(woId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "work order id must be a valid uuid")
	}

	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	// TODO: ensure cascading deletions of associate consumables, tools, and task relationships

	return a.db.DeleteWorkOrderFromAsset(aUid, wUid)

}

func (a *App) DisassociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) error {
	return a.DisassociateTaskWithWorkOrder(assetId, taskId, workOrderId)
}

func (a *App) GetWorkOrder(assetId string, woId string) (tp.WorkOrder, error) {
	woUid, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "work order id must be a valid uuid")
	}

	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return tp.WorkOrder{}, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	wo, err := a.db.GetWorkOrder(woUid)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	if wo.AssetId != aUid {
		return tp.WorkOrder{}, ae.New(ae.CodeNotFound, fmt.Sprintf("work order with id [%s] not found on asset with id [%s]", woId, assetId))
	}

	return wo, nil
}

func (a *App) ListWorkOrdersByAsset(assetId string) ([]tp.WorkOrder, error) {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return nil, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	return a.db.ListWorkOrdersByAsset(aUid)
}

func (a *App) ListWorkOrderStatus() ([]string, error) {
	keys := make([]string, 0, len(tp.ValidWorkOrderStatuses))
	for k := range tp.ValidWorkOrderStatuses {
		keys = append(keys, k)
	}

	return keys, nil
}

func (a *App) UpdateWorkOrder(assetId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// check asset and work order existence and coherency
	gwo, err := a.GetWorkOrder(assetId, woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	if wo.Id != uuid.Nil && wo.Id != gwo.Id {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, fmt.Sprintf("work order id mismatch [%s] and [%s]", wo.Id, gwo.Id))
	}
	wo.Id = gwo.Id

	if wo.AssetId != uuid.Nil && wo.AssetId != gwo.AssetId {
		return tp.WorkOrder{}, ae.New(ae.CodeInvalid, fmt.Sprintf("asset id mismatch [%s] and [%s]", wo.AssetId, gwo.AssetId))
	}
	wo.AssetId = gwo.AssetId

	err = a.validateWorkOrder(wo)
	if err != nil {
		return tp.WorkOrder{}, errors.Wrapf(err, "UpdateWorkOrder validation failed")
	}

	return a.db.UpdateWorkOrder(wo)
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

	_, aFound, err := a.assetExists(wo.AssetId.String())
	if err != nil {
		return errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", wo.AssetId))
	}

	return nil
}

func (a *App) workOrderExists(id string) (uuid.UUID, bool, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "work order id must be a valid uuid")
	}

	_, err = a.db.GetWorkOrder(uid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return uid, false, nil
		}
		return uid, false, err
	}
	return uid, true, nil
}
