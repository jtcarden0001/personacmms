package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseToolRoute = fmt.Sprintf("%s/tools", routePrefix)
var toolTitle = "ToolTitle"
var individualToolRoute = fmt.Sprintf("%s/:%s", baseToolRoute, toolTitle)

func (h *Api) registerToolRoutes() {

	h.router.POST(baseToolRoute, h.CreateTool)
	h.router.DELETE(individualToolRoute, h.DeleteTool)
	h.router.GET(baseToolRoute, h.ListTools)
	h.router.GET(individualToolRoute, h.GetTool)
	h.router.PUT(individualToolRoute, h.UpdateTool)
}

// CreateTool godoc
//
//	@Summary		Create a tool
//	@Description	Create a tool
//	@Tags			tools
//	@Accept			json
//	@Produce		json
//	@Param			tool	body		tp.Tool	true	"Tool object"
//	@Success		201		{object}	tp.Tool
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/tools [post]
func (h *Api) CreateTool(c *gin.Context) {
	var tool tp.Tool
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tool, err := h.app.CreateTool(tool)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, tool))
}

// DeleteTool godoc
//
//	@Summary		Delete a tool
//	@Description	Delete a tool
//	@Tags			tools
//	@Param			toolTitle	path	string	true	"Tool Title"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/tools/{toolTitle} [delete]
func (h *Api) DeleteTool(c *gin.Context) {
	err := h.app.DeleteTool(c.Param(toolTitle))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTool godoc
//
//	@Summary		Get a tool
//	@Description	Get a tool
//	@Tags			tools
//	@Param			toolTitle	path	string	true	"Tool Title"
//	@Produce		json
//	@Success		200	{object}	tp.Tool
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/tools/{toolTitle} [get]
func (h *Api) GetTool(c *gin.Context) {
	tool, err := h.app.GetTool(c.Param(toolTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}

// ListTools godoc
//
//	@Summary		List tools
//	@Description	List all tools
//	@Tags			tools
//	@Produce		json
//	@Success		200	{object}	[]tp.Tool
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/tools [get]
func (h *Api) ListTools(c *gin.Context) {
	tools, err := h.app.ListTools()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tools))
}

// UpdateTool godoc
//
//	@Summary		Update a tool
//	@Description	Update a tool
//	@Tags			tools
//	@Accept			json
//	@Produce		json
//	@Param			toolTitle	path		string	true	"Tool Title"
//	@Param			tool		body		tp.Tool	true	"Tool object"
//	@Success		200			{object}	tp.Tool
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/tools/{toolTitle} [put]
func (h *Api) UpdateTool(c *gin.Context) {
	var tool tp.Tool
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tool, err := h.app.UpdateTool(c.Param(toolTitle), tool)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}
