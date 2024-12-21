package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseAssetRoute = fmt.Sprintf("%s/assets", indGroupRoute)
var assetTitle = "assetTitle"
var indAssetRoute = fmt.Sprintf("%s/:%s", baseAssetRoute, assetTitle)

func (h *Api) registerAssetRoutes() {
	h.router.POST(baseAssetRoute, h.createAsset)
	h.router.DELETE(indAssetRoute, h.deleteAsset)
	h.router.GET(baseAssetRoute, h.listAssets)
	h.router.GET(indAssetRoute, h.getAsset)
	h.router.PUT(indAssetRoute, h.updateAsset)
}

// createAsset godoc
//
//	@Summary		Create an asset
//	@Description	Create an asset
//	@Tags			assets
//	@Accept			json
//	@Produce		json
//	@Param			groupTitle	path		string		true	"Group Title"
//	@Param			asset		body		tp.Asset	true	"Asset object"
//	@Success		201			{object}	tp.Asset
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/groups/{groupTitle}/assets [post]
func (h *Api) createAsset(c *gin.Context) {
	var a tp.Asset
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.app.CreateAsset(c.Param(groupTitle), a)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, a))
}

// deleteAsset godoc
//
//	@Summary		Delete an asset
//	@Description	Delete an asset
//	@Tags			assets
//	@Accept			json
//	@Produce		json
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups/{groupTitle}/assets/{assetTitle} [delete]
func (h *Api) deleteAsset(c *gin.Context) {
	err := h.app.DeleteAsset(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// listAssets godoc
//
//	@Summary		List assets
//	@Description	List all assets belonging to a group
//	@Tags			assets
//	@Produce		json
//	@Param			groupTitle	path		string	true	"Group Title"
//	@Success		200			{object}	[]tp.Asset
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/groups/{groupTitle}/assets [get]
func (h *Api) listAssets(c *gin.Context) {
	asset, err := h.app.ListAssets(c.Param(groupTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// getAsset godoc
//
//	@Summary		Get an asset
//	@Description	Get an asset
//	@Tags			assets
//	@Produce		json
//	@Param			groupTitle	path		string	true	"Group Title"
//	@Param			assetTitle	path		string	true	"Asset Title"
//	@Success		200			{object}	tp.Asset
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/groups/{groupTitle}/assets/{assetTitle} [get]
func (h *Api) getAsset(c *gin.Context) {
	asset, err := h.app.GetAsset(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// updateAsset godoc
//
//	@Summary		Update an asset
//	@Description	Update an asset
//	@Tags			assets
//	@Accept			json
//	@Produce		json
//	@Param			groupTitle	path		string		true	"Group Title"
//	@Param			assetTitle	path		string		true	"Asset Title"
//	@Param			asset		body		tp.Asset	true	"Asset object"
//	@Success		200			{object}	tp.Asset
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/groups/{groupTitle}/assets/{assetTitle} [put]
func (h *Api) updateAsset(c *gin.Context) {
	var a tp.Asset
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.app.UpdateAsset(c.Param(groupTitle), c.Param(assetTitle), a)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, a))
}
