package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var rootTaskToolRoute = fmt.Sprintf("%s/task-tools", routePrefix)
var taskToolRoute = fmt.Sprintf("%s/tools", indTaskRoute)
var toolId = "ToolId"
var indTaskToolRoute = fmt.Sprintf("%s/:%s", taskToolRoute, toolId)

func (h *Api) registerTaskToolRoutes() {
	h.router.POST(rootTaskToolRoute, h.createTaskToolBody)
	h.router.POST(indTaskToolRoute, h.createTaskToolPath)
	h.router.DELETE(indTaskToolRoute, h.deleteTaskTool)
	h.router.GET(taskToolRoute, h.listTaskTools)
	h.router.GET(indTaskToolRoute, h.getTaskTool)
}

// CreateTaskToolBody godoc
//
//	@Summary		Create a relationship between an asset task and a tool with json body
//	@Description	Create a relationship between an asset task and a tool with json body
//	@Accept			json
//	@Param			taskTool	body	tp.TaskTool	true	"Asset Task Tool object"
//	@Produce		json
//	@Success		201	{object}	tp.TaskTool
//	@Router			/task-tools [post]
func (h *Api) createTaskToolBody(c *gin.Context) {
	var taskTool tp.TaskTool
	if err := c.BindJSON(&taskTool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskTool, err := h.app.CreateTaskTool(taskTool)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskTool))
}

//	CreateTaskToolPath godoc
//
// @Summary		Create a relationship between an asset task and a tool with path parameters
// @Description	Create a relationship between an asset task and a tool with path parameters
// @Param			groupTitle	path	string	true	"Group Title"
// @Param			assetTitle	path	string	true	"Asset Title"
// @Param			taskId		path	string	true	"Asset Task Id"
// @Param			toolId		path	string	true	"Tool Id"
// @Produce		json
// @Success		201	{object}	tp.TaskTool
// @Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/tools/{toolId} [post]
func (h *Api) createTaskToolPath(c *gin.Context) {
	taskTool, err := h.app.CreateTaskToolWithValidation(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskTool))
}

// DeleteTaskTool godoc
//
//	@Summary		Delete a relationship between an asset task and a tool
//	@Description	Delete a relationship between an asset task and a tool
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Param			toolId		path	string	true	"Tool Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/tools/{toolId} [delete]
func (h *Api) deleteTaskTool(c *gin.Context) {
	err := h.app.DeleteTaskTool(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTaskTools godoc
//
//	@Summary		List asset task tools
//	@Description	List all asset task tools
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.TaskTool
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/tools [get]
func (h *Api) listTaskTools(c *gin.Context) {
	taskTools, err := h.app.ListTaskTools(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTools))
}

// GetTaskTool godoc
//
//	@Summary		Get an asset task tool
//	@Description	Get an asset task tool
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Param			toolId		path	string	true	"Tool Id"
//	@Produce		json
//	@Success		200	{object}	tp.TaskTool
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/tools/{toolId} [get]
func (h *Api) getTaskTool(c *gin.Context) {
	taskTool, err := h.app.GetTaskTool(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(toolId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTool))
}
