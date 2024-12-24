package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /assets/{assetId}/work-orders (JSON) done
// - GET  /assets/{assetId}/work-orders/{workOrderId} done
// - GET  /assets/{assetId}/work-orders done
// - PUT  /assets/{assetId}/work-orders/{workOrderId} (JSON) done
// - DEL  /assets/{assetId}/work-orders/{workOrderId} done
//
// - PUT  /assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} done
// - DEL  /assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} done

var workOrderId = "workOrderId"
var workOrderGp = "work-orders"
var workOrderResource = fmt.Sprintf("%s/:%s", workOrderGp, workOrderId)

var baseWorkOrderRoute = fmt.Sprintf("%s/%s", indAssetRoute, workOrderGp)
var indWorkOrderRoute = fmt.Sprintf("%s/%s", indAssetRoute, workOrderResource)

func (h *Api) registerWorkOrderRoutes() {
	h.router.POST(baseWorkOrderRoute, h.createWorkOrder)

	h.router.DELETE(indWorkOrderRoute, h.deleteWorkOrder)
	h.router.DELETE(fmt.Sprintf("%s/%s", indTaskRoute, workOrderResource), h.disassociateWorkOrderWithTask)

	h.router.GET(indWorkOrderRoute, h.getWorkOrder)
	h.router.GET(baseWorkOrderRoute, h.listWorkOrders)

	h.router.PUT(fmt.Sprintf("%s/%s", indTaskRoute, workOrderResource), h.associateWorkOrderWithTask)
	h.router.PUT(indWorkOrderRoute, h.updateWorkOrder)
}

// AssociateWorkOrderWithTask godoc
//
//	@Summary		Associate a work order with a task
//	@Description	Associate a work order with a task
//	@Tags			work-orders
//	@Produce		json
//	@Param			assetId		path		string	true	"Asset ID"
//	@Param			taskId		path		string	true	"Task ID"
//	@Param			workOrderId	path		string	true	"Work Order ID"
//	@Success		200			{object}	tp.WorkOrder
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} [put]
func (h *Api) associateWorkOrderWithTask(c *gin.Context) {
	workOrder, err := h.app.AssociateWorkOrderWithTask(c.Param(assetId), c.Param(taskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}

// CreateWorkOrder godoc
//
//	@Summary		Create a work order
//	@Description	Create a work order
//	@Tags			work-orders
//	@Accept			json
//	@Produce		json
//	@Param			assetId		path		string			true	"Asset Id"
//	@Param			taskId		path		string			true	"Asset Task Id"
//	@Param			workOrder	body		tp.WorkOrder	true	"Work Order object"
//	@Success		201			{object}	tp.WorkOrder
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders [post]
func (h *Api) createWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.CreateWorkOrder(c.Param(assetId), c.Param(taskId), workOrder)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, workOrder))
}

// DeleteTaskWorkOrder godoc
//
//	@Summary		Delete an asset task work order
//	@Description	Delete an asset task work order
//	@Tags			work-orders
//	@Param			assetId		path	string	true	"Asset Id"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} [delete]
func (h *Api) deleteWorkOrder(c *gin.Context) {
	err := h.app.DeleteWorkOrder(c.Param(assetId), c.Param(taskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateWorkOrderWithTask godoc
//
//	@Summary		Disassociate a work order with a task
//	@Description	Disassociate a work order with a task
//	@Tags			work-orders
//	@Param			assetId		path	string	true	"Asset Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/tasks [delete]
func (h *Api) disassociateWorkOrderWithTask(c *gin.Context) {
	err := h.app.DisassociateWorkOrderWithTask(c.Param(assetId), c.Param(workOrderId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTaskWorkOrder godoc
//
//	@Summary		Get an asset task work order
//	@Description	Get an asset task work order
//	@Tags			work-orders
//	@Produce		json
//	@Param			assetId		path		string	true	"Asset Id"
//	@Param			taskId		path		string	true	"Asset Task Id"
//	@Param			workOrderId	path		string	true	"Work Order Id"
//	@Success		200			{object}	tp.WorkOrder
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} [get]
func (h *Api) getWorkOrder(c *gin.Context) {
	workOrder, err := h.app.GetWorkOrder(c.Param(assetId), c.Param(taskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}

// ListTaskWorkOrders godoc
//
//	@Summary		List asset task work orders
//	@Description	List all asset task work orders
//	@Tags			work-orders
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			taskId	path		string	true	"Asset Task Id"
//	@Success		200		{object}	[]tp.WorkOrder
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders [get]
func (h *Api) listWorkOrders(c *gin.Context) {
	workOrders, err := h.app.ListWorkOrders(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrders))
}

// UpdateTaskWorkOrder godoc
//
//	@Summary		Update an asset task work order
//	@Description	Update an asset task work order
//	@Tags			work-orders
//	@Accept			json
//	@Produce		json
//	@Param			assetId		path		string			true	"Asset Id"
//	@Param			taskId		path		string			true	"Asset Task Id"
//	@Param			workOrderId	path		string			true	"Work Order Id"
//	@Param			workOrder	body		tp.WorkOrder	true	"Work Order object"
//	@Success		200			{object}	tp.WorkOrder
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/work-orders/{workOrderId} [put]
func (h *Api) updateWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.UpdateWorkOrder(c.Param(assetId), c.Param(taskId), c.Param(workOrderId), workOrder)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}
