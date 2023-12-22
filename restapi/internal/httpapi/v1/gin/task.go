package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:taskId", baseRoute)

	h.router.POST(baseRoute, h.createTask)
	h.router.DELETE(individualRoute, h.deleteTask)
	h.router.GET(baseRoute, h.getAllTaskByEquipment)
	h.router.GET(individualRoute, h.getTask)
	h.router.PUT(individualRoute, h.updateTask)
}

func (h *HttpApi) createTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllTaskByEquipment(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
