package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type WorkOrder interface {
	CreateWorkOrder(string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
	DeleteAssetTaskWorkOrder(string, string, string, string) error
	ListAssetTaskWorkOrders(string, string, string) ([]tp.WorkOrder, error)
	GetAssetTaskWorkOrder(string, string, string, string) (tp.WorkOrder, error)
	UpdateAssetTaskWorkOrder(string, string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
}

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

	return a.db.ListAssetTaskWorkOrders(atIdParsed)
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
