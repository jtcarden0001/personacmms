package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var rootAssetTaskToolRoute = fmt.Sprintf("%s/asset-task-tools", routePrefix)
var assetTaskToolRoute = fmt.Sprintf("%s/tools", baseAssetTaskRoute)
var toolId = "ToolId"
var indAssetTaskToolRoute = fmt.Sprintf("%s/:%s", assetTaskToolRoute, toolId)

func (h *Api) registerAssetTaskToolRoutes() {
	h.router.POST(rootAssetTaskToolRoute, h.createAssetTaskToolBody)
	h.router.POST(indAssetTaskToolRoute, h.createAssetTaskToolPath)
	h.router.DELETE(indAssetTaskToolRoute, h.deleteAssetTaskTool)
	h.router.GET(assetTaskToolRoute, h.listAssetTaskTools)
	h.router.GET(indAssetTaskToolRoute, h.getAssetTaskTool)
}

// CreateAssetTaskToolBody godoc
//
//	@Summary		Create a relationship between an asset task and a tool with json body
//	@Description	Create a relationship between an asset task and a tool with json body
//	@Accept			json
//	@Param 			assetTaskTool 	body 	tp.AssetTaskTool 	true 	"Asset Task Tool object"
//	@Produce		json
//	@Success		201	{object}	tp.AssetTaskTool
//	@Router			/asset-task-tools [post]
func (h *Api) createAssetTaskToolBody(c *gin.Context) {
	var assetTaskTool tp.AssetTaskTool
	if err := c.BindJSON(&assetTaskTool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTaskTool, err := h.app.CreateAssetTaskTool(assetTaskTool)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskTool))
}

//	 CreateAssetTaskToolPath godoc
//
//		@Summary		Create a relationship between an asset task and a tool with path parameters
//		@Description	Create a relationship between an asset task and a tool with path parameters
//		@Param			groupTitle	path	string	true	"Group Title"
//		@Param			assetTitle	path	string	true	"Asset Title"
//		@Param			assetTaskId	path	string	true	"Asset Task Id"
//		@Param			toolId	path	string	true	"Tool Id"
//		@Produce		json
//		@Success		201	{object}	tp.AssetTaskTool
//	 @Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/tools/{toolId} [post]
func (h *Api) createAssetTaskToolPath(c *gin.Context) {
	assetTaskTool, err := h.app.CreateAssetTaskToolWithValidation(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskTool))
}

// DeleteAssetTaskTool godoc
//
//	@Summary		Delete a relationship between an asset task and a tool
//	@Description	Delete a relationship between an asset task and a tool
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			toolId		path	string	true	"Tool Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/tools/{toolId} [delete]
func (h *Api) deleteAssetTaskTool(c *gin.Context) {
	err := h.app.DeleteAssetTaskTool(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListAssetTaskTools godoc
//
//	@Summary		List asset task tools
//	@Description	List all asset task tools
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.AssetTaskTool
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/tools [get]
func (h *Api) listAssetTaskTools(c *gin.Context) {
	assetTaskTools, err := h.app.ListAssetTaskTools(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTaskTools))
}

// GetAssetTaskTool godoc
//
//	@Summary		Get an asset task tool
//	@Description	Get an asset task tool
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			toolId		path	string	true	"Tool Id"
//	@Produce		json
//	@Success		200	{object}	tp.AssetTaskTool
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/tools/{toolId} [get]
func (h *Api) getAssetTaskTool(c *gin.Context) {
	assetTaskTool, err := h.app.GetAssetTaskTool(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTaskTool))
}
