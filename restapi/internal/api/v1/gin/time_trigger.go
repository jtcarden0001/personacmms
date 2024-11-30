package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseTimeTriggerRoute = fmt.Sprintf("%s/time-triggers", indTaskRoute)
var timeTriggerId = "TimeTriggerId"
var indTimeTriggerRoute = fmt.Sprintf("%s/:%s", baseTimeTriggerRoute, timeTriggerId)

func (h *Api) registerTimeTriggerRoutes() {
	h.router.GET(baseTimeTriggerRoute, h.listTimeTriggers)
	h.router.GET(indTimeTriggerRoute, h.getTimeTrigger)
	h.router.POST(baseTimeTriggerRoute, h.createTimeTrigger)
	h.router.PUT(indTimeTriggerRoute, h.updateTimeTrigger)
	h.router.DELETE(indTimeTriggerRoute, h.deleteTimeTrigger)
}

// CreateTimeTrigger godoc
//
//	@Summary		Create a time trigger
//	@Description	Create a time trigger
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Param			timeTrigger	body	tp.TimeTrigger	true	"Time Trigger object"
//	@Produce		json
//	@Success		201	{object}	tp.TimeTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/time-triggers [post]
func (h *Api) createTimeTrigger(c *gin.Context) {
	var timeTrigger tp.TimeTrigger
	if err := c.BindJSON(&timeTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeTrigger, err := h.app.CreateTimeTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), timeTrigger)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, timeTrigger))
}

// DeleteTimeTrigger godoc
//
//	@Summary		Delete a time trigger
//	@Description	Delete a time trigger
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Param			timeTriggerId	path	string	true	"Time Trigger Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/time-triggers/{timeTriggerId} [delete]
func (h *Api) deleteTimeTrigger(c *gin.Context) {
	err := h.app.DeleteTimeTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(timeTriggerId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTimeTrigger godoc
//
//	@Summary		Get a time trigger
//	@Description	Get a time trigger
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Param			timeTriggerId	path	string	true	"Time Trigger Id"
//	@Produce		json
//	@Success		200	{object}	tp.TimeTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/time-triggers/{timeTriggerId} [get]
func (h *Api) getTimeTrigger(c *gin.Context) {
	timeTrigger, err := h.app.GetTimeTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(timeTriggerId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTrigger))
}

// ListTimeTriggers godoc
//
//	@Summary		List time triggers
//	@Description	List all time triggers
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.TimeTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/time-triggers [get]
func (h *Api) listTimeTriggers(c *gin.Context) {
	timeTriggers, err := h.app.ListTimeTriggers(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTriggers))
}

// UpdateTimeTrigger godoc
//
//	@Summary		Update a time trigger
//	@Description	Update a time trigger
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			taskId	path	string	true	"Asset Task Id"
//	@Param			timeTriggerId	path	string	true	"Time Trigger Id"
//	@Param			timeTrigger	body	tp.TimeTrigger	true	"Time Trigger object"
//	@Produce		json
//	@Success		200	{object}	tp.TimeTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/time-triggers/{timeTriggerId} [put]
func (h *Api) updateTimeTrigger(c *gin.Context) {
	var timeTrigger tp.TimeTrigger
	if err := c.BindJSON(&timeTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeTrigger, err := h.app.UpdateTimeTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(timeTriggerId), timeTrigger)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeTrigger))
}
