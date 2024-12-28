package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateConsumableWithTask(assetId string, taskId string, consumableId string) (tp.Consumable, error) {
	// TODO: implement supporting functionality in the store layer
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "AssociateConsumableWithTask not implemented")
}

func (a *App) AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) (tp.Consumable, error) {
	// TODO: implement supporting functionality in the store layer
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "AssociateConsumableWithWorkOrder not implemented")
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

	return a.db.DeleteConsumable(consumableUuid)
}

func (a *App) DisassociateConsumableWithTask(assetId string, taskId string, consumableId string) error {
	// TODO: implement supporting functionality in the store layer
	return ae.New(ae.CodeNotImplemented, "DisassociateConsumableWithTask not implemented")
}

func (a *App) DisassociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) error {
	// TODO: implement supporting functionality in the store layer
	return ae.New(ae.CodeNotImplemented, "DisassociateConsumableWithWorkOrder not implemented")
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

func (a *App) consumableExists(consumableId uuid.UUID) (bool, error) {
	_, err := a.db.GetConsumable(consumableId)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
