package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:id", baseRoute)

	h.router.POST(baseRoute, h.createConsumable)
	h.router.DELETE(individualRoute, h.deleteConsumable)
	h.router.GET(baseRoute, h.getAllConsumable)
	h.router.GET(individualRoute, h.getConsumable)
	h.router.PUT(individualRoute, h.updateConsumable)
}

func (h *HttpApi) createConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateConsumable(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
