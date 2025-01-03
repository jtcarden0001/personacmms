package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /assets/{assetId}/tasks (JSON) done
// - GET  /assets/{assetId}/tasks/{taskId} done
// - GET  /assets/{assetId}/tasks done
// - PUT  /assets/{assetId}/tasks/{taskId} (JSON) done
// - DEL  /assets/{assetId}/tasks/{taskId} done
//
// - DEL  /assets/{assetId}/work-orders/{workOrderId}/tasks done

var taskId = "TaskId"
var taskGp = "tasks"
var taskResource = fmt.Sprintf("%s/:%s", taskGp, taskId)

var baseTaskRoute = fmt.Sprintf("%s/%s", indAssetRoute, taskGp)
var indTaskRoute = fmt.Sprintf("%s/%s", indAssetRoute, taskResource)

func (h *Api) registerTaskRoutes() {
	h.router.POST(baseTaskRoute, h.createTask)

	h.router.DELETE(indTaskRoute, h.deleteTask)
	h.router.DELETE(fmt.Sprintf("%s/%s/%s", indAssetRoute, workOrderResource, taskGp), h.disassociateTaskWithWorkOrder)

	h.router.GET(indTaskRoute, h.getTask)
	h.router.GET(baseTaskRoute, h.listTasksByAsset)

	h.router.PUT(indTaskRoute, h.updateTask)
}

// CreateTask godoc
//
//	@Summary		Create a task
//	@Description	Create a task for an asset
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			task	body		tp.Task	true	"Task object"
//	@Success		201		{object}	tp.Task
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks [post]
func (h *Api) createTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	task, err := h.app.CreateTask(c.Param(assetId), task)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, task))
}

// DeleteTask godoc
//
//	@Summary		Delete a task
//	@Description	Delete a task
//	@Tags			tasks
//	@Param			assetId	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Task Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId} [delete]
func (h *Api) deleteTask(c *gin.Context) {
	err := h.app.DeleteTask(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateTaskWithWorkOrder godoc
//
//	@Summary		Disassociate a task with a work order
//	@Description	Disassociate a task with a work order
//	@Tags			tasks
//	@Param			assetId		path	string	true	"Asset Id"
//	@Param			workOrderId	path	string	true	"Work Order Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/tasks [delete]
func (h *Api) disassociateTaskWithWorkOrder(c *gin.Context) {
	err := h.app.DisassociateTaskWithWorkOrder(c.Param(assetId), c.Param(workOrderId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTask godoc
//
//	@Summary		Get a task
//	@Description	Get a task
//	@Tags			tasks
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			taskId	path		string	true	"Task Id"
//	@Success		200		{object}	tp.Task
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId} [get]
func (h *Api) getTask(c *gin.Context) {
	task, err := h.app.GetTask(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, task))
}

// ListTasksByAsset godoc
//
//	@Summary		List tasks by asset
//	@Description	List tasks by asset
//	@Tags			tasks
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Success		200		{object}	[]tp.Task
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks [get]
func (h *Api) listTasksByAsset(c *gin.Context) {
	tasks, err := h.app.ListTasksByAsset(c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tasks))
}

// UpdateTask godoc
//
//	@Summary		Update a task
//	@Description	Update a task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			taskId	path		string	true	"Task Id"
//	@Param			task	body		tp.Task	true	"Task object"
//	@Success		200		{object}	tp.Task
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId} [put]
func (h *Api) updateTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	task, err := h.app.UpdateTask(c.Param(assetId), c.Param(taskId), task)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, task))
}
