package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var rootTaskConsumableRoute = fmt.Sprintf("%s/task-consumables", routePrefix)
var taskConsumableRoute = fmt.Sprintf("%s/consumables", indTaskRoute)
var consumableId = "ConsumableId"
var indTaskConsumableRoute = fmt.Sprintf("%s/:%s", taskConsumableRoute, consumableId)

func (h *Api) registerTaskConsumableRoutes() {
	h.router.POST(rootTaskConsumableRoute, h.createTaskConsumableBody)
	h.router.POST(indTaskConsumableRoute, h.createTaskConsumablePath)
	h.router.DELETE(indTaskConsumableRoute, h.deleteTaskConsumable)
	h.router.GET(taskConsumableRoute, h.listTaskConsumables)
	h.router.GET(indTaskConsumableRoute, h.getTaskConsumable)
	h.router.PUT(indTaskConsumableRoute, h.updateTaskConsumableBody)
	h.router.PUT(indTaskConsumableRoute, h.updateTaskConsumablePath)
}

// CreateTaskConsumableBody godoc
//
//	@Summary		Create a relationship between an asset task and a consumable with json body
//	@Description	Create a relationship between an asset task and a consumable with json body
//	@Accept			json
//	@Param			taskConsumable	body	tp.TaskConsumable	true	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.TaskConsumable
//	@Router			/task-consumables [post]
func (h *Api) createTaskConsumableBody(c *gin.Context) {
	var taskConsumable tp.TaskConsumable
	if err := c.BindJSON(&taskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskConsumable, err := h.app.CreateTaskConsumable(taskConsumable)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskConsumable))
}

// CreateTaskConsumablePath godoc
//
// @Summary		Create a relationship between an asset task and a consumable with path parameters
// @Description	Create a relationship between an asset task and a consumable with path parameters
// @Accept			json
// @Param			groupTitle		path	string						true	"Group Title"
// @Param			assetTitle		path	string						true	"Asset Title"
// @Param			taskId			path	string						true	"Asset Task ID"
// @Param			consumableId	path	string						true	"Consumable ID"
// @Param			taskConsumable	body	tp.TaskConsumableForPath	true	"Asset Task Consumable object"
// @Produce		json
// @Success		201	{object}	tp.TaskConsumable
// @Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/consumables/{consumableId} [post]
func (h *Api) createTaskConsumablePath(c *gin.Context) {
	var taskConsumable tp.TaskConsumable
	if err := c.BindJSON(&taskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskConsumable, err := h.app.CreateTaskConsumableWithValidation(
		c.Param(groupTitle),
		c.Param(assetTitle),
		c.Param(taskId),
		c.Param(consumableId),
		taskConsumable.QuantityNote,
	)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskConsumable))
}

// DeleteTaskConsumable godoc
//
//	@Summary		Delete an asset task consumable
//	@Description	Delete an asset task consumable
//	@Param			groupTitle		path	string	true	"Group Title"
//	@Param			assetTitle		path	string	true	"Asset Title"
//	@Param			taskId			path	string	true	"Asset Task ID"
//	@Param			consumableId	path	string	true	"Consumable ID"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/consumables/{consumableId} [delete]
func (h *Api) deleteTaskConsumable(c *gin.Context) {
	err := h.app.DeleteTaskConsumable(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTaskConsumable godoc
//
//	@Summary		Get an asset task consumable
//	@Description	Get an asset task consumable
//	@Param			groupTitle		path	string	true	"Group Title"
//	@Param			assetTitle		path	string	true	"Asset Title"
//	@Param			taskId			path	string	true	"Asset Task ID"
//	@Param			consumableId	path	string	true	"Consumable ID"
//	@Produce		json
//	@Success		200	{object}	tp.TaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/consumables/{consumableId} [get]
func (h *Api) getTaskConsumable(c *gin.Context) {
	taskConsumable, err := h.app.GetTaskConsumable(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskConsumable))
}

// ListTaskConsumables godoc
//
//	@Summary		List asset task consumables
//	@Description	List all asset task consumables
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			taskId		path	string	true	"Asset Task ID"
//	@Produce		json
//	@Success		200	{object}	[]tp.TaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/consumables [get]
func (h *Api) listTaskConsumables(c *gin.Context) {
	taskConsumables, err := h.app.ListTaskConsumables(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskConsumables))
}

// UpdateTaskConsumableBody godoc
//
//	@Summary		Update an asset task consumable with json body
//	@Description	Update an asset task consumable with json body
//	@Accept			json
//	@Param			taskConsumable	body	tp.TaskConsumable	true	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.TaskConsumable
//	@Router			/task-consumables/{consumableId} [put]
func (h *Api) updateTaskConsumableBody(c *gin.Context) {
	var taskConsumable tp.TaskConsumable
	if err := c.BindJSON(&taskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskConsumable, err := h.app.UpdateTaskConsumable(taskConsumable)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskConsumable))
}

// UpdateTaskConsumablePath godoc
//
//	@Summary		Update an asset task consumable with path parameters
//	@Description	Update an asset task consumable with path parameters
//	@Param			groupTitle		path	string						true	"Group Title"
//	@Param			assetTitle		path	string						true	"Asset Title"
//	@Param			taskId			path	string						true	"Asset Task ID"
//	@Param			consumableId	path	string						true	"Consumable ID"
//	@Param			taskConsumable	body	tp.TaskConsumableForPath	true	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.TaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/consumables/{consumableId} [put]
func (h *Api) updateTaskConsumablePath(c *gin.Context) {
	var taskConsumable tp.TaskConsumable
	if err := c.BindJSON(&taskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskConsumable, err := h.app.UpdateTaskConsumableWithValidation(
		c.Param(groupTitle),
		c.Param(assetTitle),
		c.Param(taskId),
		c.Param(consumableId),
		taskConsumable.QuantityNote,
	)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskConsumable))
}
