package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

func (h *HttpApi) registerEquipmentRoutes(r *gin.Engine) {
	r.POST("/equipment", h.CreateEquipment)
	r.DELETE("/equipment/:id", h.DeleteEquipment)
	r.GET("/equipment", h.GetAllEquipment)
	r.GET("/equipment/:id", h.GetEquipment)
	r.PATCH("/equipment/:id", h.UpdateEquipment)
}

func (h *HttpApi) CreateEquipment(c *gin.Context) {
	var e tp.Equipment
	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateEquipment(e.Title, e.Year, e.Make, e.ModelNumber, e.Description)
	if err != nil {
		// TODO: revisit this error to make sure it is correct
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		e.Id = id
		c.IndentedJSON(http.StatusCreated, e)
	}
}

func (h *HttpApi) DeleteEquipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

func (h *HttpApi) GetAllEquipment(c *gin.Context) {
	equipment, err := h.app.GetAllEquipment()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, equipment)
	}
}

func (h *HttpApi) GetEquipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

func (h *HttpApi) UpdateEquipment(c *gin.Context) {
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
	err = h.app.UpdateEquipment(e.Id, e.Title, e.Year, e.Make, e.ModelNumber, e.Description)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, e)
	}
}
