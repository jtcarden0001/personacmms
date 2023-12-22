package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerTaskConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(baseRoute, h.createTaskConsumable)
	h.router.DELETE(individualRoute, h.deleteTaskConsumable)
	h.router.GET(baseRoute, h.getAllTaskConsumableByTask)
	h.router.GET(individualRoute, h.getTaskConsumable)
	h.router.PUT(individualRoute, h.updateTaskConsumable)
}

func (h *HttpApi) createTaskConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteTaskConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllTaskConsumableByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getTaskConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateTaskConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
