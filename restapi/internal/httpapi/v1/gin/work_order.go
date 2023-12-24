package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerWorkOrderRoutes() {
	baseRouteByTask := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/work-orders", routePrefix)
	individualRouteByTask := fmt.Sprintf("%s/:workOrderId", baseRouteByTask)

	h.router.POST(baseRouteByTask, h.createWorkOrderByTask)
	h.router.DELETE(individualRouteByTask, h.deleteWorkOrderByTask)
	h.router.GET(baseRouteByTask, h.getAllWorkOrderByTask)
	h.router.GET(individualRouteByTask, h.getWorkOrderByTask)
	h.router.PUT(individualRouteByTask, h.updateWorkOrderByTask)
}

func (h *HttpApi) createWorkOrderByTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var wo tp.WorkOrder
	if err := c.BindJSON(&wo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo.TaskId = taskId
	id, err := h.app.CreateWorkOrder(wo.TaskId, wo.StatusId, wo.CreatedDate, wo.CompletedDate)
	if err != nil {
		processAppError(c, err)
	} else {
		wo.Id = id
		c.IndentedJSON(201, wo) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	workOrderId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteWorkOrder(workOrderId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woss, err := h.app.GetAllWorkOrder()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, woss) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo, err := h.app.GetWorkOrder(woId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, wo) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var wo tp.WorkOrder
	if err := c.BindJSON(&wo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo.Id = woId
	err = h.app.UpdateWorkOrder(wo.Id, wo.StatusId, wo.CreatedDate, wo.CompletedDate)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
