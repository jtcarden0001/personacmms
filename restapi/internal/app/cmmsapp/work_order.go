package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateWorkOrder(groupTitle string, assetTitle string, assetTaskId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	//TODO: validate and populate workorder

	return a.db.CreateWorkOrder(wo)
}

func (a *App) DeleteAssetTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string) error {
	// TODO: validate work order

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return err
	}

	return a.db.DeleteWorkOrder(woIdParsed)
}

func (a *App) ListAssetTaskWorkOrders(groupTitle string, assetTitle string, atId string) ([]tp.WorkOrder, error) {
	// TODO: validate
	atIdParsed, err := uuid.Parse(atId)
	if err != nil {
		return []tp.WorkOrder{}, err
	}

	allWorkOrders, err := a.db.ListWorkOrders()
	if err != nil {
		return []tp.WorkOrder{}, err
	}

	// filter work orders by asset task id
	var assetTaskWorkOrders []tp.WorkOrder
	for _, wo := range allWorkOrders {
		if wo.AssetTaskId == atIdParsed {
			assetTaskWorkOrders = append(assetTaskWorkOrders, wo)
		}
	}

	return assetTaskWorkOrders, nil
}

func (a *App) GetAssetTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string) (tp.WorkOrder, error) {
	// TODO: validate work order

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.GetWorkOrder(woIdParsed)
}

func (a *App) UpdateAssetTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// TODO: validate and populate workorder

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.UpdateWorkOrder(woIdParsed, wo)
}
