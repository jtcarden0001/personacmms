package cmmsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) (apitp.WorkOrderResponse, error) {
	wUid, wFound, err := a.workOrderExists(workOrderId)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "error checking work order exists")
	}

	if !wFound {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("work order with id [%s] not found", workOrderId))
	}

	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.WorkOrderResponse{}, err
	}

	stWOResponse, err := a.db.AssociateWorkOrderWithTask(task.Id, wUid)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "AssociateWorkOrderWithTask failed")
	}

	return convertStoreWorkOrderToApiWorkOrderResponse(stWOResponse)
}

func (a *App) CreateWorkOrder(assetId string, wo apitp.WorkOrderRequest) (apitp.WorkOrderResponse, error) {
	if wo.Id != uuid.Nil {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeInvalid, "work order id must be nil on create, we will create an id for you")
	}
	wo.Id = uuid.New()

	aUid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	if wo.AssetId != uuid.Nil && wo.AssetId != aUid {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset id mismatch [%s] does not match [%s]", wo.AssetId, assetId))
	}

	wo.AssetId = aUid
	err = a.validateWorkOrder(wo)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "CreateWorkOrder validation failed")
	}

	stWORequest, err := convertApiWorkOrderRequestToStoreWorkOrder(wo)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "CreateWorkOrder - error converting to store type")
	}

	stWOResponse, err := a.db.CreateWorkOrder(stWORequest)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "CreateWorkOrder - error creating work order")
	}

	return convertStoreWorkOrderToApiWorkOrderResponse(stWOResponse)
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

func (a *App) GetWorkOrder(assetId string, woId string) (apitp.WorkOrderResponse, error) {
	woUid, err := uuid.Parse(woId)
	if err != nil {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeInvalid, "work order id must be a valid uuid")
	}

	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	wo, err := a.db.GetWorkOrder(woUid)
	if err != nil {
		return apitp.WorkOrderResponse{}, err
	}

	if wo.AssetId != aUid {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("work order with id [%s] not found on asset with id [%s]", woId, assetId))
	}

	return convertStoreWorkOrderToApiWorkOrderResponse(wo)
}

func (a *App) ListWorkOrdersByAsset(assetId string) ([]apitp.WorkOrderResponse, error) {
	aUid, aFound, err := a.assetExists(assetId)
	if err != nil {
		return nil, errors.Wrapf(err, "error checking asset exists")
	}

	if !aFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("asset with id [%s] not found", assetId))
	}

	stWOResponses, err := a.db.ListWorkOrdersByAsset(aUid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListWorkOrdersByAsset failed")
	}

	return convertStoreWorkOrderListToApiWorkOrderResponseList(stWOResponses)
}

func (a *App) ListWorkOrderStatus() ([]string, error) {
	keys := make([]string, 0, len(apitp.ValidWorkOrderStatuses))
	for k := range apitp.ValidWorkOrderStatuses {
		keys = append(keys, k)
	}

	return keys, nil
}

func (a *App) UpdateWorkOrder(assetId string, woId string, wo apitp.WorkOrderRequest) (apitp.WorkOrderResponse, error) {
	// check asset and work order existence and coherency
	gwo, err := a.GetWorkOrder(assetId, woId)
	if err != nil {
		return apitp.WorkOrderResponse{}, err
	}

	if wo.Id != uuid.Nil && wo.Id != gwo.Id {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("work order id mismatch [%s] and [%s]", wo.Id, gwo.Id))
	}
	wo.Id = gwo.Id

	if wo.AssetId != uuid.Nil && wo.AssetId != gwo.AssetId {
		return apitp.WorkOrderResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("asset id mismatch [%s] and [%s]", wo.AssetId, gwo.AssetId))
	}
	wo.AssetId = gwo.AssetId

	err = a.validateWorkOrder(wo)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "UpdateWorkOrder validation failed")
	}

	stWORequest, err := convertApiWorkOrderRequestToStoreWorkOrder(wo)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "UpdateWorkOrder conversion failed")
	}

	stWOResponse, err := a.db.UpdateWorkOrder(stWORequest)
	if err != nil {
		return apitp.WorkOrderResponse{}, errors.Wrapf(err, "UpdateWorkOrder failed")
	}

	return convertStoreWorkOrderToApiWorkOrderResponse(stWOResponse)
}

func (a *App) validateWorkOrder(wo apitp.WorkOrderRequest) error {
	if wo.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "work order id is required")
	}

	if len(wo.Title) < apitp.MinEntityTitleLength || len(wo.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("work order title must be between [%d] and [%d] characters",
				apitp.MinEntityTitleLength,
				apitp.MaxEntityTitleLength))
	}

	if wo.CreatedDate.After(time.Now()) {
		return ae.New(ae.CodeInvalid, "work order created date cannot be in the future")
	}

	if !apitp.ValidWorkOrderStatuses[wo.Status] {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("work order status must be one of [%s]", apitp.PrintValidWorkOrderStatuses()))
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

func convertApiWorkOrderRequestToStoreWorkOrder(wo apitp.WorkOrderRequest) (storetp.WorkOrder, error) {
	return storetp.WorkOrder{}, ae.New(ae.CodeNotImplemented, "convertApiWorkOrderRequestToStoreWorkOrder not implemented")
}

func convertStoreWorkOrderListToApiWorkOrderResponseList(wos []storetp.WorkOrder) ([]apitp.WorkOrderResponse, error) {
	return []apitp.WorkOrderResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreWorkOrderListToApiWorkOrderResponseList not implemented")
}

func convertStoreWorkOrderToApiWorkOrderResponse(wo storetp.WorkOrder) (apitp.WorkOrderResponse, error) {
	return apitp.WorkOrderResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreWorkOrderToApiWorkOrderResponse not implemented")
}
