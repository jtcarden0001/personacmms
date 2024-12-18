package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseTaskRoute = fmt.Sprintf("%s/tasks", indAssetRoute)
var taskId = "TaskId"
var indTaskRoute = fmt.Sprintf("%s/:%s", baseTaskRoute, taskId)

// general route that treats all asset tasks the same, can CRUD any type through this route
func (h *Api) registerTaskRoutes() {
	h.router.POST(baseTaskRoute, h.createTask)
	h.router.DELETE(indTaskRoute, h.deleteTask)
	h.router.GET(baseTaskRoute, h.listTasks)
	h.router.GET(indTaskRoute, h.getTask)
	h.router.PUT(indTaskRoute, h.updateTask)
}

// CreateTask godoc
//
//	@Summary		Create an asset task
//	@Description	Create an asset task
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			task	body	tp.Task	true	"Asset Task object"
//	@Produce		json
//	@Success		201	{object}	tp.Task
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks [post]
func (h *Api) createTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.app.CreateTask(c.Param(groupTitle), c.Param(assetTitle), task)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, task))
}

// DeleteTask godoc
//
//	@Summary		Delete an asset task
//	@Description	Delete an asset task
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId} [delete]
func (h *Api) deleteTask(c *gin.Context) {
	err := h.app.DeleteTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTasks godoc
//
//	@Summary		List asset tasks
//	@Description	List all asset tasks
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.Task
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks [get]
func (h *Api) listTasks(c *gin.Context) {
	tasks, err := h.app.ListTasks(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tasks))
}

// GetTask godoc
//
//	@Summary		Get an asset task
//	@Description	Get an asset task
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	tp.Task
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId} [get]
func (h *Api) getTask(c *gin.Context) {
	task, err := h.app.GetTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, task))
}

// UpdateTask godoc
//
//	@Summary		Update an asset task
//	@Description	Update an asset task
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Param			task	body	tp.Task	true	"Asset Task object"
//	@Produce		json
//	@Success		200	{object}	tp.Task
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId} [put]
func (h *Api) updateTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.app.UpdateTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), task)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, task))
}
