package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerTaskToolRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:toolId", baseRoute)

	h.router.POST(individualRoute, h.createTaskTool)
	h.router.DELETE(individualRoute, h.deleteTaskTool)
	h.router.GET(baseRoute, h.getAllTaskToolByTask)
	h.router.GET(individualRoute, h.getTaskTool)
}

func (h *HttpApi) createTaskTool(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.CreateTaskTool(taskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(201, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteTaskTool(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTaskTool(taskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllTaskToolByTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskTools, err := h.app.GetAllTaskToolByTaskId(taskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, taskTools) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getTaskTool(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskTool, err := h.app.GetTaskTool(taskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, taskTool) // switch to .JSON() for performance
	}
}
