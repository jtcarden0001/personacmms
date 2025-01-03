package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /assets/{assetId}/tasks/{taskId}/date-triggers (JSON) done
// - GET  /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTiggerId} done
// - GET  /assets/{assetId}/tasks/{taskId}/date-triggers done
// - PUT  /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTiggerId} (JSON) done
// - DEL  /assets/{assetId}/tasks/{taskId}/date-triggers/{dateTiggerId} done

var dateTriggerId = "dateTriggerId"
var dateTriggerGp = "date-triggers"
var dateTriggerResource = fmt.Sprintf("%s/:%s", dateTriggerGp, dateTriggerId)

var baseDateTriggerRoute = fmt.Sprintf("%s/%s", indTaskRoute, dateTriggerGp)
var indDateTriggerRoute = fmt.Sprintf("%s/%s", indTaskRoute, dateTriggerResource)

func (h *Api) registerDateTriggerRoutes() {
	h.router.POST(baseDateTriggerRoute, h.createDateTrigger)

	h.router.DELETE(indDateTriggerRoute, h.deleteDateTrigger)

	h.router.GET(baseDateTriggerRoute, h.listDateTriggersByAssetAndTask)
	h.router.GET(indDateTriggerRoute, h.getDateTrigger)

	h.router.PUT(indDateTriggerRoute, h.updateDateTrigger)
}

// CreateDateTrigger godoc
//
//	@Summary		Create a date trigger
//	@Description	Create a date trigger
//	@Tags			date-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId		path		string			true	"Asset Id"
//	@Param			taskId		path		string			true	"Task Id"
//	@Param			dateTrigger	body		tp.DateTrigger	true	"Date Trigger object"
//	@Success		201			{object}	tp.DateTrigger
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/date-triggers [post]
func (h *Api) createDateTrigger(c *gin.Context) {
	var dateTrigger tp.DateTrigger
	if err := c.BindJSON(&dateTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	dateTrigger, err := h.app.CreateDateTrigger(c.Param(assetId), c.Param(taskId), dateTrigger)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, dateTrigger))
}

// DeleteDateTrigger godoc
//
//	@Summary		Delete a date trigger
//	@Description	Delete a date trigger
//	@Tags			date-triggers
//	@Param			assetId			path	string	true	"Asset Id"
//	@Param			taskId			path	string	true	"Task Id"
//	@Param			dateTriggerId	path	string	true	"Date Trigger Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} [delete]
func (h *Api) deleteDateTrigger(c *gin.Context) {
	err := h.app.DeleteDateTrigger(c.Param(assetId), c.Param(taskId), c.Param(dateTriggerId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetDateTrigger godoc
//
//	@Summary		Get a date trigger
//	@Description	Get a date trigger
//	@Tags			date-triggers
//	@Produce		json
//	@Param			assetId			path		string	true	"Asset Id"
//	@Param			taskId			path		string	true	"Task Id"
//	@Param			dateTriggerId	path		string	true	"Date Trigger Id"
//	@Success		200				{object}	tp.DateTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} [get]
func (h *Api) getDateTrigger(c *gin.Context) {
	dateTrigger, err := h.app.GetDateTrigger(c.Param(assetId), c.Param(taskId), c.Param(dateTriggerId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTrigger))
}

// ListDateTriggers godoc
//
//	@Summary		List date triggers
//	@Description	List date triggers for a task
//	@Tags			date-triggers
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			taskId	path		string	true	"Task Id"
//	@Success		200		{object}	[]tp.DateTrigger
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/date-triggers [get]
func (h *Api) listDateTriggersByAssetAndTask(c *gin.Context) {
	dateTriggers, err := h.app.ListDateTriggersByAssetAndTask(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTriggers))
}

// UpdateDateTrigger godoc
//
//	@Summary		Update a date trigger
//	@Description	Update a date trigger
//	@Tags			date-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string			true	"Asset Id"
//	@Param			taskId			path		string			true	"Task Id"
//	@Param			dateTriggerId	path		string			true	"Date Trigger Id"
//	@Param			dateTrigger		body		tp.DateTrigger	true	"Date Trigger object"
//	@Success		200				{object}	tp.DateTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/date-triggers/{dateTriggerId} [put]
func (h *Api) updateDateTrigger(c *gin.Context) {
	var dateTrigger tp.DateTrigger
	if err := c.BindJSON(&dateTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	dateTrigger, err := h.app.UpdateDateTrigger(c.Param(assetId), c.Param(taskId), c.Param(dateTriggerId), dateTrigger)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTrigger))
}
