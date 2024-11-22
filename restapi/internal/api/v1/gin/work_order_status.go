package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerWorkOrderStatusRoutes() {
	baseRoute := fmt.Sprintf("%s/work-order-statuses", routePrefix)
	individualRoute := fmt.Sprintf("%s/:workOrderStatusTitle", baseRoute)

	h.router.POST(baseRoute, h.createWorkOrderStatus)
	h.router.DELETE(individualRoute, h.deleteWorkOrderStatus)
	h.router.GET(baseRoute, h.listWorkOrderStatus)
	h.router.GET(individualRoute, h.getWorkOrderStatus)
	h.router.PUT(individualRoute, h.updateWorkOrderStatus)
}

// CreateWorkOrderStatus godoc
//
//	@Summary		Create a work order status
//	@Description	Create a work order status
//	@Accept			json
//	@Param			workOrderStatus	body	tp.WorkOrderStatus	true	"Work Order Status object"
//	@Produce		json
//	@Success		201	{object}	tp.WorkOrderStatus
//	@Router			/work-order-statuses [post]
func (h *Api) createWorkOrderStatus(c *gin.Context) {
	var wos tp.WorkOrderStatus
	if err := c.BindJSON(&wos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wos, err := h.app.CreateWorkOrderStatus(wos)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, wos))
}

// DeleteWorkOrderStatus godoc
//
//	@Summary		Delete a work order status
//	@Description	Delete a work order status
//	@Param			workOrderStatusTitle	path	string	true	"Work Order Status Title"
//	@Success		204
//	@Failure		404
//	@Router			/work-order-statuses/{workOrderStatusTitle} [delete]
func (h *Api) deleteWorkOrderStatus(c *gin.Context) {
	err := h.app.DeleteWorkOrderStatus(c.Param("workOrderStatusTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListWorkOrderStatus godoc
//
//	@Summary		List work order statuses
//	@Description	List all work order statuses
//	@Produce		json
//	@Success		200	{object}	[]tp.WorkOrderStatus
//	@Router			/work-order-statuses [get]
func (h *Api) listWorkOrderStatus(c *gin.Context) {
	wos, err := h.app.ListWorkOrderStatus()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, wos))
}

// GetWorkOrderStatus godoc
//
//	@Summary		Get a work order status
//	@Description	Get a work order status
//	@Param			workOrderStatusTitle	path	string	true	"Work Order Status Title"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrderStatus
//	@Router			/work-order-statuses/{workOrderStatusTitle} [get]
func (h *Api) getWorkOrderStatus(c *gin.Context) {
	wos, err := h.app.GetWorkOrderStatus(c.Param("workOrderStatusTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, wos))
}

// UpdateWorkOrderStatus godoc
//
//	@Summary		Update a work order status
//	@Description	Update a work order status
//	@Param			workOrderStatusTitle	path	string	true	"Work Order Status Title"
//	@Accept			json
//	@Param			workOrderStatus	body	tp.WorkOrderStatus	true	"Work Order Status object"
//	@Produce		json
//	@Success		200	{object}	tp.WorkOrderStatus
//	@Router			/work-order-statuses/{workOrderStatusTitle} [put]
func (h *Api) updateWorkOrderStatus(c *gin.Context) {
	var wos tp.WorkOrderStatus
	if err := c.BindJSON(&wos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wos, err := h.app.UpdateWorkOrderStatus(c.Param("workOrderStatusTitle"), wos)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, wos))
}
