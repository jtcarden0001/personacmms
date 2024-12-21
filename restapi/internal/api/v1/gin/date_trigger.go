package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseDateTriggerRoute = fmt.Sprintf("%s/date-triggers", indTaskRoute)
var dateTriggerId = "DateTriggerId"
var indDateTriggerRoute = fmt.Sprintf("%s/:%s", baseDateTriggerRoute, dateTriggerId)

func (h *Api) registerDateTriggerRoutes() {
	h.router.GET(baseDateTriggerRoute, h.listDateTriggers)
	h.router.GET(indDateTriggerRoute, h.getDateTrigger)
	h.router.POST(baseDateTriggerRoute, h.createDateTrigger)
	h.router.PUT(indDateTriggerRoute, h.updateDateTrigger)
	h.router.DELETE(indDateTriggerRoute, h.deleteDateTrigger)
}

// CreateDateTrigger godoc
//
//	@Summary		Create a date trigger
//	@Description	Create a date trigger
//	@Accept			json
//	@Param			groupTitle	path	string			true	"Group Title"
//	@Param			assetTitle	path	string			true	"Asset Id"
//	@Param			taskId		path	string			true	"Asset Task Id"
//	@Param			dateTrigger	body	tp.DateTrigger	true	"Date Trigger object"
//	@Produce		json
//	@Success		201	{object}	tp.DateTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/date-triggers [post]
func (h *Api) createDateTrigger(c *gin.Context) {
	var dateTrigger tp.DateTrigger
	if err := c.BindJSON(&dateTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateTrigger, err := h.app.CreateDateTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), dateTrigger)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, dateTrigger))
}

// DeleteDateTrigger godoc
//
//	@Summary		Delete a date trigger
//	@Description	Delete a date trigger
//	@Param			groupTitle		path	string	true	"Group Title"
//	@Param			assetTitle		path	string	true	"Asset Title"
//	@Param			taskId			path	string	true	"Asset Task Id"
//	@Param			dateTriggerId	path	string	true	"Date Trigger Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/date-triggers/{dateTriggerId} [delete]
func (h *Api) deleteDateTrigger(c *gin.Context) {
	err := h.app.DeleteDateTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(dateTriggerId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetDateTrigger godoc
//
//	@Summary		Get a date trigger
//	@Description	Get a date trigger
//	@Param			groupTitle		path	string	true	"Group Title"
//	@Param			assetTitle		path	string	true	"Asset Title"
//	@Param			taskId			path	string	true	"Asset Task Id"
//	@Param			dateTriggerId	path	string	true	"Date Trigger Id"
//	@Produce		json
//	@Success		200	{object}	tp.DateTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/date-triggers/{dateTriggerId} [get]
func (h *Api) getDateTrigger(c *gin.Context) {
	dateTrigger, err := h.app.GetDateTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(dateTriggerId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTrigger))
}

// ListDateTriggers godoc
//
//	@Summary		List date triggers
//	@Description	List date triggers
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			taskId		path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.DateTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/date-triggers [get]
func (h *Api) listDateTriggers(c *gin.Context) {
	dateTriggers, err := h.app.ListDateTriggers(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTriggers))
}

// UpdateDateTrigger godoc
//
//	@Summary		Update a date trigger
//	@Description	Update a date trigger
//	@Accept			json
//	@Param			groupTitle		path	string			true	"Group Title"
//	@Param			assetTitle		path	string			true	"Asset Title"
//	@Param			taskId			path	string			true	"Asset Task Id"
//	@Param			dateTriggerId	path	string			true	"Date Trigger Id"
//	@Param			dateTrigger		body	tp.DateTrigger	true	"Date Trigger object"
//	@Produce		json
//	@Success		200	{object}	tp.DateTrigger
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{taskId}/date-triggers/{dateTriggerId} [put]
func (h *Api) updateDateTrigger(c *gin.Context) {
	var dateTrigger tp.DateTrigger
	if err := c.BindJSON(&dateTrigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateTrigger, err := h.app.UpdateDateTrigger(c.Param(groupTitle), c.Param(assetTitle), c.Param(taskId), c.Param(dateTriggerId), dateTrigger)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, dateTrigger))
}
