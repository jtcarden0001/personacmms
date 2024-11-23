package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerPreventativeTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/preventativeTasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:preventativeTaskTitle", baseRoute)

	h.router.POST(baseRoute, h.createPreventativeTask)
	h.router.DELETE(individualRoute, h.deletePreventativeTask)
	h.router.GET(baseRoute, h.listPreventativeTasks)
	h.router.GET(individualRoute, h.getPreventativeTask)
	h.router.PUT(individualRoute, h.updatePreventativeTask)
}

// CreatePreventativeTask godoc
//
//	@Summary		Create a preventativeTask
//	@Description	Create a preventativeTask
//	@Accept			json
//	@Param			preventativeTask	body	tp.PreventativeTask	true	"PreventativeTask object"
//	@Produce		json
//	@Success		201	{object}	tp.PreventativeTask
//	@Router			/preventativeTasks [post]
func (h *Api) createPreventativeTask(c *gin.Context) {
	var preventativeTask tp.PreventativeTask
	if err := c.BindJSON(&preventativeTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	preventativeTask, err := h.app.CreatePreventativeTask(preventativeTask)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, preventativeTask))
}

// DeletePreventativeTask godoc
//
//	@Summary		Delete a preventativeTask
//	@Description	Delete a preventativeTask
//	@Param			preventativeTaskTitle	path	string	true	"PreventativeTask Title"
//	@Success		204
//	@Failure		404
//	@Router			/preventativeTasks/{preventativeTaskTitle} [delete]
func (h *Api) deletePreventativeTask(c *gin.Context) {
	err := h.app.DeletePreventativeTask(c.Param("preventativeTaskTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListPreventativeTasks godoc
//
//	@Summary		List preventativeTasks
//	@Description	List all preventativeTasks
//	@Produce		json
//	@Success		200	{object}	[]tp.PreventativeTask
//	@Router			/preventativeTasks [get]
func (h *Api) listPreventativeTasks(c *gin.Context) {
	preventativeTasks, err := h.app.ListPreventativeTasks()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, preventativeTasks))
}

// GetPreventativeTask godoc
//
//	@Summary		Get a preventativeTask
//	@Description	Get a preventativeTask
//	@Param			preventativeTaskTitle	path	string	true	"PreventativeTask Title"
//	@Produce		json
//	@Success		200	{object}	tp.PreventativeTask
//	@Router			/preventativeTasks/{preventativeTaskTitle} [get]
func (h *Api) getPreventativeTask(c *gin.Context) {
	preventativeTask, err := h.app.GetPreventativeTask(c.Param("preventativeTaskTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, preventativeTask))
}

// UpdatePreventativeTask godoc
//
//	@Summary		Update a preventativeTask
//	@Description	Update a preventativeTask
//	@Accept			json
//	@Param			preventativeTaskTitle	path	string	true	"PreventativeTask Title"
//	@Param			preventativeTask		body	tp.PreventativeTask	true	"PreventativeTask object"
//	@Produce		json
//	@Success		204
//	@Router			/preventativeTasks/{preventativeTaskTitle} [put]
func (h *Api) updatePreventativeTask(c *gin.Context) {
	var preventativeTask tp.PreventativeTask
	if err := c.BindJSON(&preventativeTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	preventativeTask, err := h.app.UpdatePreventativeTask(c.Param("preventativeTaskTitle"), preventativeTask)
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}
