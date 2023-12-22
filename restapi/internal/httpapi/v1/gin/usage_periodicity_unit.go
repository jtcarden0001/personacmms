package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerUsagePeriodicityUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/usage-periodicity-units", routePrefix)
	individualRoute := fmt.Sprintf("%s/:usagePeriodicityUnitId", baseRoute)

	h.router.POST(baseRoute, h.createUsagePeriodicityUnit)
	h.router.DELETE(individualRoute, h.deleteUsagePeriodicityUnit)
	h.router.GET(baseRoute, h.getAllUsagePeriodicityUnit)
	h.router.GET(individualRoute, h.getUsagePeriodicityUnit)
	h.router.PUT(individualRoute, h.updateUsagePeriodicityUnit)
}

func (h *HttpApi) createUsagePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteUsagePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllUsagePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getUsagePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateUsagePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
