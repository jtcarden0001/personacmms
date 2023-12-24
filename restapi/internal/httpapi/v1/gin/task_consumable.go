package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerTaskConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(individualRoute, h.createTaskConsumable)
	h.router.DELETE(individualRoute, h.deleteTaskConsumable)
	h.router.GET(baseRoute, h.getAllTaskConsumableByTask)
	h.router.GET(individualRoute, h.getTaskConsumable)
	h.router.PUT(individualRoute, h.updateTaskConsumable)
}

func (h *HttpApi) createTaskConsumable(c *gin.Context) {
	tc := tp.TaskConsumable{}
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.TaskId = taskId

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.ConsumableId = consumableId

	err = h.app.CreateTaskConsumable(tc.TaskId, tc.ConsumableId, tc.QuantityNote)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(201, tc) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteTaskConsumable(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTaskConsumable(taskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllTaskConsumableByTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskConsumables, err := h.app.GetAllTaskConsumableByTaskId(taskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, taskConsumables) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getTaskConsumable(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskConsumable, err := h.app.GetTaskConsumable(taskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, taskConsumable) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateTaskConsumable(c *gin.Context) {
	var tc tp.TaskConsumable
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.TaskId = taskId

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.ConsumableId = consumableId

	err = h.app.UpdateTaskConsumable(tc.TaskId, tc.ConsumableId, tc.QuantityNote)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
