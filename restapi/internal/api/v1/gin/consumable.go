package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// - POST /consumables (JSON)
// - GET  /consumables/{consumableId}
// - GET  /consumables
// - PUT  /consumables/{consumableId} (JSON)
// - DEL  /consumables/{consumableId}
//
// - PUT  /assets/{assetId}/tasks/{taskId}/consumables/{consumableId}
// - DEL  /assets/{assetId}/tasks/{taskId}/consumables/{consumableId}
// - PUT  /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId}
// - DEL  /assets/{assetId}/work-orders/{workOrderId}/consumables/{consumableId}

func (h *Api) registerConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableTitle", baseRoute)

	h.router.POST(baseRoute, h.createConsumable)
	h.router.DELETE(individualRoute, h.deleteConsumable)
	h.router.GET(baseRoute, h.listConsumables)
	h.router.GET(individualRoute, h.getConsumable)
	h.router.PUT(individualRoute, h.updateConsumable)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
//	@Param			consumableTitle	path	string	true	"Consumable ID"
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/consumables/{consumableTitle} [delete]
func (h *Api) deleteConsumable(c *gin.Context) {
	err := h.app.DeleteConsumable(c.Param("consumableTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// GetConsumable godoc
//
//	@Summary		Get a consumable
//	@Description	Get a consumable
//	@Tags			consumables
//	@Produce		json
//	@Param			consumableTitle	path		string	true	"Consumable ID"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/consumables/{consumableTitle} [get]
func (h *Api) getConsumable(c *gin.Context) {
	consumable, err := h.app.GetConsumable(c.Param("consumableTitle"))
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
//	@Param			consumableTitle	path		string			true	"Consumable ID"
//	@Param			consumable		body		tp.Consumable	true	"Consumable object"
//	@Success		200				{object}	tp.Consumable
//	@Failure		400				{object}	map[string]any
//	@Failure		404				{object}	map[string]any
//	@Failure		500				{object}	map[string]any
//	@Router			/consumables/{consumableTitle} [put]
func (h *Api) updateConsumable(c *gin.Context) {
	var consumable tp.Consumable
	if err := c.BindJSON(&consumable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consumable, err := h.app.UpdateConsumable(c.Param("consumableTitle"), consumable)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, consumable))
}
