package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTimeUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/time-units", routePrefix)
	individualRoute := fmt.Sprintf("%s/:timeUnitTitle", baseRoute)

	h.router.POST(baseRoute, h.createTimeUnit)
	h.router.DELETE(individualRoute, h.deleteTimeUnit)
	h.router.GET(baseRoute, h.getAllTimeUnit)
	h.router.GET(individualRoute, h.getTimeUnit)
	h.router.PUT(individualRoute, h.updateTimeUnit)
}

func (h *Api) createTimeUnit(c *gin.Context) {
	var tpu tp.TimeUnit
	if err := c.BindJSON(&tpu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTimeUnit(tpu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		tpu.Id = id
		c.IndentedJSON(201, tpu) // switch to .JSON() for performance
	}
}

func (h *Api) deleteTimeUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timeUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTimeUnit(tpuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllTimeUnit(c *gin.Context) {
	tpus, err := h.app.GetAllTimeUnit()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tpus) // switch to .JSON() for performance
	}
}

func (h *Api) getTimeUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timeUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tpu, err := h.app.GetTimeUnit(tpuId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, tpu) // switch to .JSON() for performance
	}
}

func (h *Api) updateTimeUnit(c *gin.Context) {
	tpuId, err := strconv.Atoi(c.Param("timeUnitId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var tpu tp.TimeUnit
	if err := c.BindJSON(&tpu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateTimeUnit(tpuId, tpu.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
