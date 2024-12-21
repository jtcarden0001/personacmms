package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /groups (JSON)
// - GET  /groups/{groupId}
// - GET  /groups
// - PUT  /groups/{groupdId}
// - DEL  /groups/{groupId}
// - GET  /assets/{assetId}/groups

var baseGroupRoute = fmt.Sprintf("%s/groups", routePrefix)
var groupTitle = "groupTitle"
var indGroupRoute = fmt.Sprintf("%s/:%s", baseGroupRoute, groupTitle)

func (h *Api) registerGroupRoutes() {

	h.router.POST(baseGroupRoute, h.createGroup)
	h.router.DELETE(indGroupRoute, h.deleteGroup)
	h.router.GET(baseGroupRoute, h.listGroups)
	h.router.GET(indGroupRoute, h.getGroup)
	h.router.PUT(indGroupRoute, h.updateGroup)
}

// CreateGroup godoc
//
//	@Summary		Create an asset group
//	@Description	Create a group
//	@Accept			json
//	@Param			group	body	tp.Group	true	"Group object"
//	@Produce		json
//	@Success		201	{object}	tp.Group
//	@Router			/groups [post]
func (h *Api) createGroup(c *gin.Context) {
	var group tp.Group
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := h.app.CreateGroup(group)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, group))
}

// DeleteGroup godoc
//
//	@Summary		Delete an asset group
//	@Description	Delete a group
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle} [delete]
func (h *Api) deleteGroup(c *gin.Context) {
	err := h.app.DeleteGroup(c.Param(groupTitle))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetGroup godoc
//
//	@Summary		Get an asset group
//	@Description	Get a group
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Produce		json
//	@Success		200	{object}	tp.Group
//	@Router			/groups/{groupTitle} [get]
func (h *Api) getGroup(c *gin.Context) {
	group, err := h.app.GetGroup(c.Param(groupTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, group))
}

// ListGroups godoc
//
//	@Summary		List asset groups
//	@Description	List all groups
//	@Produce		json
//	@Success		200	{object}	[]tp.Group
//	@Router			/groups [get]
func (h *Api) listGroups(c *gin.Context) {
	groups, err := h.app.ListGroups()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, groups))
}

// UpdateGroup godoc
//
//	@Summary		Update an asset group
//	@Description	Update a group
//	@Accept			json
//	@Param			groupTitle	path	string		true	"Group Title"
//	@Param			group		body	tp.Group	true	"Group object"
//	@Produce		json
//	@Success		200	{object}	tp.Group
//	@Router			/groups/{groupTitle} [put]
func (h *Api) updateGroup(c *gin.Context) {
	var group tp.Group
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := h.app.UpdateGroup(c.Param(groupTitle), group)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, group))
}
