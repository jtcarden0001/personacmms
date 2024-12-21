package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTaskTemplateRoutes() {
	baseRoute := fmt.Sprintf("%s/task-templates", routePrefix)
	individualRoute := fmt.Sprintf("%s/:taskTemplateTitle", baseRoute)

	h.router.POST(baseRoute, h.createTaskTemplate)
	h.router.DELETE(individualRoute, h.deleteTaskTemplate)
	h.router.GET(baseRoute, h.listTaskTemplates)
	h.router.GET(individualRoute, h.getTaskTemplate)
	h.router.PUT(individualRoute, h.updateTaskTemplate)
}

// CreateTaskTemplate godoc
//
//	@Summary		Create a task template
//	@Description	Create a task template
//	@Accept			json
//	@Param			taskTemplate	body	tp.TaskTemplate	true	"TaskTemplate object"
//	@Produce		json
//	@Success		201	{object}	tp.TaskTemplate
//	@Router			/task-templates [post]
func (h *Api) createTaskTemplate(c *gin.Context) {
	var taskTemplate tp.TaskTemplate
	if err := c.BindJSON(&taskTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskTemplate, err := h.app.CreateTaskTemplate(taskTemplate)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskTemplate))
}

// DeleteTaskTemplate godoc
//
//	@Summary		Delete a task template
//	@Description	Delete a task template
//	@Param			taskTemplateTitle	path	string	true	"TaskTemplate Title"
//	@Success		204
//	@Failure		404
//	@Router			/task-templates/{taskTemplateTitle} [delete]
func (h *Api) deleteTaskTemplate(c *gin.Context) {
	err := h.app.DeleteTaskTemplate(c.Param("taskTemplateTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTaskTemplates godoc
//
//	@Summary		List task templates
//	@Description	List all task templates
//	@Produce		json
//	@Success		200	{object}	[]tp.TaskTemplate
//	@Router			/task-templates [get]
func (h *Api) listTaskTemplates(c *gin.Context) {
	taskTemplates, err := h.app.ListTaskTemplates()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTemplates))
}

// GetTaskTemplate godoc
//
//	@Summary		Get a task template
//	@Description	Get a task template
//	@Param			taskTemplateTitle	path	string	true	"TaskTemplate Title"
//	@Produce		json
//	@Success		200	{object}	tp.TaskTemplate
//	@Router			/task-templates/{taskTemplateTitle} [get]
func (h *Api) getTaskTemplate(c *gin.Context) {
	taskTemplate, err := h.app.GetTaskTemplate(c.Param("taskTemplateTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTemplate))
}

// UpdateTaskTemplate godoc
//
//	@Summary		Update a task template
//	@Description	Update a task template
//	@Accept			json
//	@Param			taskTemplateTitle	path	string			true	"TaskTemplate Title"
//	@Param			taskTemplate		body	tp.TaskTemplate	true	"TaskTemplate object"
//	@Produce		json
//	@Success		204
//	@Router			/task-templates/{taskTemplateTitle} [put]
func (h *Api) updateTaskTemplate(c *gin.Context) {
	var taskTemplate tp.TaskTemplate
	if err := c.BindJSON(&taskTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskTemplate, err := h.app.UpdateTaskTemplate(c.Param("taskTemplateTitle"), taskTemplate)
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}
