package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTaskConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/asset/:assetId/preventativeTasks/:preventativeTaskId/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(individualRoute, h.createTaskConsumable)
	h.router.DELETE(individualRoute, h.deleteTaskConsumable)
	h.router.GET(baseRoute, h.getAllTaskConsumableByTask)
	h.router.GET(individualRoute, h.getTaskConsumable)
	h.router.PUT(individualRoute, h.updateTaskConsumable)
}

func (h *Api) createTaskConsumable(c *gin.Context) {
	tc := tp.TaskConsumable{}
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.TaskId = preventativeTaskId

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

func (h *Api) deleteTaskConsumable(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTaskConsumable(preventativeTaskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllTaskConsumableByTask(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskConsumables, err := h.app.GetAllTaskConsumableByTaskId(preventativeTaskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskConsumables) // switch to .JSON() for performance
	}
}

func (h *Api) getTaskConsumable(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskConsumable, err := h.app.GetTaskConsumable(preventativeTaskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskConsumable) // switch to .JSON() for performance
	}
}

func (h *Api) updateTaskConsumable(c *gin.Context) {
	var tc tp.TaskConsumable
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.TaskId = preventativeTaskId

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
