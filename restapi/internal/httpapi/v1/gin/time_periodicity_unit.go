package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
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
	var tpu tp.TimePeriodicityUnit
	if err := c.BindJSON(&tpu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTimePeriodicityUnit(tpu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		tpu.Id = id
		c.IndentedJSON(201, tpu) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteTimePeriodicityUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTimePeriodicityUnit(tpuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllTimePeriodicityUnit(c *gin.Context) {
	tpus, err := h.app.GetAllTimePeriodicityUnit()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tpus) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getTimePeriodicityUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tpu, err := h.app.GetTimePeriodicityUnit(tpuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tpu) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateTimePeriodicityUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timePeriodicityUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var tpu tp.TimePeriodicityUnit
	if err := c.BindJSON(&tpu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateTimePeriodicityUnit(tpuId, tpu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
