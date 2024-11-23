package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerToolRoutes() {
	baseRoute := fmt.Sprintf("%s/tools", routePrefix)
	individualRoute := fmt.Sprintf("%s/:toolTitle", baseRoute)

	h.router.POST(baseRoute, h.CreateTool)
	h.router.DELETE(individualRoute, h.DeleteTool)
	h.router.GET(baseRoute, h.ListTools)
	h.router.GET(individualRoute, h.GetTool)
	h.router.PUT(individualRoute, h.UpdateTool)
}

// CreateTool godoc
//
//	@Summary		Create a tool
//	@Description	Create a tool
//	@Accept			json
//	@Param			tool	body	tp.Tool	true	"Tool object"
//	@Produce		json
//	@Success		201	{object}	tp.Tool
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
//	@Param			toolTitle	path	string	true	"Tool Title"
//	@Success		204
//	@Failure		404
//	@Router			/tools/{toolTitle} [delete]
func (h *Api) DeleteTool(c *gin.Context) {
	err := h.app.DeleteTool(c.Param("toolTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTools godoc
//
//	@Summary		List tools
//	@Description	List all tools
//	@Produce		json
//	@Success		200	{object}	[]tp.Tool
//	@Router			/tools [get]
func (h *Api) ListTools(c *gin.Context) {
	tools, err := h.app.ListTools()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tools))
}

// GetTool godoc
//
//	@Summary		Get a tool
//	@Description	Get a tool
//	@Param			toolTitle	path	string	true	"Tool Title"
//	@Produce		json
//	@Success		200	{object}	tp.Tool
//	@Router			/tools/{toolTitle} [get]
func (h *Api) GetTool(c *gin.Context) {
	tool, err := h.app.GetTool(c.Param("toolTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}

// UpdateTool godoc
//
//	@Summary		Update a tool
//	@Description	Update a tool
//	@Accept			json
//	@Param			toolTitle	path	string	true	"Tool Title"
//	@Produce		json
//	@Success		200	{object}	tp.Tool
//	@Router			/tools/{toolTitle} [put]
func (h *Api) UpdateTool(c *gin.Context) {
	var tool tp.Tool
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tool, err := h.app.UpdateTool(c.Param("toolTitle"), tool)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}
