package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *HttpApi) registerTaskRoutes() {
	baseAllRoute := fmt.Sprintf("%s/tasks", routePrefix)
	baseEqRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:taskId", baseEqRoute)

	h.router.POST(baseEqRoute, h.createTask)
	h.router.DELETE(individualRoute, h.deleteTask)
	h.router.GET(baseAllRoute, h.getAllTask)
	h.router.GET(baseEqRoute, h.getAllTaskByEquipment)
	h.router.GET(individualRoute, h.getTask)
	h.router.PUT(individualRoute, h.updateTask)
}

func (h *HttpApi) createTask(c *gin.Context) {
	var t tp.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	equipmentId, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTask(t.Title, t.Instructions, t.TimePeriodicityQuantity, t.TimePeriodicityUnitId, t.UsagePeriodicityQuantity, t.UsagePeriodicityUnitId, equipmentId)
	if err != nil {
		processAppError(c, err)
	} else {
		t.Id = id
		c.IndentedJSON(201, t) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteTask(c *gin.Context) {
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

func (h *HttpApi) getAllTask(c *gin.Context) {
	tasks, err := h.app.GetAllTask()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tasks) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllTaskByEquipment(c *gin.Context) {
	equipmentId, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.app.GetAllTaskByEquipmentId(equipmentId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tasks) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getTask(c *gin.Context) {
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

func (h *HttpApi) updateTask(c *gin.Context) {
	var t tp.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	equipmentId, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskid, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateTask(taskid, t.Title, t.Instructions, t.TimePeriodicityQuantity, t.TimePeriodicityUnitId, t.UsagePeriodicityQuantity, t.UsagePeriodicityUnitId, equipmentId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
