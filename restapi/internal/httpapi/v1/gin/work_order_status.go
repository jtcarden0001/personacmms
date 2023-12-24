package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerWorkOrderStatusRoutes() {
	baseRoute := fmt.Sprintf("%s/work-order-statuses", routePrefix)
	individualRoute := fmt.Sprintf("%s/:workOrderStatusId", baseRoute)

	h.router.POST(baseRoute, h.createWorkOrderStatus)
	h.router.DELETE(individualRoute, h.deleteWorkOrderStatus)
	h.router.GET(baseRoute, h.getAllWorkOrderStatus)
	h.router.GET(individualRoute, h.getWorkOrderStatus)
	h.router.PUT(individualRoute, h.updateWorkOrderStatus)
}

func (h *HttpApi) createWorkOrderStatus(c *gin.Context) {
	var wos tp.WorkOrderStatus
	if err := c.BindJSON(&wos); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateWorkOrderStatus(wos.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		wos.Id = id
		c.IndentedJSON(201, wos) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteWorkOrderStatus(c *gin.Context) {
	wosId, err := strconv.Atoi(c.Param("workOrderStatusId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteWorkOrderStatus(wosId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllWorkOrderStatus(c *gin.Context) {
	woss, err := h.app.GetAllWorkOrderStatus()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, woss) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getWorkOrderStatus(c *gin.Context) {
	wosId, err := strconv.Atoi(c.Param("workOrderStatusId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wos, err := h.app.GetWorkOrderStatus(wosId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, wos) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateWorkOrderStatus(c *gin.Context) {
	wosId, err := strconv.Atoi(c.Param("workOrderStatusId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var wos tp.WorkOrderStatus
	if err := c.BindJSON(&wos); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wos.Id = wosId
	err = h.app.UpdateWorkOrderStatus(wos.Id, wos.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
