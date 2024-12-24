package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /tools (JSON)
// - GET  /tools/{toolId}
// - GET  /tools
// - PUT  /tools/{toolId} (JSON)
// - DEL  /tools/{toolId}
//
// - PUT  /assets/{assetId}/tasks/{taskId}/tools/{toolId}
// - DEL  /assets/{assetId}/tasks/{taskId}/tools/{toolId}
// - PUT  /assets/{assetId}/work-orders/{workOrderId}/tools/{toolId}
// - DEL  /assets/{assetId}/work-orders/{workOrderId}/tools/{toolId}

var toolId = "toolId"
var toolGp = "tools"
var toolResource = fmt.Sprintf("%s/:%s", toolGp, toolId)

var baseToolRoute = fmt.Sprintf("%s/%s", routePrefix, toolGp)
var indToolRoute = fmt.Sprintf("%s/%s", routePrefix, toolResource)

func (h *Api) registerToolRoutes() {
	h.router.POST(baseToolRoute, h.createTool)

	h.router.DELETE(indToolRoute, h.deleteTool)
	h.router.DELETE(fmt.Sprintf("%s/%s", indTaskRoute, toolResource), h.disassociateToolWithTask)
	h.router.DELETE(fmt.Sprintf("%s/%s", indWorkOrderRoute, toolResource), h.disassociateToolWithWorkOrder)

	h.router.GET(baseToolRoute, h.listTools)
	h.router.GET(indToolRoute, h.getTool)

	h.router.PUT(fmt.Sprintf("%s/%s", indTaskRoute, toolResource), h.associateToolWithTask)
	h.router.PUT(fmt.Sprintf("%s/%s", indWorkOrderRoute, toolResource), h.associateToolWithWorkOrder)
	h.router.PUT(indToolRoute, h.updateTool)
}

// AssociateToolWithTask godoc
//
//	@Summary		Associate a tool with a task
//	@Description	Associate a tool with a task
//	@Tags			tools
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset ID"
//	@Param			taskId	path		string	true	"Task ID"
//	@Param			toolId	path		string	true	"Tool ID"
//	@Success		200		{object}	tp.Tool
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/tools/{toolId} [put]
func (h *Api) associateToolWithTask(c *gin.Context) {
	tool, err := h.app.AssociateToolWithTask(c.Param(assetId), c.Param(taskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}

// AssociateToolWithWorkOrder godoc
//
//	@Summary		Associate a tool with a work order
//	@Description	Associate a tool with a work order
//	@Tags			tools
//	@Produce		json
//	@Param			assetId		path		string	true	"Asset ID"
//	@Param			workOrderId	path		string	true	"Work Order ID"
//	@Param			toolId		path		string	true	"Tool ID"
//	@Success		200			{object}	tp.Tool
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/tools/{toolId} [put]
func (h *Api) associateToolWithWorkOrder(c *gin.Context) {
	tool, err := h.app.AssociateToolWithWorkOrder(c.Param(assetId), c.Param(workOrderId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
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
func (h *Api) createTool(c *gin.Context) {
	var tool tp.Tool
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
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
//	@Param			toolId	path	string	true	"Tool Title"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/tools/{toolId} [delete]
func (h *Api) deleteTool(c *gin.Context) {
	err := h.app.DeleteTool(c.Param(toolId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateToolWithTask godoc
//
//	@Summary		Disassociate a tool with a task
//	@Description	Disassociate a tool with a task
//	@Tags			tools
//	@Param			assetId	path	string	true	"Asset ID"
//	@Param			taskId	path	string	true	"Task ID"
//	@Param			toolId	path	string	true	"Tool ID"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/tools/{toolId} [delete]
func (h *Api) disassociateToolWithTask(c *gin.Context) {
	err := h.app.DisassociateToolWithTask(c.Param(assetId), c.Param(taskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateToolWithWorkOrder godoc
//
//	@Summary		Disassociate a tool with a work order
//	@Description	Disassociate a tool with a work order
//	@Tags			tools
//	@Param			assetId		path	string	true	"Asset ID"
//	@Param			workOrderId	path	string	true	"Work Order ID"
//	@Param			toolId		path	string	true	"Tool ID"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/tools/{toolId} [delete]
func (h *Api) disassociateToolWithWorkOrder(c *gin.Context) {
	err := h.app.DisassociateToolWithWorkOrder(c.Param(assetId), c.Param(workOrderId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// getTool godoc
//
//	@Summary		Get a tool
//	@Description	Get a tool
//	@Tags			tools
//	@Param			toolId	path	string	true	"Tool Title"
//	@Produce		json
//	@Success		200	{object}	tp.Tool
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/tools/{toolId} [get]
func (h *Api) getTool(c *gin.Context) {
	tool, err := h.app.GetTool(c.Param(toolId))
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
func (h *Api) listTools(c *gin.Context) {
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
//	@Param			toolId	path		string	true	"Tool Title"
//	@Param			tool	body		tp.Tool	true	"Tool object"
//	@Success		200		{object}	tp.Tool
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/tools/{toolId} [put]
func (h *Api) updateTool(c *gin.Context) {
	var tool tp.Tool
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	tool, err := h.app.UpdateTool(c.Param(toolId), tool)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tool))
}
