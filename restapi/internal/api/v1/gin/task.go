package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTaskRoutes() {
	baseAllRoute := fmt.Sprintf("%s/tasks", routePrefix)
	baseEqRoute := fmt.Sprintf("%s/asset/:assetId/tasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:taskId", baseEqRoute)

	h.router.POST(baseEqRoute, h.createTask)
	h.router.DELETE(individualRoute, h.deleteTask)
	h.router.GET(baseAllRoute, h.getAllTask)
	h.router.GET(baseEqRoute, h.getAllTaskByAsset)
	h.router.GET(individualRoute, h.getTask)
	h.router.PUT(individualRoute, h.updateTask)
}

func (h *Api) createTask(c *gin.Context) {
	var t tp.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTask(t.Title, t.Instructions, t.TimePeriodicityQuantity, t.TimePeriodicityUnitId, t.UsagePeriodicityQuantity, t.UsagePeriodicityUnitId, assetId)
	if err != nil {
		processAppError(c, err)
	} else {
		t.Id = id
		c.IndentedJSON(201, t) // switch to .JSON() for performance
	}
}

func (h *Api) deleteTask(c *gin.Context) {
	taskid, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTask(taskid)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllTask(c *gin.Context) {
	tasks, err := h.app.GetAllTask()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tasks) // switch to .JSON() for performance
	}
}

func (h *Api) getAllTaskByAsset(c *gin.Context) {
	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.app.GetAllTaskByAssetId(assetId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tasks) // switch to .JSON() for performance
	}
}

func (h *Api) getTask(c *gin.Context) {
	taskid, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := h.app.GetTask(taskid)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, task) // switch to .JSON() for performance
	}
}

func (h *Api) updateTask(c *gin.Context) {
	var t tp.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskid, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateTask(taskid, t.Title, t.Instructions, t.TimePeriodicityQuantity, t.TimePeriodicityUnitId, t.UsagePeriodicityQuantity, t.UsagePeriodicityUnitId, assetId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
