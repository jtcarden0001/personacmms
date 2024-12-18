package cmmsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// create a work order
func (a *App) CreateWorkOrder(groupTitle string, assetTitle string, taskId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// validate task namespace
	t, err := a.GetTask(groupTitle, assetTitle, taskId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	if err := a.validateAndInterpolateWorkOrder(t.Id, &wo); err != nil {
		return tp.WorkOrder{}, err
	}

	return a.db.CreateWorkOrder(wo)
}

// delete a work order
func (a *App) DeleteWorkOrder(groupTitle string, assetTitle string, tId string, woId string) error {
	// validate the namespace coherency of the task and retrieve the work order
	wo, err := a.GetWorkOrder(groupTitle, assetTitle, tId, woId)
	if err != nil {
		return err
	}

	return a.db.DeleteWorkOrder(wo.Id)
}

// get a work orderfor a task TODO: update name to be more specific
func (a *App) GetWorkOrder(groupTitle string, assetTitle string, atId string, woId string) (tp.WorkOrder, error) {
	// validate task namespace coherency
	t, err := a.GetTask(groupTitle, assetTitle, atId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	woIdParsed, err := uuid.Parse(woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	wo, err := a.db.GetWorkOrderForTask(t.Id, woIdParsed)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	return wo, nil
}

// list work orders for a task TODO: update name to be more specific
func (a *App) ListWorkOrders(groupTitle string, assetTitle string, atId string) ([]tp.WorkOrder, error) {
	// validate task namespace coherency
	t, err := a.GetTask(groupTitle, assetTitle, atId)
	if err != nil {
		return []tp.WorkOrder{}, err
	}

	wos, err := a.db.ListWorkOrdersByTaskId(t.Id)
	if err != nil {
		return []tp.WorkOrder{}, err
	}

	return wos, nil
}

// update a work order
func (a *App) UpdateWorkOrder(groupTitle string, assetTitle string, tId string, woId string, wo tp.WorkOrder) (tp.WorkOrder, error) {
	// validate the namespace coherency of the task and retrieve the work order
	oldWo, err := a.GetWorkOrder(groupTitle, assetTitle, tId, woId)
	if err != nil {
		return tp.WorkOrder{}, err
	}

	// validate created date
	if wo.CreatedDate.IsZero() {
		wo.CreatedDate = time.Now()
	} else {
		if wo.CreatedDate.After(time.Now()) {
			return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "work order created date is in the future")
		}

		if wo.CompletedDate != nil && wo.CompletedDate.Before(wo.CreatedDate) {
			return tp.WorkOrder{}, ae.New(ae.CodeInvalid, "work order completed date is before created date")
		}
	}

	// validate status title
	if wo.StatusTitle == "" {
		wo.StatusTitle = tp.WorkOrderStatusNew
	} else {
		if !tp.IsValidWorkOrderStatus(wo.StatusTitle) {
			return tp.WorkOrder{}, ae.New(
				ae.CodeInvalid,
				fmt.Sprintf("work order status title is invalid, %s", ae.CreateInvalidWorkOrderStatusMessage()),
			)
		}
	}

	return a.db.UpdateWorkOrder(oldWo.Id, wo)
}

// TODO: figure out this common validation code
func (a *App) validateAndInterpolateWorkOrder(taskId uuid.UUID, wo *tp.WorkOrder) error {
	// validate task id
	if wo.TaskId != uuid.Nil && wo.TaskId != taskId {
		return ae.New(ae.CodeInvalid, "work order task id field is invalid")
	}
	wo.TaskId = taskId

	// validate created date
	if wo.CreatedDate.IsZero() {
		wo.CreatedDate = time.Now()
	} else {
		if wo.CreatedDate.After(time.Now()) {
			return ae.New(ae.CodeInvalid, "work order created date is in the future")
		}

		if wo.CompletedDate != nil && wo.CompletedDate.Before(wo.CreatedDate) {
			return ae.New(ae.CodeInvalid, "work order completed date is before created date")
		}
	}

	// validate status title
	if wo.StatusTitle == "" {
		wo.StatusTitle = tp.WorkOrderStatusNew
	} else {
		if !tp.IsValidWorkOrderStatus(wo.StatusTitle) {
			return ae.New(
				ae.CodeInvalid,
				fmt.Sprintf("work order status title is invalid, %s", ae.CreateInvalidWorkOrderStatusMessage()),
			)
		}
	}

	return nil
}
