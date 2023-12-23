package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerEquipmentCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryId", baseRoute)

	h.router.POST(baseRoute, h.createEquipmentCategory)
	h.router.DELETE(individualRoute, h.deleteEquipmentCategory)
	h.router.GET(baseRoute, h.getAllEquipmentCategory)
	h.router.GET(individualRoute, h.getEquipmentCategory)
	h.router.PUT(individualRoute, h.updateEquipmentCategory) // accepts object id in url, disregards id in body, may revisit this design
}

func (h *HttpApi) createEquipmentCategory(c *gin.Context) {
	var ec tp.EquipmentCategory
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateEquipmentCategory(ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		ec.Id = id
		c.IndentedJSON(201, ec)
	}
}

func (h *HttpApi) deleteEquipmentCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteEquipmentCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{})
	}
}

func (h *HttpApi) getAllEquipmentCategory(c *gin.Context) {
	equipmentCategories, err := h.app.GetAllEquipmentCategory()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, equipmentCategories)
	}
}

func (h *HttpApi) getEquipmentCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	equipmentCategory, err := h.app.GetEquipmentCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, equipmentCategory)
	}
}

func (h *HttpApi) updateEquipmentCategory(c *gin.Context) {
	var ec tp.EquipmentCategory
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateEquipmentCategory(id, ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{})
	}
}
