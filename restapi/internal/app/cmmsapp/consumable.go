package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateConsumableWithTask(assetId string, taskId string, consumableId string, cq apitp.ConsumableQuantityRequest) (apitp.ConsumableQuantityResponse, error) {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	// TODO check that cq doesnt conflict with path params
	stConsResponse, err := a.db.AssociateConsumableWithTask(task.Id, cUid, cq.Quantity)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, errors.Wrapf(err, "AssociateConsumableWithTask failed")
	}

	return convertStoreConsumableQuantityToApiConsumableQuantityResponse(stConsResponse)
}

func (a *App) AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string, cq apitp.ConsumableQuantityRequest) (apitp.ConsumableQuantityResponse, error) {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	stConsResponse, err := a.db.AssociateConsumableWithWorkOrder(workOrder.Id, cUid, cq.Quantity)
	if err != nil {
		return apitp.ConsumableQuantityResponse{}, errors.Wrapf(err, "AssociateConsumableWithWorkOrder failed")
	}

	return convertStoreConsumableQuantityToApiConsumableQuantityResponse(stConsResponse)
}

func (a *App) CreateConsumable(consumable apitp.ConsumableRequest) (apitp.ConsumableResponse, error) {
	if consumable.Id != uuid.Nil {
		return apitp.ConsumableResponse{}, ae.New(ae.CodeInvalid, "consumable id must be nil on create, we will create an id for you")
	}
	consumable.Id = uuid.New()

	err := a.validateConsumable(consumable)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "CreateConsumable validation failed")
	}

	stConsRequest, err := convertApiConsumableRequestToStoreConsumable(consumable)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "CreateConsumable ConvertApiConsumableRequestToStoreConsumable failed")
	}

	stConsResponse, err := a.db.CreateConsumable(stConsRequest)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "CreateConsumable CreateConsumable failed")
	}

	return convertStoreConsumableToApiConsumableResponse(stConsResponse)
}

func (a *App) DeleteConsumable(consumableId string) error {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	// TODO: block deletion if the consumable is in use

	return a.db.DeleteConsumable(consumableUuid)
}

func (a *App) DisassociateConsumableWithTask(assetId string, taskId string, consumableId string) error {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	return a.db.DisassociateConsumableWithTask(task.Id, cUid)
}

func (a *App) DisassociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) error {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	return a.db.DisassociateConsumableWithWorkOrder(workOrder.Id, cUid)
}

func (a *App) GetConsumable(consumableId string) (apitp.ConsumableResponse, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
		return apitp.ConsumableResponse{}, ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	stConsResponse, err := a.db.GetConsumable(consumableUuid)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "GetConsumable failed")
	}

	return convertStoreConsumableToApiConsumableResponse(stConsResponse)
}

func (a *App) ListConsumables() ([]apitp.ConsumableResponse, error) {
	stConsResponses, err := a.db.ListConsumables()
	if err != nil {
		return nil, errors.Wrapf(err, "ListConsumables failed")
	}

	return convertStoreConsumableListToApiConsumableResponseList(stConsResponses)
}

func (a *App) UpdateConsumable(consumableId string, consumable apitp.ConsumableRequest) (apitp.ConsumableResponse, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
		return apitp.ConsumableResponse{}, ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	if consumable.Id != uuid.Nil && consumable.Id != consumableUuid {
		return apitp.ConsumableResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("consumable id mismatch between [%s] and [%s]", consumableId, consumable.Id.String()))
	}

	consumable.Id = consumableUuid
	err = a.validateConsumable(consumable)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "UpdateConsumable validation failed")
	}

	stConsRequest, err := convertApiConsumableRequestToStoreConsumable(consumable)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "UpdateConsumable ConvertApiConsumableRequestToStoreConsumable failed")
	}

	stConsResponse, err := a.db.UpdateConsumable(stConsRequest)
	if err != nil {
		return apitp.ConsumableResponse{}, errors.Wrapf(err, "UpdateConsumable UpdateConsumable failed")
	}

	return convertStoreConsumableToApiConsumableResponse(stConsResponse)
}

func (a *App) validateConsumable(consumable apitp.ConsumableRequest) error {
	if consumable.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "consumable id is required")
	}

	if len(consumable.Title) < apitp.MinEntityTitleLength || len(consumable.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("consumable title length must be between [%d] and [%d] characters",
				apitp.MinEntityTitleLength,
				apitp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) consumableExists(consumableId string) (uuid.UUID, bool, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil || consumableUuid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	_, err = a.db.GetConsumable(consumableUuid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return consumableUuid, false, nil
		}
		return consumableUuid, false, err
	}

	return consumableUuid, true, nil
}

func convertApiConsumableRequestToStoreConsumable(consumableRequest apitp.ConsumableRequest) (storetp.Consumable, error) {
	return storetp.Consumable{}, ae.New(ae.CodeNotImplemented, "ConvertApiConsumableRequestToStoreConsumable not implemented")
}

func convertStoreConsumableListToApiConsumableResponseList(storeConsumables []storetp.Consumable) ([]apitp.ConsumableResponse, error) {
	return []apitp.ConsumableResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableListToApiConsumableResponseList not implemented")
}

func convertStoreConsumableQuantityListToApiConsumableQuantityResponseList(storeConsumables []storetp.ConsumableQuantity) ([]apitp.ConsumableQuantityResponse, error) {
	return []apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableQuantityListToApiConsumableQuantityResponseList not implemented")
}

func convertStoreConsumableQuantityToApiConsumableQuantityResponse(storeConsumable storetp.ConsumableQuantity) (apitp.ConsumableQuantityResponse, error) {
	return apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableQuantityToApiConsumableQuantityResponse not implemented")
}

func convertStoreConsumableToApiConsumableResponse(storeConsumable storetp.Consumable) (apitp.ConsumableResponse, error) {
	return apitp.ConsumableResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableToApiConsumableResponse not implemented")
}
