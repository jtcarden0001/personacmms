package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// start with long routes then evaluate short routes later
var baseLongWorkOrderRoute = fmt.Sprintf("%s/work-orders", indTaskRoute)
var workOrderId = "workOrderId"
var indLongWorkOrderRoute = fmt.Sprintf("%s/:workOrderId", baseLongWorkOrderRoute)

func (h *Api) registerWorkOrderRoutes() {
	h.router.POST(baseLongWorkOrderRoute, h.createWorkOrder)
	h.router.DELETE(indLongWorkOrderRoute, h.deleteTaskWorkOrder)
	h.router.GET(baseLongWorkOrderRoute, h.listTaskWorkOrders)
	h.router.GET(indLongWorkOrderRoute, h.getTaskWorkOrder)
	h.router.PUT(indLongWorkOrderRoute, h.updateTaskWorkOrder)
}

// CreateWorkOrder godoc
//
//	@Summary		Create a work order
//	@Description	Create a work order
//	@Accept			json
//	@Param			groupTitle	path	string			true	"Group Title"
//	@Param			assetTitle	path	string			true	"Asset Id"
//	@Param			taskId		path	string			true	"Asset Task Id"
//	@Param			workOrder	body	tp.WorkOrder	true	"Work Order object"
//	@Produce		json
//	@Success		201	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/work-orders [post]
func (h *Api) createWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.CreateWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), workOrder)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, workOrder))
}

// DeleteTaskWorkOrder godoc
//
//	@Summary		Delete an asset task work order
//	@Description	Delete an asset task work order
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/work-orders/{workOrderId} [delete]
func (h *Api) deleteTaskWorkOrder(c *gin.Context) {
	err := h.app.DeleteWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTaskWorkOrders godoc
//
//	@Summary		List asset task work orders
//	@Description	List all asset task work orders
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/work-orders [get]
func (h *Api) listTaskWorkOrders(c *gin.Context) {
	workOrders, err := h.app.ListWorkOrders(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrders))
}

// GetTaskWorkOrder godoc
//
//	@Summary		Get an asset task work order
//	@Description	Get an asset task work order
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/work-orders/{workOrderId} [get]
func (h *Api) getTaskWorkOrder(c *gin.Context) {
	workOrder, err := h.app.GetWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}

// UpdateTaskWorkOrder godoc
//
//	@Summary		Update an asset task work order
//	@Description	Update an asset task work order
//	@Accept			json
//	@Param			groupTitle	path	string			true	"Group Title"
//	@Param			assetTitle	path	string			true	"Asset Id"
//	@Param			taskId		path	string			true	"Asset Task Id"
//	@Param			workOrderId	path	string			true	"Work Order Id"
//	@Param			workOrder	body	tp.WorkOrder	true	"Work Order object"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/work-orders/{workOrderId} [put]
func (h *Api) updateTaskWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.UpdateWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(workOrderId), workOrder)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}
