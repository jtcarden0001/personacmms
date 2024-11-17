package gin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerToolRoutes() {
	baseRoute := fmt.Sprintf("%s/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:id", baseRoute)

	h.router.POST(baseRoute, h.CreateTool)
	h.router.DELETE(individualRoute, h.DeleteTool)
	h.router.GET(baseRoute, h.GetAllTools)
	h.router.GET(individualRoute, h.GetTool)
	h.router.PUT(individualRoute, h.UpdateTool)
}

func (h *Api) CreateTool(c *gin.Context) {
	var t tp.Tool
	if err := c.BindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateTool(t.Title, t.Size)
	if err != nil {
		processAppError(c, err)
	} else {
		t.Id = id
		c.IndentedJSON(http.StatusCreated, t)
	}
}

func (h *Api) DeleteTool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteTool(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusNoContent, gin.H{})
	}
}

func (h *Api) GetAllTools(c *gin.Context) {
	tools, err := h.app.GetAllTool()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, tools)
	}
}

func (h *Api) GetTool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tool, err := h.app.GetTool(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, tool)
	}
}

func (h *Api) UpdateTool(c *gin.Context) {
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
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, t)
	}
}
