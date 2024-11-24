package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/tasks", routePrefix)
	individualRoute := fmt.Sprintf("%s/:taskTitle", baseRoute)

	h.router.POST(baseRoute, h.createTask)
	h.router.DELETE(individualRoute, h.deleteTask)
	h.router.GET(baseRoute, h.listTasks)
	h.router.GET(individualRoute, h.getTask)
	h.router.PUT(individualRoute, h.updateTask)
}

// CreateTask godoc
//
//	@Summary		Create a task
//	@Description	Create a task
//	@Accept			json
//	@Param			task	body	tp.Task	true	"Task object"
//	@Produce		json
//	@Success		201	{object}	tp.Task
//	@Router			/tasks [post]
func (h *Api) createTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.app.CreateTask(task)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, task))
}

// DeleteTask godoc
//
//	@Summary		Delete a task
//	@Description	Delete a task
//	@Param			taskTitle	path	string	true	"Task Title"
//	@Success		204
//	@Failure		404
//	@Router			/tasks/{taskTitle} [delete]
func (h *Api) deleteTask(c *gin.Context) {
	err := h.app.DeleteTask(c.Param("taskTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTasks godoc
//
//	@Summary		List tasks
//	@Description	List all tasks
//	@Produce		json
//	@Success		200	{object}	[]tp.Task
//	@Router			/tasks [get]
func (h *Api) listTasks(c *gin.Context) {
	tasks, err := h.app.ListTasks()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, tasks))
}

// GetTask godoc
//
//	@Summary		Get a task
//	@Description	Get a task
//	@Param			taskTitle	path	string	true	"Task Title"
//	@Produce		json
//	@Success		200	{object}	tp.Task
//	@Router			/tasks/{taskTitle} [get]
func (h *Api) getTask(c *gin.Context) {
	task, err := h.app.GetTask(c.Param("taskTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, task))
}

// UpdateTask godoc
//
//	@Summary		Update a task
//	@Description	Update a task
//	@Accept			json
//	@Param			taskTitle	path	string	true	"Task Title"
//	@Param			task		body	tp.Task	true	"Task object"
//	@Produce		json
//	@Success		204
//	@Router			/tasks/{taskTitle} [put]
func (h *Api) updateTask(c *gin.Context) {
	var task tp.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.app.UpdateTask(c.Param("taskTitle"), task)
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}
