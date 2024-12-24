package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /assets/{assetId}/tasks/{taskId}/usage-triggers (JSON) done
// - GET  /assets/{assetId}/tasks/{taskId}/usage-triggers/{usagedTiggerId} done
// - GET  /assets/{assetId}/tasks/{taskId}/usage-triggers done
// - PUT  /assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTiggerId} (JSON) done
// - DEL  /assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTiggerId} done

var baseUsageTriggerRoute = fmt.Sprintf("%s/usage-triggers", indTaskRoute)
var usageTriggerId = "UsageTriggerId"
var indUsageTriggerRoute = fmt.Sprintf("%s/:%s", baseUsageTriggerRoute, usageTriggerId)

func (h *Api) registerUsageTriggerRoutes() {
	h.router.POST(baseUsageTriggerRoute, h.createUsageTrigger)

	h.router.DELETE(indUsageTriggerRoute, h.deleteUsageTrigger)

	h.router.GET(baseUsageTriggerRoute, h.listUsageTriggersByAssetAndTask)
	h.router.GET(indUsageTriggerRoute, h.getUsageTrigger)

	h.router.PUT(indUsageTriggerRoute, h.updateUsageTrigger)
}

// CreateUsageTrigger godoc
//
//	@Summary		Create a usage trigger
//	@Description	Create a usage trigger for a task
//	@Tags			usage-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string			true	"Asset Id"
//	@Param			taskId			path		string			true	"Asset Task Id"
//	@Param			usageTrigger	body		tp.UsageTrigger	true	"Usage Trigger object"
//	@Success		201				{object}	tp.UsageTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/usage-triggers [post]
func (h *Api) createUsageTrigger(c *gin.Context) {
	var usageTrigger tp.UsageTrigger
	if err := c.BindJSON(&usageTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usageTrigger, err := h.app.CreateUsageTrigger(c.Param(assetId), c.Param(taskId), usageTrigger)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, usageTrigger))
}

// DeleteUsageTrigger godoc
//
//	@Summary		Delete a usage trigger
//	@Description	Delete a usage trigger
//	@Tags			usage-triggers
//	@Param			assetId			path	string	true	"Asset Id"
//	@Param			taskId			path	string	true	"Asset Task Id"
//	@Param			usageTriggerId	path	string	true	"Usage Trigger Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} [delete]
func (h *Api) deleteUsageTrigger(c *gin.Context) {
	err := h.app.DeleteUsageTrigger(c.Param(assetId), c.Param(taskId), c.Param(usageTriggerId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetUsageTrigger godoc
//
//	@Summary		Get a usage trigger
//	@Description	Get a usage trigger
//	@Tags			usage-triggers
//	@Produce		json
//	@Param			assetId			path		string	true	"Asset Id"
//	@Param			taskId			path		string	true	"Asset Task Id"
//	@Param			usageTriggerId	path		string	true	"Usage Trigger Id"
//	@Success		200				{object}	tp.UsageTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} [get]
func (h *Api) getUsageTrigger(c *gin.Context) {
	usageTrigger, err := h.app.GetUsageTrigger(c.Param(assetId), c.Param(taskId), c.Param(usageTriggerId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageTrigger))
}

// ListUsageTriggers godoc
//
//	@Summary		List usage triggers
//	@Description	List usage triggers for a task
//	@Tags			usage-triggers
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Param			taskId	path		string	true	"Asset Task Id"
//	@Success		200		{object}	[]tp.UsageTrigger
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/usage-triggers [get]
func (h *Api) listUsageTriggersByAssetAndTask(c *gin.Context) {
	usageTriggers, err := h.app.ListUsageTriggersByAssetAndTask(c.Param(assetId), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageTriggers))
}

// UpdateUsageTrigger godoc
//
//	@Summary		Update a usage trigger
//	@Description	Update a usage trigger
//	@Tags			usage-triggers
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string			true	"Asset Id"
//	@Param			taskId			path		string			true	"Asset Task Id"
//	@Param			usageTriggerId	path		string			true	"Usage Trigger Id"
//	@Param			usageTrigger	body		tp.UsageTrigger	true	"Usage Trigger object"
//	@Success		200				{object}	tp.UsageTrigger
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId} [put]
func (h *Api) updateUsageTrigger(c *gin.Context) {
	var usageTrigger tp.UsageTrigger
	if err := c.BindJSON(&usageTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usageTrigger, err := h.app.UpdateUsageTrigger(c.Param(assetId), c.Param(taskId), c.Param(usageTriggerId), usageTrigger)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageTrigger))
}
