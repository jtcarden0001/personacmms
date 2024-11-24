package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Api) registerTaskToolRoutes() {
	baseRoute := fmt.Sprintf("%s/asset/:assetId/preventativeTasks/:preventativeTaskId/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:toolId", baseRoute)

	h.router.POST(individualRoute, h.createTaskTool)
	h.router.DELETE(individualRoute, h.deleteTaskTool)
	h.router.GET(baseRoute, h.getAllTaskToolByTask)
	h.router.GET(individualRoute, h.getTaskTool)
}

func (h *Api) createTaskTool(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.CreateTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(201, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) deleteTaskTool(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllTaskToolByTask(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskTools, err := h.app.GetAllTaskToolByTaskId(preventativeTaskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskTools) // switch to .JSON() for performance
	}
}

func (h *Api) getTaskTool(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	toolId, err := strconv.Atoi(c.Param("toolId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskTool, err := h.app.GetTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskTool) // switch to .JSON() for performance
	}
}
