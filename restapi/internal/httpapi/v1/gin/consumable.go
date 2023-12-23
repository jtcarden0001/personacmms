package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(baseRoute, h.createConsumable)
	h.router.DELETE(individualRoute, h.deleteConsumable)
	h.router.GET(baseRoute, h.getAllConsumable)
	h.router.GET(individualRoute, h.getConsumable)
	h.router.PUT(individualRoute, h.updateConsumable)
}

func (h *HttpApi) createConsumable(c *gin.Context) {
	var co tp.Consumable
	if err := c.BindJSON(&co); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateConsumable(co.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		co.Id = id
		c.IndentedJSON(201, co)
	}
}

func (h *HttpApi) deleteConsumable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteConsumable(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{})
	}
}

func (h *HttpApi) getAllConsumable(c *gin.Context) {
	consumables, err := h.app.GetAllConsumable()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, consumables)
	}
}

func (h *HttpApi) getConsumable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumable, err := h.app.GetConsumable(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, consumable)
	}
}

func (h *HttpApi) updateConsumable(c *gin.Context) {
	var co tp.Consumable

	if err := c.BindJSON(&co); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateConsumable(id, co.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{})
	}
}
