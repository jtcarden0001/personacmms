package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerWorkOrderRoutes() {
	baseRouteByTask := fmt.Sprintf("%s/equipment/:equipmentId/tasks/:taskId/work-orders", routePrefix)
	individualRouteByTask := fmt.Sprintf("%s/:workOrderId", baseRouteByTask)

	h.router.POST(baseRouteByTask, h.createWorkOrderByTask)
	h.router.DELETE(individualRouteByTask, h.deleteWorkOrderByTask)
	h.router.GET(baseRouteByTask, h.getAllWorkOrderByTask)
	h.router.GET(individualRouteByTask, h.getWorkOrderByTask)
	h.router.PUT(individualRouteByTask, h.updateWorkOrderByTask)
}

func (h *HttpApi) createWorkOrderByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteWorkOrderByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllWorkOrderByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getWorkOrderByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateWorkOrderByTask(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
