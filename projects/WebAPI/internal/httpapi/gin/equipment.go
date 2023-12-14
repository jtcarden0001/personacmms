package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

func (h *HttpApi) registerEquipmentRoutes(r *gin.Engine) {
	r.GET("/equipment", h.GetAllEquipment)
	r.POST("/equipment", h.CreateEquipment)
}

func (h *HttpApi) CreateEquipment(c *gin.Context) {
	var e tp.Equipment
	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.app.CreateEquipment(e.Title, e.Description)
	if err != nil {
		c.IndentedJSON(http.StatusCreated, e)
	} else {
		// TODO: revisit this error to make sure it is correct
		c.JSON(400, gin.H{"error": err.Error()})
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
