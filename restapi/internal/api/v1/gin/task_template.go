package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /task-templates (JSON) done
// - GET  /task-templates/{taskTemplateId} done
// - GET  /task-templates done
// - PUT  /task-templates/{taskTemplateId} (JSON) done
// - DEL  /task-templates/{taskTemplateId} done

var taskTemplateId = "taskTemplateId"
var taskTemplateGp = "task-templates"
var taskTemplateResource = fmt.Sprintf("%s/:%s", taskTemplateGp, taskTemplateId)

var baseTaskTemplateRoute = fmt.Sprintf("%s/%s", routePrefix, taskTemplateGp)
var indTaskTemplateRoute = fmt.Sprintf("%s/%s", routePrefix, taskTemplateResource)

func (h *Api) registerTaskTemplateRoutes() {
	h.router.POST(baseTaskTemplateRoute, h.createTaskTemplate)

	h.router.DELETE(indTaskTemplateRoute, h.deleteTaskTemplate)

	h.router.GET(indTaskTemplateRoute, h.getTaskTemplate)
	h.router.GET(baseTaskTemplateRoute, h.listTaskTemplates)

	h.router.PUT(indTaskTemplateRoute, h.updateTaskTemplate)
}

// CreateTaskTemplate godoc
//
//	@Summary		Create a task template
//	@Description	Create a task template
//	@Tags			task-templates
//	@Accept			json
//	@Produce		json
//	@Param			taskTemplate	body	tp.TaskTemplate	true	"TaskTemplate object"
//	@Success		201	{object}	tp.TaskTemplate
//	@Failure		400	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/task-templates [post]
func (h *Api) createTaskTemplate(c *gin.Context) {
	var taskTemplate tp.TaskTemplate
	if err := c.BindJSON(&taskTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	taskTemplate, err := h.app.CreateTaskTemplate(taskTemplate)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, taskTemplate))
}

// DeleteTaskTemplate godoc
//
//	@Summary		Delete a task template
//	@Description	Delete a task template
//	@Tags			task-templates
//	@Param			taskTemplateId	path	string	true	"TaskTemplate Id"
//	@Success		204
//	@Failure		400 {object} map[string]any
//	@Failure		404 {object} map[string]any
//	@Failure		500 {object} map[string]any
//	@Router			/task-templates/{taskTemplateId} [delete]
func (h *Api) deleteTaskTemplate(c *gin.Context) {
	err := h.app.DeleteTaskTemplate(c.Param(taskTemplateId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetTaskTemplate godoc
//
//	@Summary		Get a task template
//	@Description	Get a task template
//	@Tags			task-templates
//	@Produce		json
//	@Param			taskTemplateId	path	string	true	"TaskTemplate Id"
//	@Success		200	{object}	tp.TaskTemplate
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/task-templates/{taskTemplateId} [get]
func (h *Api) getTaskTemplate(c *gin.Context) {
	taskTemplate, err := h.app.GetTaskTemplate(c.Param(taskTemplateId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTemplate))
}

// ListTaskTemplates godoc
//
//	@Summary		List task templates
//	@Description	List all task templates
//	@Tags			task-templates
//	@Produce		json
//	@Success		200	{object}	[]tp.TaskTemplate
//	@Failure		500	{object}	map[string]any
//	@Router			/task-templates [get]
func (h *Api) listTaskTemplates(c *gin.Context) {
	taskTemplates, err := h.app.ListTaskTemplates()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, taskTemplates))
}

// UpdateTaskTemplate godoc
//
//	@Summary		Update a task template
//	@Description	Update a task template
//	@Accept			json
//	@Param			taskTemplateId	path	string			true	"TaskTemplate Id"
//	@Param			taskTemplate		body	tp.TaskTemplate	true	"TaskTemplate object"
//	@Produce		json
//	@Success		200	{object}	tp.TaskTemplate
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/task-templates/{taskTemplateId} [put]
func (h *Api) updateTaskTemplate(c *gin.Context) {
	var taskTemplate tp.TaskTemplate
	if err := c.BindJSON(&taskTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	taskTemplate, err := h.app.UpdateTaskTemplate(c.Param(taskTemplateId), taskTemplate)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, nil))
}
