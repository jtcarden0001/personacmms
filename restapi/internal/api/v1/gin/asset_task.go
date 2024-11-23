package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseAssetTaskRoute = fmt.Sprintf("%s/tasks", indAssetRoute)
var assetTaskId = "AssetTaskId"
var indAssetTaskRoute = fmt.Sprintf("%s/:%s", baseAssetTaskRoute, assetTaskId)

// general route that treats all asset tasks the same, can CRUD any type through this route
func (h *Api) registerAssetTaskRoutes() {
	h.router.POST(baseAssetTaskRoute, h.createAssetTask)
	h.router.DELETE(indAssetTaskRoute, h.deleteAssetTask)
	h.router.GET(baseAssetTaskRoute, h.listAssetTasks)
	h.router.GET(indAssetTaskRoute, h.getAssetTask)
	h.router.PUT(indAssetTaskRoute, h.updateAssetTask)
}

// CreateAssetTask godoc
//
//	@Summary		Create an asset task
//	@Description	Create an asset task
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTask	body	tp.AssetTask	true	"Asset Task object"
//	@Produce		json
//	@Success		201	{object}	tp.AssetTask
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks [post]
func (h *Api) createAssetTask(c *gin.Context) {
	var assetTask tp.AssetTask
	if err := c.BindJSON(&assetTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTask, err := h.app.CreateAssetTask(c.Param(groupTitle), c.Param(assetTitle), assetTask)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTask))
}

// DeleteAssetTask godoc
//
//	@Summary		Delete an asset task
//	@Description	Delete an asset task
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId} [delete]
func (h *Api) deleteAssetTask(c *gin.Context) {
	err := h.app.DeleteAssetTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListAssetTasks godoc
//
//	@Summary		List asset tasks
//	@Description	List all asset tasks
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	[]tp.AssetTask
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks [get]
func (h *Api) listAssetTasks(c *gin.Context) {
	assetTasks, err := h.app.ListAssetTasks(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTasks))
}

// GetAssetTask godoc
//
//	@Summary		Get an asset task
//	@Description	Get an asset task
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Produce		json
//	@Success		200	{object}	tp.AssetTask
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId} [get]
func (h *Api) getAssetTask(c *gin.Context) {
	assetTask, err := h.app.GetAssetTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTask))
}

// UpdateAssetTask godoc
//
//	@Summary		Update an asset task
//	@Description	Update an asset task
//	@Accept			json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Id"
//	@Param			assetTaskId	path	string	true	"Asset Task Id"
//	@Param			assetTask	body	tp.AssetTask	true	"Asset Task object"
//	@Produce		json
//	@Success		200	{object}	tp.AssetTask
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/tasks/{assetTaskId} [put]
func (h *Api) updateAssetTask(c *gin.Context) {
	var assetTask tp.AssetTask
	if err := c.BindJSON(&assetTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTask, err := h.app.UpdateAssetTask(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), assetTask)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTask))
}
