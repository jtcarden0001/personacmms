package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerGroupRoutes() {
	baseRoute := fmt.Sprintf("%s/groups", routePrefix)
	individualRoute := fmt.Sprintf("%s/:groupTitle", baseRoute)

	h.router.POST(baseRoute, h.createGroup)
	h.router.DELETE(individualRoute, h.deleteGroup)
	h.router.GET(baseRoute, h.listGroups)
	// h.router.GET(individualRoute, h.getGroup)
	// h.router.PUT(individualRoute, h.updateGroup)
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
	if err != nil {
		processAppError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, group) // switch to .JSON() for performance
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
	title := c.Param("groupTitle")
	err := h.app.DeleteGroup(title)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{}) // switch to .JSON() for performance
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
	if err != nil {
		processAppError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, groups)
}
