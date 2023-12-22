package gin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerEquipmentRoutes() {
	baseRoute := fmt.Sprintf("%s/equipment", routePrefix)
	individualRoute := fmt.Sprintf("%s/:equipmentId", baseRoute)

	h.router.POST(baseRoute, h.createEquipment)
	h.router.DELETE(individualRoute, h.deleteEquipment)
	h.router.GET(baseRoute, h.getAllEquipment)
	h.router.GET(individualRoute, h.getEquipment)
	h.router.PUT(individualRoute, h.updateEquipment)
}

func (h *HttpApi) createEquipment(c *gin.Context) {
	var e tp.Equipment
	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateEquipment(e.Title, e.Year, e.Make, e.ModelNumber, e.Description, e.CategoryId)
	if err != nil {
		// TODO: revisit this error to make sure it is correct
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		e.Id = id
		c.IndentedJSON(http.StatusCreated, e)
	}
}

func (h *HttpApi) deleteEquipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteEquipment(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusNoContent, gin.H{})
	}
}

func (h *HttpApi) getAllEquipment(c *gin.Context) {
	equipment, err := h.app.GetAllEquipment()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, equipment)
	}
}

func (h *HttpApi) getEquipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	equipment, err := h.app.GetEquipment(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, equipment)
	}
}

func (h *HttpApi) updateEquipment(c *gin.Context) {
	var e tp.Equipment

	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	e.Id = id // ignoring the id in the body and using the id in the url
	err = h.app.UpdateEquipment(e.Id, e.Title, e.Year, e.Make, e.ModelNumber, e.Description, e.CategoryId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, e)
	}
}
