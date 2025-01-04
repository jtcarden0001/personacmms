package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
)

// - POST /assets (JSON) done
// - GET  /assets/{assetId} done
// - GET  /assets done
// - PUT  /assets/{assetId} (JSON) done
// - DEL  /assets/{assetId} done
//
// - GET  /categories/{categoryId}/assets done
// - GET  /categories/{categoryId}/groups/{groupId}/assets done
// - PUT  /categories/{categoryId}/assets/{assetId} (JSON) done
// - DEL  /categories/{categoryId}/assets/{assetId} done
//
// - GET  /groups/{groupId}/assets done
// - PUT  /groups/{groupId}/assets/{assetId} (JSON) done
// - DEL  /groups/{groupId}/assets/{assetId} done

var assetGp = "assets"
var assetId = "assetId"
var assetResource = fmt.Sprintf("%s/:%s", assetGp, assetId)

var baseAssetRoute = fmt.Sprintf("%s/%s", routePrefix, assetGp)
var indAssetRoute = fmt.Sprintf("%s/%s", routePrefix, assetResource)

func (h *Api) registerAssetRoutes() {
	h.router.POST(baseAssetRoute, h.createAsset)

	h.router.DELETE(indAssetRoute, h.deleteAsset)
	h.router.DELETE(fmt.Sprintf("%s/%s", indCategoryRoute, assetResource), h.disassociateAssetWithCategory)
	h.router.DELETE(fmt.Sprintf("%s/%s", indGroupRoute, assetResource), h.disassociateAssetWithGroup)

	h.router.GET(indAssetRoute, h.getAsset)
	h.router.GET(baseAssetRoute, h.listAssets)
	h.router.GET(fmt.Sprintf("%s/%s", indGroupRoute, assetGp), h.listAssetsByGroup)
	h.router.GET(fmt.Sprintf("%s/%s", indCategoryRoute, assetResource), h.listAssetsByCategory)
	h.router.GET(fmt.Sprintf("%s/%s/%s", indCategoryRoute, groupResource, assetGp), h.listAssetsByCategoryAndGroup)

	h.router.PUT(fmt.Sprintf("%s/%s", indCategoryRoute, assetResource), h.associateAssetWithCategory)
	h.router.PUT(fmt.Sprintf("%s/%s", indGroupRoute, assetResource), h.associateAssetWithGroup)
	h.router.PUT(indAssetRoute, h.updateAsset)
}

// associateAssetWithCategory godoc
//
//	@Summary		Associate an asset with a category
//	@Description	Associate an asset with a category
//	@Tags			assets
//	@Param			categoryId	path	string	true	"Category Id"
//	@Param			assetId		path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId}/assets/{assetId} [put]
func (h *Api) associateAssetWithCategory(c *gin.Context) {
	asset, err := h.app.AssociateAssetWithCategory(c.Param(categoryId), c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// associateAssetWithGroup godoc
//
//	@Summary		Associate an asset with a group
//	@Description	Associate an asset with a group
//	@Tags			assets
//	@Param			groupId	path	string	true	"Group Id"
//	@Param			assetId	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups/{groupId}/assets/{assetId} [put]
func (h *Api) associateAssetWithGroup(c *gin.Context) {
	asset, err := h.app.AssociateAssetWithGroup(c.Param(groupId), c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// createAsset godoc
//
//	@Summary		Create an asset
//	@Description	Create an asset
//	@Tags			assets
//	@Accept			json
//	@Param			asset	body	apitp.AssetRequest	true	"Asset object"
//	@Produce		json
//	@Success		201	{object}	apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets [post]
func (h *Api) createAsset(c *gin.Context) {
	var a apitp.AssetRequest
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	ar, err := h.app.CreateAsset(a)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, ar))
}

// deleteAsset godoc
//
//	@Summary		Delete an asset
//	@Description	Delete an asset
//	@Tags			assets
//	@Param			assetId	path	string	true	"Asset Title"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId} [delete]
func (h *Api) deleteAsset(c *gin.Context) {
	err := h.app.DeleteAsset(c.Param(assetId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// disassociateAssetWithCategory godoc
//
//	@Summary		Disassociate an asset with a category
//	@Description	Disassociate an asset with a category
//	@Tags			assets
//	@Param			categoryId	path	string	true	"Category Id"
//	@Param			assetId		path	string	true	"Asset Id"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId}/assets/{assetId} [delete]
func (h *Api) disassociateAssetWithCategory(c *gin.Context) {
	err := h.app.DisassociateAssetWithCategory(c.Param(categoryId), c.Param(assetId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// disassociateAssetWithGroup godoc
//
//	@Summary		Disassociate an asset with a group
//	@Description	Disassociate an asset with a group
//	@Tags			assets
//	@Param			groupId	path	string	true	"Group Id"
//	@Param			assetId	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups/{groupId}/assets/{assetId} [delete]
func (h *Api) disassociateAssetWithGroup(c *gin.Context) {
	err := h.app.DisassociateAssetWithGroup(c.Param(groupId), c.Param(assetId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// getAsset godoc
//
//	@Summary		Get an asset
//	@Description	Get an asset
//	@Tags			assets
//	@Param			assetId	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId} [get]
func (h *Api) getAsset(c *gin.Context) {
	asset, err := h.app.GetAsset(c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// listAssets godoc
//
//	@Summary		List assets
//	@Description	List all assets
//	@Tags			assets
//	@Produce		json
//	@Success		200	{object}	[]apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets [get]
func (h *Api) listAssets(c *gin.Context) {
	asset, err := h.app.ListAssets()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

// listAssetsByCategory godoc
//
//	@Summary		List assets by category
//	@Description	List all assets by category
//	@Tags			assets
//	@Param			categoryId	path	string	true	"Category Id"
//	@Produce		json
//	@Success		200	{object}	[]apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId}/assets [get]
func (h *Api) listAssetsByCategory(c *gin.Context) {
	assets, err := h.app.ListAssetsByCategory(c.Param(categoryId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assets))
}

// listAssetsByCategoryAndGroup godoc
//
//	@Summary		List assets by category and group
//	@Description	List all assets by category and group
//	@Tags			assets
//	@Param			categoryId	path	string	true	"Category Id"
//	@Param			groupId		path	string	true	"Group Id"
//	@Produce		json
//	@Success		200	{object}	[]apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId}/groups/{groupId}/assets [get]
func (h *Api) listAssetsByCategoryAndGroup(c *gin.Context) {
	assets, err := h.app.ListAssetsByCategoryAndGroup(c.Param(categoryId), c.Param(groupId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assets))
}

// listAssetsByGroup godoc
//
//	@Summary		List assets by group
//	@Description	List all assets by group
//	@Tags			assets
//	@Param			groupId	path	string	true	"Group Id"
//	@Produce		json
//	@Success		200	{object}	[]apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/groups/{groupId}/assets [get]
func (h *Api) listAssetsByGroup(c *gin.Context) {
	assets, err := h.app.ListAssetsByGroup(c.Param(groupId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assets))
}

// updateAsset godoc
//
//	@Summary		Update an asset
//	@Description	Update an asset
//	@Tags			assets
//	@Accept			json
//	@Param			assetId	path	string				true	"Asset Id"
//	@Param			asset	body	apitp.AssetRequest	true	"Asset object"
//	@Produce		json
//	@Success		200	{object}	apitp.AssetResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId} [put]
func (h *Api) updateAsset(c *gin.Context) {
	var a apitp.AssetRequest
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	ar, err := h.app.UpdateAsset(c.Param(assetId), a)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, ar))
}
