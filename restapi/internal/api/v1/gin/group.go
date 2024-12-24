package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /groups (JSON) done
// - GET  /groups/{groupId} done
// - GET  /groups done
// - PUT  /groups/{groupdId} done
// - DEL  /groups/{groupId} done
//
// - GET  /assets/{assetId}/groups

var groupId = "groupId"
var groupGp = "groups"
var groupResource = fmt.Sprintf("%s/:%s", groupGp, groupId)

var baseGroupRoute = fmt.Sprintf("%s/%s", routePrefix, groupGp)
var indGroupRoute = fmt.Sprintf("%s/%s", routePrefix, groupResource)

func (h *Api) registerGroupRoutes() {
	h.router.POST(baseGroupRoute, h.createGroup)

	h.router.DELETE(indGroupRoute, h.deleteGroup)

	h.router.GET(indGroupRoute, h.getGroup)
	h.router.GET(baseGroupRoute, h.listGroups)
	h.router.GET(fmt.Sprintf("%s/%s", indAssetRoute, groupGp), h.listGroupsByAsset)

	h.router.PUT(indGroupRoute, h.updateGroup)
}

// CreateGroup godoc
//
//	@Summary		Create a group
//	@Description	Create a group
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			group	body		tp.Group	true	"Group object"
//	@Success		201		{object}	tp.Group
//	@Failure		400		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/groups [post]
func (h *Api) createGroup(c *gin.Context) {
	var group tp.Group
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	group, err := h.app.CreateGroup(group)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, group))
}

// DeleteGroup godoc
//
//	@Summary		Delete a group
//	@Description	Delete a group
//	@Tags			groups
//	@Param			groupId	path	string	true	"Group Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups/{groupId} [delete]
func (h *Api) deleteGroup(c *gin.Context) {
	err := h.app.DeleteGroup(c.Param(groupId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetGroup godoc
//
//	@Summary		Get a group
//	@Description	Get a group
//	@Tags			groups
//	@Produce		json
//	@Param			groupId	path		string	true	"Group Id"
//	@Success		200		{object}	tp.Group
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/groups/{groupId} [get]
func (h *Api) getGroup(c *gin.Context) {
	group, err := h.app.GetGroup(c.Param(groupId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, group))
}

// ListGroups godoc
//
//	@Summary		List all groups
//	@Description	List all groups
//	@Tags			groups
//	@Produce		json
//	@Success		200	{object}	[]tp.Group
//	@Failure		400	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups [get]
func (h *Api) listGroups(c *gin.Context) {
	groups, err := h.app.ListGroups()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, groups))
}

// ListGroupsByAsset godoc
//
//	@Summary		List groups by asset
//	@Description	List groups by asset
//	@Tags			groups
//	@Produce		json
//	@Param			assetId	path		string	true	"Asset Id"
//	@Success		200		{object}	[]tp.Group
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/assets/{assetId}/groups [get]
func (h *Api) listGroupsByAsset(c *gin.Context) {
	groups, err := h.app.ListGroupsByAsset(c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, groups))
}

// UpdateGroup godoc
//
//	@Summary		Update a group
//	@Description	Update a group
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			groupId	path		string		true	"Group Id"
//	@Param			group	body		tp.Group	true	"Group object"
//	@Success		200		{object}	tp.Group
//	@Failure		400		{object}	map[string]any
//	@Failure		404		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/groups/{groupId} [put]
func (h *Api) updateGroup(c *gin.Context) {
	var group tp.Group
	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	group, err := h.app.UpdateGroup(c.Param(groupId), group)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, group))
}
