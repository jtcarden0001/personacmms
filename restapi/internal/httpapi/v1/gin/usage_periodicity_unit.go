package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
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
	var upu tp.UsagePeriodicityUnit
	if err := c.BindJSON(&upu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateUsagePeriodicityUnit(upu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		upu.Id = id
		c.IndentedJSON(201, upu) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteUsagePeriodicityUnit(c *gin.Context) {
	upuId, err := strconv.Atoi(c.Param("usagePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteUsagePeriodicityUnit(upuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllUsagePeriodicityUnit(c *gin.Context) {
	upus, err := h.app.GetAllUsagePeriodicityUnit()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, upus) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getUsagePeriodicityUnit(c *gin.Context) {
	upuId, err := strconv.Atoi(c.Param("usagePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	upu, err := h.app.GetUsagePeriodicityUnit(upuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, upu) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateUsagePeriodicityUnit(c *gin.Context) {
	upuId, err := strconv.Atoi(c.Param("usagePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var upu tp.UsagePeriodicityUnit
	if err := c.BindJSON(&upu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	upu.Id = upuId
	err = h.app.UpdateUsagePeriodicityUnit(upu.Id, upu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, upu) // switch to .JSON() for performance
	}
}
