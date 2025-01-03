package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /consumables (JSON) done
// - GET  /consumables/{consumableId} done
// - GET  /consumables done
// - PUT  /consumables/{consumableId} (JSON) done
// - DEL  /consumables/{consumableId} done
//
// - PUT  /assets/{assetId}/tasks/{taskId}/consumables/{consumableId} done
// - DEL  /assets/{assetId}/tasks/{taskId}/consumables/{consumableId} done
// - PUT  /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} done
// - DEL  /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} done

var consumableId = "consumableId"
var consumableGp = "consumables"
var consumableResource = fmt.Sprintf("%s/:%s", consumableGp, consumableId)

var baseConsumableRoute = fmt.Sprintf("%s/%s", routePrefix, consumableGp)
var indConsumableRoute = fmt.Sprintf("%s/%s", routePrefix, consumableResource)

func (h *Api) registerConsumableRoutes() {
	h.router.POST(baseConsumableRoute, h.createConsumable)

	h.router.DELETE(indConsumableRoute, h.deleteConsumable)
	h.router.DELETE(fmt.Sprintf("%s/%s/%s", indAssetRoute, taskResource, consumableResource), h.disassociateConsumableWithTask)
	h.router.DELETE(fmt.Sprintf("%s/%s/%s", indAssetRoute, workOrderResource, consumableResource), h.disassociateConsumableWithWorkOrder)

	h.router.GET(baseConsumableRoute, h.listConsumables)
	h.router.GET(indConsumableRoute, h.getConsumable)

	h.router.PUT(fmt.Sprintf("%s/%s/%s", indAssetRoute, taskResource, consumableResource), h.associateConsumableWithTask)
	h.router.PUT(fmt.Sprintf("%s/%s/%s", indAssetRoute, workOrderResource, consumableResource), h.associateConsumableWithWorkOrder)
	h.router.PUT(indConsumableRoute, h.updateConsumable)
}

// AssociateConsumableWithTask godoc
//
//	@Summary		Associate a consumable with a task
//	@Description	Associate a consumable with a task
//	@Tags			consumables
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string					true	"Asset ID"
//	@Param			taskId			path		string					true	"Task ID"
//	@Param			consumableId	path		string					true	"Consumable ID"
//	@Param			consumable		body		tp.ConsumableQuantity	true	"Consumable object"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/consumables/{consumableId} [put]
func (h *Api) associateConsumableWithTask(c *gin.Context) {
	var consumable tp.ConsumableQuantity
	if err := c.BindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	consumable, err := h.app.AssociateConsumableWithTask(c.Param(assetId), c.Param(taskId), c.Param(consumableId), consumable)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumable))
}

// AssociateConsumableWithWorkOrder godoc
//
//	@Summary		Associate a consumable with a work order
//	@Description	Associate a consumable with a work order
//	@Tags			consumables
//	@Accept			json
//	@Produce		json
//	@Param			assetId			path		string					true	"Asset Id"
//	@Param			workOrderId		path		string					true	"Work Order Id"
//	@Param			consumableId	path		string					true	"Consumable Id"
//	@Param			consumable		body		tp.ConsumableQuantity	true	"Consumable object"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} [put]
func (h *Api) associateConsumableWithWorkOrder(c *gin.Context) {
	var consumable tp.ConsumableQuantity
	if err := c.BindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	consumable, err := h.app.AssociateConsumableWithWorkOrder(c.Param(assetId), c.Param(workOrderId), c.Param(consumableId), consumable)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumable))
}

// CreateConsumable godoc
//
//	@Summary		Create a consumable
//	@Description	Create a consumable
//	@Tags			consumables
//	@Accept			json
//	@Produce		json
//	@Param			consumable	body		tp.Consumable	true	"Consumable object"
//	@Success		201			{object}	tp.Consumable
//	@Failure		400			{object}	map[string]any
//	@Failure		404			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/consumables [post]
func (h *Api) createConsumable(c *gin.Context) {
	var consumable tp.Consumable
	if err := c.BindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	consumable, err := h.app.CreateConsumable(consumable)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, consumable))
}

// DeleteConsumable godoc
//
//	@Summary		Delete a consumable
//	@Description	Delete a consumable
//	@Tags			consumables
//	@Param			consumableId	path	string	true	"Consumable Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/consumables/{consumableId} [delete]
func (h *Api) deleteConsumable(c *gin.Context) {
	err := h.app.DeleteConsumable(c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateConsumableWithTask godoc
//
//	@Summary		Disassociate a consumable with a task
//	@Description	Disassociate a consumable with a task
//	@Tags			consumables
//	@Param			assetId			path	string	true	"Asset Id"
//	@Param			taskId			path	string	true	"Task Id"
//	@Param			consumableId	path	string	true	"Consumable Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/tasks/{taskId}/consumables/{consumableId} [delete]
func (h *Api) disassociateConsumableWithTask(c *gin.Context) {
	err := h.app.DisassociateConsumableWithTask(c.Param(assetId), c.Param(taskId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// DisassociateConsumableWithWorkOrder godoc
//
//	@Summary		Disassociate a consumable with a work order
//	@Description	Disassociate a consumable with a work order
//	@Tags			consumables
//	@Param			assetId			path	string	true	"Asset Id"
//	@Param			workOrderId		path	string	true	"Work Order Id"
//	@Param			consumableId	path	string	true	"Consumable Id"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId} [delete]
func (h *Api) disassociateConsumableWithWorkOrder(c *gin.Context) {
	err := h.app.DisassociateConsumableWithWorkOrder(c.Param(assetId), c.Param(workOrderId), c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetConsumable godoc
//
//	@Summary		Get a consumable
//	@Description	Get a consumable
//	@Tags			consumables
//	@Produce		json
//	@Param			consumableId	path		string	true	"Consumable Id"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/consumables/{consumableId} [get]
func (h *Api) getConsumable(c *gin.Context) {
	consumable, err := h.app.GetConsumable(c.Param(consumableId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumable))
}

// ListConsumables godoc
//
//	@Summary		List consumables
//	@Description	List all consumables
//	@Tags			consumables
//	@Produce		json
//	@Success		200	{object}	[]tp.Consumable
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/consumables [get]
func (h *Api) listConsumables(c *gin.Context) {
	consumables, err := h.app.ListConsumables()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumables))
}

// UpdateConsumable godoc
//
//	@Summary		Update a consumable
//	@Description	Update a consumable
//	@Tags			consumables
//	@Accept			json
//	@Produce		json
//	@Param			consumableId	path		string			true	"Consumable Id"
//	@Param			consumable		body		tp.Consumable	true	"Consumable object"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/consumables/{consumableId} [put]
func (h *Api) updateConsumable(c *gin.Context) {
	var consumable tp.Consumable
	if err := c.BindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	consumable, err := h.app.UpdateConsumable(c.Param(consumableId), consumable)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumable))
}
