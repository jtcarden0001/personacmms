package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) AssociateConsumableWithTask(assetId string, taskId string, consumableId string) (tp.Consumable, error) {
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "AssociateConsumableWithTask not implemented")
}

func (a *App) AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) (tp.Consumable, error) {
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "AssociateConsumableWithWorkOrder not implemented")
}

func (a *App) CreateConsumable(consumable tp.Consumable) (tp.Consumable, error) {
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "CreateConsumable not implemented")
}

func (a *App) DeleteConsumable(consumableId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteConsumable not implemented")
}

func (a *App) DisassociateConsumableWithTask(assetId string, taskId string, consumableId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateConsumableWithTask not implemented")
}

func (a *App) DisassociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateConsumableWithWorkOrder not implemented")
}

func (a *App) GetConsumable(consumableId string) (tp.Consumable, error) {
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "GetConsumable not implemented")
}

func (a *App) ListConsumables() ([]tp.Consumable, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListConsumables not implemented")
}

func (a *App) UpdateConsumable(consumableId string, consumable tp.Consumable) (tp.Consumable, error) {
	return tp.Consumable{}, ae.New(ae.CodeNotImplemented, "UpdateConsumable not implemented")
}
