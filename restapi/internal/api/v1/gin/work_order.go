package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// start with long routes then evaluate short routes later
var baseLongWorkOrderRoute = fmt.Sprintf("%s/work-orders", indAssetTaskRoute)
var workOrderId = "workOrderId"
var indLongWorkOrderRoute = fmt.Sprintf("%s/:workOrderId", baseLongWorkOrderRoute)

func (h *Api) registerWorkOrderRoutes() {
	h.router.POST(baseLongWorkOrderRoute, h.createWorkOrder)
	h.router.DELETE(indLongWorkOrderRoute, h.deleteAssetTaskWorkOrder)
	h.router.GET(baseLongWorkOrderRoute, h.listAssetTaskWorkOrders)
	h.router.GET(indLongWorkOrderRoute, h.getAssetTaskWorkOrder)
	h.router.PUT(indLongWorkOrderRoute, h.updateAssetTaskWorkOrder)
}

// CreateWorkOrder godoc
//
//	@Summary		Create a work order
//	@Description	Create a work order
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			workOrder	body	tp.WorkOrder	true	"Work Order object"
//	@Produce		json
//	@Success		201	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId}/work-orders [post]
func (h *Api) createWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.CreateWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), workOrder)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, workOrder))
}

// DeleteAssetTaskWorkOrder godoc
//
//	@Summary		Delete an asset task work order
//	@Description	Delete an asset task work order
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId}/work-orders/{workOrderId} [delete]
func (h *Api) deleteAssetTaskWorkOrder(c *gin.Context) {
	err := h.app.DeleteAssetTaskWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListAssetTaskWorkOrders godoc
//
//	@Summary		List asset task work orders
//	@Description	List all asset task work orders
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId}/work-orders [get]
func (h *Api) listAssetTaskWorkOrders(c *gin.Context) {
	workOrders, err := h.app.ListAssetTaskWorkOrders(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrders))
}

// GetAssetTaskWorkOrder godoc
//
//	@Summary		Get an asset task work order
//	@Description	Get an asset task work order
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId}/work-orders/{workOrderId} [get]
func (h *Api) getAssetTaskWorkOrder(c *gin.Context) {
	workOrder, err := h.app.GetAssetTaskWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(workOrderId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}

// UpdateAssetTaskWorkOrder godoc
//
//	@Summary		Update an asset task work order
//	@Description	Update an asset task work order
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Param			workOrder	body	tp.WorkOrder	true	"Work Order object"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrder
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId}/work-orders/{workOrderId} [put]
func (h *Api) updateAssetTaskWorkOrder(c *gin.Context) {
	var workOrder tp.WorkOrder
	if err := c.BindJSON(&workOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workOrder, err := h.app.UpdateAssetTaskWorkOrder(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(workOrderId), workOrder)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, workOrder))
}
