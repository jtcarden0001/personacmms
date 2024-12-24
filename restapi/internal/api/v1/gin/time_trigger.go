package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /assets/{assetId}/tasks/{taskId}/time-triggers (JSON) done
// - GET  /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTiggerId} done
// - GET  /assets/{assetId}/tasks/{taskId}/time-triggers done
// - PUT  /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTiggerId} (JSON) done
// - DEL  /assets/{assetId}/tasks/{taskId}/time-triggers/{timeTiggerId} done

var timeTriggerId = "TimeTriggerId"
var timeTriggerGp = "time-triggers"
var timeTriggerResource = fmt.Sprintf("%s/:%s", timeTriggerGp, timeTriggerId)

var baseTimeTriggerRoute = fmt.Sprintf("%s/%s", indTaskRoute, timeTriggerGp)
var indTimeTriggerRoute = fmt.Sprintf("%s/%s", indTaskRoute, timeTriggerResource)

func (h *Api) registerTimeTriggerRoutes() {
	h.router.POST(baseTimeTriggerRoute, h.createTimeTrigger)

	h.router.DELETE(indTimeTriggerRoute, h.deleteTimeTrigger)

	h.router.GET(indTimeTriggerRoute, h.getTimeTrigger)
	h.router.GET(baseTimeTriggerRoute, h.listTimeTriggers)

	h.router.PUT(indTimeTriggerRoute, h.updateTimeTrigger)
}

// CreateTimeTrigger godoc
//
//	@Summary		Create a time trigger
//	@Description	Create a time trigger
//	@Tags			time-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId		path		string			true	"Asset Id"
//	@Param			taskId		path		string			true	"Asset Task Id"
//	@Param			timeTrigger	body		tp.TimeTrigger	true	"Time Trigger object"
//	@Success		201			{object}	tp.TimeTrigger
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/time-triggers [post]
func (h *Api) createTimeTrigger(c *gin.Context) {
	var timeTrigger tp.TimeTrigger
	if err := c.BindJSON(&timeTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	timeTrigger, err := h.app.CreateTimeTrigger(c.Param(assetId), c.Param(taskId), timeTrigger)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, timeTrigger))
}

// DeleteTimeTrigger godoc
//
//	@Summary		Delete a time trigger
//	@Description	Delete a time trigger
//	@Tags			time-triggers
//	@Param			assetId			path	string	true	"Asset Id"
//	@Param			taskId			path	string	true	"Asset Task Id"
//	@Param			timeTriggerId	path	string	true	"Time Trigger Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} [delete]
func (h *Api) deleteTimeTrigger(c *gin.Context) {
	err := h.app.DeleteTimeTrigger(c.Param(assetId), c.Param(taskId), c.Param(timeTriggerId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTimeTrigger godoc
//
//	@Summary		Get a time trigger
//	@Description	Get a time trigger
//	@Tags			time-triggers
//	@Produce		json
//	@Param			assetId			path		string	true	"Asset Id"
//	@Param			taskId			path		string	true	"Asset Task Id"
//	@Param			timeTriggerId	path		string	true	"Time Trigger Id"
//	@Success		200				{object}	tp.TimeTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} [get]
func (h *Api) getTimeTrigger(c *gin.Context) {
	timeTrigger, err := h.app.GetTimeTrigger(c.Param(assetId), c.Param(taskId), c.Param(timeTriggerId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTrigger))
}

// ListTimeTriggers godoc
//
//	@Summary		List time triggers
//	@Description	List all time triggers
//	@Tags			time-triggers
//	@Produce		json
//	@Param			assetId		path		string	true	"Asset Id"
//	@Param			taskId		path		string	true	"Asset Task Id"
//	@Success		200			{object}	[]tp.TimeTrigger
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/time-triggers [get]
func (h *Api) listTimeTriggers(c *gin.Context) {
	timeTriggers, err := h.app.ListTimeTriggers(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTriggers))
}

// UpdateTimeTrigger godoc
//
//	@Summary		Update a time trigger
//	@Description	Update a time trigger
//	@Tags			time-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string			true	"Asset Id"
//	@Param			taskId			path		string			true	"Task Id"
//	@Param			timeTriggerId	path		string			true	"Time Trigger Id"
//	@Param			timeTrigger		body		tp.TimeTrigger	true	"Time Trigger object"
//	@Success		200				{object}	tp.TimeTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/time-triggers/{timeTriggerId} [put]
func (h *Api) updateTimeTrigger(c *gin.Context) {
	var timeTrigger tp.TimeTrigger
	if err := c.BindJSON(&timeTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	timeTrigger, err := h.app.UpdateTimeTrigger(c.Param(assetId), c.Param(taskId), c.Param(timeTriggerId), timeTrigger)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTrigger))
}
