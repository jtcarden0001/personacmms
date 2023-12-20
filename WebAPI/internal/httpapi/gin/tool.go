package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

func (h *HttpApi) registerToolRoutes(r *gin.Engine) {
	r.POST("/tool", h.CreateTool)
	r.DELETE("/tool/:id", h.DeleteTool)
	r.GET("/tool", h.GetAllTools)
	r.GET("/tool/:id", h.GetTool)
	r.PATCH("/tool/:id", h.UpdateTool)
}

func (h *HttpApi) CreateTool(c *gin.Context) {
	var t tp.Tool
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTool(t.Title, t.Size)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		t.Id = id
		c.IndentedJSON(http.StatusCreated, t)
	}
}

func (h *HttpApi) DeleteTool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTool(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusNoContent, gin.H{})
	}
}

func (h *HttpApi) GetAllTools(c *gin.Context) {
	tools, err := h.app.GetAllTools()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, tools)
	}
}

func (h *HttpApi) GetTool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tool, err := h.app.GetTool(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, tool)
	}
}

func (h *HttpApi) UpdateTool(c *gin.Context) {
	var t tp.Tool
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	t.Id = id
	err = h.app.UpdateTool(t.Id, t.Title, t.Size)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, t)
	}
}
