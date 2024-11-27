package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var rootAssetTaskConsumableRoute = fmt.Sprintf("%s/asset-task-consumables", routePrefix)
var assetTaskConsumableRoute = fmt.Sprintf("%s/consumables", indAssetTaskRoute)
var consumableId = "ConsumableId"
var indAssetTaskConsumableRoute = fmt.Sprintf("%s/:%s", assetTaskConsumableRoute, consumableId)

func (h *Api) registerAssetTaskConsumableRoutes() {
	h.router.POST(rootAssetTaskConsumableRoute, h.createAssetTaskConsumableBody)
	h.router.POST(indAssetTaskConsumableRoute, h.createAssetTaskConsumablePath)
	h.router.DELETE(indAssetTaskConsumableRoute, h.deleteAssetTaskConsumable)
	h.router.GET(assetTaskConsumableRoute, h.listAssetTaskConsumables)
	h.router.GET(indAssetTaskConsumableRoute, h.getAssetTaskConsumable)
	h.router.PUT(indAssetTaskConsumableRoute, h.updateAssetTaskConsumableBody)
	h.router.PUT(indAssetTaskConsumableRoute, h.updateAssetTaskConsumablePath)
}

// CreateAssetTaskConsumableBody godoc
//
//	@Summary		Create a relationship between an asset task and a consumable with json body
//	@Description	Create a relationship between an asset task and a consumable with json body
//	@Accept			json
//	@Param 			assetTaskConsumable 	body 	tp.AssetTaskConsumable 	true 	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.AssetTaskConsumable
//	@Router			/asset-task-consumables [post]
func (h *Api) createAssetTaskConsumableBody(c *gin.Context) {
	var assetTaskConsumable tp.AssetTaskConsumable
	if err := c.BindJSON(&assetTaskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTaskConsumable, err := h.app.CreateAssetTaskConsumable(assetTaskConsumable)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskConsumable))
}

//	CreateAssetTaskConsumablePath godoc
//
// @Summary			Create a relationship between an asset task and a consumable with path parameters
// @Description		Create a relationship between an asset task and a consumable with path parameters
// @Accept			json
// @Param			groupTitle	path	string	true	"Group Title"
// @Param			assetTitle	path	string	true	"Asset Title"
// @Param			assetTaskId	path	string	true	"Asset Task ID"
// @Param			consumableId	path	string	true	"Consumable ID"
// @Param 			assetTaskConsumable 	body 	tp.AssetTaskConsumableForPath 	true 	"Asset Task Consumable object"
// @Produce			json
// @Success			201	{object}	tp.AssetTaskConsumable
// @Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/consumables/{consumableId} [post]
func (h *Api) createAssetTaskConsumablePath(c *gin.Context) {
	var assetTaskConsumable tp.AssetTaskConsumable
	if err := c.BindJSON(&assetTaskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTaskConsumable, err := h.app.CreateAssetTaskConsumableWithValidation(
		c.Param(groupTitle),
		c.Param(assetTitle),
		c.Param(assetTaskId),
		c.Param(consumableId),
		assetTaskConsumable.QuantityNote,
	)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskConsumable))
}

// DeleteAssetTaskConsumable godoc
//
//	@Summary		Delete an asset task consumable
//	@Description	Delete an asset task consumable
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task ID"
//	@Param			consumableId	path	string	true	"Consumable ID"
//	@Success		204
//	@Failure		404
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/consumables/{consumableId} [delete]
func (h *Api) deleteAssetTaskConsumable(c *gin.Context) {
	err := h.app.DeleteAssetTaskConsumable(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListAssetTaskConsumables godoc
//
//	@Summary		List asset task consumables
//	@Description	List all asset task consumables
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task ID"
//	@Produce		json
//	@Success		200	{object}	[]tp.AssetTaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/consumables [get]
func (h *Api) listAssetTaskConsumables(c *gin.Context) {
	assetTaskConsumables, err := h.app.ListAssetTaskConsumables(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTaskConsumables))
}

// GetAssetTaskConsumable godoc
//
//	@Summary		Get an asset task consumable
//	@Description	Get an asset task consumable
//	@Param			groupTitle	path	string	true	"Group Title"
//	@Param			assetTitle	path	string	true	"Asset Title"
//	@Param			assetTaskId	path	string	true	"Asset Task ID"
//	@Param			consumableId	path	string	true	"Consumable ID"
//	@Produce		json
//	@Success		200	{object}	tp.AssetTaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/consumables/{consumableId} [get]
func (h *Api) getAssetTaskConsumable(c *gin.Context) {
	assetTaskConsumable, err := h.app.GetAssetTaskConsumable(c.Param(groupTitle), c.Param(assetTitle), c.Param(assetTaskId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, assetTaskConsumable))
}

// UpdateAssetTaskConsumableBody godoc
//
//	@Summary		Update an asset task consumable with json body
//	@Description	Update an asset task consumable with json body
//	@Accept			json
//	@Param 			assetTaskConsumable 	body 	tp.AssetTaskConsumable 	true 	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.AssetTaskConsumable
//	@Router			/consumables/{consumableId} [put]
func (h *Api) updateAssetTaskConsumableBody(c *gin.Context) {
	var assetTaskConsumable tp.AssetTaskConsumable
	if err := c.BindJSON(&assetTaskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTaskConsumable, err := h.app.UpdateAssetTaskConsumable(assetTaskConsumable)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskConsumable))
}

// UpdateAssetTaskConsumablePath godoc
//
//	@Summary		Update an asset task consumable with path parameters
//	@Description	Update an asset task consumable with path parameters
//	@Param 			groupTitle	path	string	true	"Group Title"
//	@Param 			assetTitle	path	string	true	"Asset Title"
//	@Param 			assetTaskId	path	string	true	"Asset Task ID"
//	@Param 			consumableId	path	string	true	"Consumable ID"
//	@Param 			assetTaskConsumable 	body 	tp.AssetTaskConsumableForPath 	true 	"Asset Task Consumable object"
//	@Produce		json
//	@Success		201	{object}	tp.AssetTaskConsumable
//	@Router			/groups/{groupTitle}/assets/{assetTitle}/asset-tasks/{assetTaskId}/consumables/{consumableId} [put]
func (h *Api) updateAssetTaskConsumablePath(c *gin.Context) {
	var assetTaskConsumable tp.AssetTaskConsumable
	if err := c.BindJSON(&assetTaskConsumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetTaskConsumable, err := h.app.UpdateAssetTaskConsumableWithValidation(
		c.Param(groupTitle),
		c.Param(assetTitle),
		c.Param(assetTaskId),
		c.Param(consumableId),
		assetTaskConsumable.QuantityNote,
	)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, assetTaskConsumable))
}
