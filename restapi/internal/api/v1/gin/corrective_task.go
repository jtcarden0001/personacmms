package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerCorrectiveTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/correctiveTasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:correctiveTaskTitle", baseRoute)

	h.router.POST(baseRoute, h.createCorrectiveTask)
	h.router.DELETE(individualRoute, h.deleteCorrectiveTask)
	h.router.GET(baseRoute, h.listCorrectiveTasks)
	h.router.GET(individualRoute, h.getCorrectiveTask)
	h.router.PUT(individualRoute, h.updateCorrectiveTask)
}

// CreateCorrectiveTask godoc
//
//	@Summary		Create a correctiveTask
//	@Description	Create a correctiveTask
//	@Accept			json
//	@Param			correctiveTask	body	tp.CorrectiveTask	true	"CorrectiveTask object"
//	@Produce		json
//	@Success		201	{object}	tp.CorrectiveTask
//	@Router			/correctiveTasks [post]
func (h *Api) createCorrectiveTask(c *gin.Context) {
	var correctiveTask tp.CorrectiveTask
	if err := c.BindJSON(&correctiveTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	correctiveTask, err := h.app.CreateCorrectiveTask(correctiveTask)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, correctiveTask))
}

// DeleteCorrectiveTask godoc
//
//	@Summary		Delete a correctiveTask
//	@Description	Delete a correctiveTask
//	@Param			correctiveTaskTitle	path	string	true	"CorrectiveTask Title"
//	@Success		204
//	@Failure		404
//	@Router			/correctiveTasks/{correctiveTaskTitle} [delete]
func (h *Api) deleteCorrectiveTask(c *gin.Context) {
	err := h.app.DeleteCorrectiveTask(c.Param("correctiveTaskTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListCorrectiveTasks godoc
//
//	@Summary		List correctiveTasks
//	@Description	List all correctiveTasks
//	@Produce		json
//	@Success		200	{object}	[]tp.CorrectiveTask
//	@Router			/correctiveTasks [get]
func (h *Api) listCorrectiveTasks(c *gin.Context) {
	correctiveTasks, err := h.app.ListCorrectiveTasks()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, correctiveTasks))
}

// GetCorrectiveTask godoc
//
//	@Summary		Get a correctiveTask
//	@Description	Get a correctiveTask
//	@Param			correctiveTaskTitle	path	string	true	"CorrectiveTask Title"
//	@Produce		json
//	@Success		200	{object}	tp.CorrectiveTask
//	@Router			/correctiveTasks/{correctiveTaskTitle} [get]
func (h *Api) getCorrectiveTask(c *gin.Context) {
	correctiveTask, err := h.app.GetCorrectiveTask(c.Param("correctiveTaskTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, correctiveTask))
}

// UpdateCorrectiveTask godoc
//
//	@Summary		Update a correctiveTask
//	@Description	Update a correctiveTask
//	@Accept			json
//	@Param			correctiveTaskTitle	path	string	true	"CorrectiveTask Title"
//	@Router			/correctiveTasks/{correctiveTaskTitle} [put]
func (h *Api) updateCorrectiveTask(c *gin.Context) {
	var correctiveTask tp.CorrectiveTask
	if err := c.BindJSON(&correctiveTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	correctiveTask, err := h.app.UpdateCorrectiveTask(c.Param("correctiveTaskTitle"), correctiveTask)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, correctiveTask))
}
