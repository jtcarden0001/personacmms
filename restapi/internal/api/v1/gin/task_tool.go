package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Api) registerPreventativeTaskToolRoutes() {
	baseRoute := fmt.Sprintf("%s/asset/:assetId/preventativeTasks/:preventativeTaskId/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:toolId", baseRoute)

	h.router.POST(individualRoute, h.createPreventativeTaskTool)
	h.router.DELETE(individualRoute, h.deletePreventativeTaskTool)
	h.router.GET(baseRoute, h.getAllPreventativeTaskToolByPreventativeTask)
	h.router.GET(individualRoute, h.getPreventativeTaskTool)
}

func (h *Api) createPreventativeTaskTool(c *gin.Context) {
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

	err = h.app.CreatePreventativeTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(201, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) deletePreventativeTaskTool(c *gin.Context) {
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

	err = h.app.DeletePreventativeTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllPreventativeTaskToolByPreventativeTask(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskTools, err := h.app.GetAllPreventativeTaskToolByPreventativeTaskId(preventativeTaskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskTools) // switch to .JSON() for performance
	}
}

func (h *Api) getPreventativeTaskTool(c *gin.Context) {
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

	preventativeTaskTool, err := h.app.GetPreventativeTaskTool(preventativeTaskId, toolId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskTool) // switch to .JSON() for performance
	}
}
