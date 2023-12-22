package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerTimePeriodicityUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/time-periodicity-units", routePrefix)
	individualRoute := fmt.Sprintf("%s/:timePeriodicityUnitId", baseRoute)

	h.router.POST(baseRoute, h.createTimePeriodicityUnit)
	h.router.DELETE(individualRoute, h.deleteTimePeriodicityUnit)
	h.router.GET(baseRoute, h.getAllTimePeriodicityUnit)
	h.router.GET(individualRoute, h.getTimePeriodicityUnit)
	h.router.PUT(individualRoute, h.updateTimePeriodicityUnit)
}

func (h *HttpApi) createTimePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteTimePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllTimePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getTimePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateTimePeriodicityUnit(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
