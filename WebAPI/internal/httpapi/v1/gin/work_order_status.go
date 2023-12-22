package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerWorkOrderStatusRoutes() {
	baseRoute := fmt.Sprintf("%s/work-order-statuses", routePrefix)
	individualRoute := fmt.Sprintf("%s/:workOrderStatusId", baseRoute)

	h.router.POST(baseRoute, h.createWorkOrderStatus)
	h.router.DELETE(individualRoute, h.deleteWorkOrderStatus)
	h.router.GET(baseRoute, h.getAllWorkOrderStatus)
	h.router.GET(individualRoute, h.getWorkOrderStatus)
	h.router.PUT(individualRoute, h.updateWorkOrderStatus)
}

func (h *HttpApi) createWorkOrderStatus(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteWorkOrderStatus(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllWorkOrderStatus(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getWorkOrderStatus(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateWorkOrderStatus(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
