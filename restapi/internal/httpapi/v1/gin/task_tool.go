package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerTaskToolRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:toolId", baseRoute)

	h.router.POST(baseRoute, h.createTaskTool)
	h.router.DELETE(individualRoute, h.deleteTaskTool)
	h.router.GET(baseRoute, h.getAllTaskToolByTask)
	h.router.GET(individualRoute, h.getTaskTool)
}

func (h *HttpApi) createTaskTool(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteTaskTool(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllTaskToolByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getTaskTool(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateTaskTool(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
