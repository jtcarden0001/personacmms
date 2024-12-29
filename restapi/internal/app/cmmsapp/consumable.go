package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateConsumableWithTask(assetId string, taskId string, consumableId string, cq tp.ConsumableQuantity) (tp.ConsumableQuantity, error) {
	// check asset and task exists and task is associated with asset
	task, err := a.GetTask(assetId, taskId)
	if err != nil {
		return tp.ConsumableQuantity{}, err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return tp.ConsumableQuantity{}, errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return tp.ConsumableQuantity{}, ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	// TODO check that cq doesnt conflict with path params

	return a.db.AssociateConsumableWithTask(task.Id, cUid, cq.Quantity)

}

func (a *App) AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string, cq tp.ConsumableQuantity) (tp.ConsumableQuantity, error) {
	// check asset and work order exists and work order is associated with asset
	workOrder, err := a.GetWorkOrder(assetId, workOrderId)
	if err != nil {
		return tp.ConsumableQuantity{}, err
	}

	cUid, cFound, err := a.consumableExists(consumableId)
	if err != nil {
		return tp.ConsumableQuantity{}, errors.Wrapf(err, "error checking consumable exists")
	}

	if !cFound {
		return tp.ConsumableQuantity{}, ae.New(ae.CodeNotFound, fmt.Sprintf("consumable with id [%s] not found", consumableId))
	}

	return a.db.AssociateConsumableWithWorkOrder(workOrder.Id, cUid)
}

func (a *App) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	if consumable.Id != uuid.Nil {
		return tp.Consumable{}, ae.New(ae.CodeInvalid, "consumable id must be nil on create, we will create an id for you")
	}
	consumable.Id = uuid.New()

	err := a.validateConsumable(consumable)
	if err != nil {
		return tp.Consumable{}, errors.Wrapf(err, "CreateConsumable validation failed")
	}

	return a.db.CreateConsumable(consumable)
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

func (a *App) GetConsumable(consumableId string) (tp.Consumable, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.Consumable{}, ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	return a.db.GetConsumable(consumableUuid)
}

func (a *App) ListConsumables() ([]tp.Consumable, error) {
	return a.db.ListConsumables()
}

func (a *App) UpdateConsumable(consumableId string, consumable tp.Consumable) (tp.Consumable, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
		return tp.Consumable{}, ae.New(ae.CodeInvalid, "consumable id must be a valid uuid")
	}

	if consumable.Id != uuid.Nil && consumable.Id != consumableUuid {
		return tp.Consumable{}, ae.New(ae.CodeInvalid, fmt.Sprintf("consumable id mismatch between [%s] and [%s]", consumableId, consumable.Id.String()))
	}

	consumable.Id = consumableUuid
	err = a.validateConsumable(consumable)
	if err != nil {
		return tp.Consumable{}, errors.Wrapf(err, "UpdateConsumable validation failed")
	}

	return a.db.UpdateConsumable(consumable)
}

func (a *App) validateConsumable(consumable tp.Consumable) error {
	if consumable.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "consumable id is required")
	}

	if len(consumable.Title) < tp.MinEntityTitleLength || len(consumable.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("consumable title length must be between [%d] and [%d] characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) consumableExists(consumableId string) (uuid.UUID, bool, error) {
	consumableUuid, err := uuid.Parse(consumableId)
	if err != nil {
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
