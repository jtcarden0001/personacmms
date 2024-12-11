package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateWorkOrder(groupTitle string, assetTitle string, taskId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	if err := a.validateWorkOrder(wo); err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.CreateWorkOrder(wo)
}

func (a *App) DeleteTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string) error {
	// TODO: validate work order

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return err
	}

	return a.db.DeleteWorkOrder(woIdParsed)
}

func (a *App) ListTaskWorkOrders(groupTitle string, assetTitle string, atId string) ([]tp.WorkOrder, error) {
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
	var taskWorkOrders []tp.WorkOrder
	for _, wo := range allWorkOrders {
		if wo.TaskId == atIdParsed {
			taskWorkOrders = append(taskWorkOrders, wo)
		}
	}

	return taskWorkOrders, nil
}

func (a *App) GetTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string) (tp.WorkOrder, error) {
	// TODO: validate work order

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.GetWorkOrder(woIdParsed)
}

func (a *App) UpdateTaskWorkOrder(groupTitle string, assetTitle string, atId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// TODO: validate and populate workorder

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.UpdateWorkOrder(woIdParsed, wo)
}

func (a *App) validateWorkOrder(wo tp.WorkOrder) error {
	return nil
}
