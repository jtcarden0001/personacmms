package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerUsageUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/usage-units", routePrefix)
	individualRoute := fmt.Sprintf("%s/:usageUnitTitle", baseRoute)

	h.router.POST(baseRoute, h.createUsageUnit)
	h.router.DELETE(individualRoute, h.deleteUsageUnit)
	h.router.GET(baseRoute, h.listUsageUnits)
	h.router.GET(individualRoute, h.getUsageUnit)
	h.router.PUT(individualRoute, h.updateUsageUnit)
}

// CreateUsageUnit godoc
//
//	@Summary		Create a usage unit
//	@Description	Create a usage unit
//	@Accept			json
//	@Param			usageUnit	body	tp.UsageUnit	true	"Usage Unit object"
//	@Produce		json
//	@Success		201	{object}	tp.UsageUnit
//	@Router			/usage-units [post]
func (h *Api) createUsageUnit(c *gin.Context) {
	var usageUnit tp.UsageUnit
	if err := c.BindJSON(&usageUnit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usageUnit, err := h.app.CreateUsageUnit(usageUnit)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, usageUnit))
}

// DeleteUsageUnit godoc
//
//	@Summary		Delete a usage unit
//	@Description	Delete a usage unit
//	@Param			usageUnitTitle	path	string	true	"Usage Unit Title"
//	@Success		204
//	@Failure		404
//	@Router			/usage-units/{usageUnitTitle} [delete]
func (h *Api) deleteUsageUnit(c *gin.Context) {
	err := h.app.DeleteUsageUnit(c.Param("usageUnitTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListUsageUnits godoc
//
//	@Summary		List usage units
//	@Description	List all usage units
//	@Produce		json
//	@Success		200	{object}	[]tp.UsageUnit
//	@Router			/usage-units [get]
func (h *Api) listUsageUnits(c *gin.Context) {
	usageUnits, err := h.app.ListUsageUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageUnits))
}

// GetUsageUnit godoc
//
//	@Summary		Get a usage unit
//	@Description	Get a usage unit
//	@Param			usageUnitTitle	path	string	true	"Usage Unit Title"
//	@Produce		json
//	@Success		200	{object}	tp.UsageUnit
//	@Router			/usage-units/{usageUnitTitle} [get]
func (h *Api) getUsageUnit(c *gin.Context) {
	usageUnit, err := h.app.GetUsageUnit(c.Param("usageUnitTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageUnit))
}

// UpdateUsageUnit godoc
//
//	@Summary		Update a usage unit
//	@Description	Update a usage unit
//	@Accept			json
//	@Param			usageUnitTitle	path	string		true	"Usage Unit Title"
//	@Router			/usage-units/{usageUnitTitle} [put]
func (h *Api) updateUsageUnit(c *gin.Context) {
	var usageUnit tp.UsageUnit
	if err := c.BindJSON(&usageUnit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usageUnit, err := h.app.UpdateUsageUnit(c.Param("usageUnitTitle"), usageUnit)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageUnit))
}
