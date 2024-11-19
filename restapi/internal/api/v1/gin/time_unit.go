package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTimeUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/time-units", routePrefix)
	individualRoute := fmt.Sprintf("%s/:timeUnitTitle", baseRoute)

	h.router.POST(baseRoute, h.createTimeUnit)
	h.router.DELETE(individualRoute, h.deleteTimeUnit)
	h.router.GET(baseRoute, h.listTimeUnits)
	h.router.GET(individualRoute, h.getTimeUnit)
	h.router.PUT(individualRoute, h.updateTimeUnit)
}

// CreateTimeUnit godoc
//
//	@Summary		Create a time unit
//	@Description	Create a time unit
//	@Accept			json
//	@Param			timeUnit	body	tp.TimeUnit	true	"Time Unit object"
//	@Produce		json
//	@Success		201	{object}	tp.TimeUnit
//	@Router			/time-units [post]
func (h *Api) createTimeUnit(c *gin.Context) {
	var timeUnit tp.TimeUnit
	if err := c.BindJSON(&timeUnit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeUnit, err := h.app.CreateTimeUnit(timeUnit)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, timeUnit))
}

// DeleteTimeUnit godoc
//
//	@Summary		Delete a time unit
//	@Description	Delete a time unit
//	@Param			timeUnitTitle	path	string	true	"Time Unit Title"
//	@Success		204
//	@Failure		404
//	@Router			/time-units/{timeUnitTitle} [delete]
func (h *Api) deleteTimeUnit(c *gin.Context) {
	err := h.app.DeleteTimeUnit(c.Param("timeUnitTitle"))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// ListTimeUnits godoc
//
//	@Summary		List time units
//	@Description	List all time units
//	@Produce		json
//	@Success		200	{object}	[]tp.TimeUnit
//	@Router			/time-units [get]
func (h *Api) listTimeUnits(c *gin.Context) {
	timeUnits, err := h.app.ListTimeUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnits))
}

// GetTimeUnit godoc
//
//	@Summary		Get a time unit
//	@Description	Get a time unit
//	@Param			timeUnitTitle	path	string	true	"Time Unit Title"
//	@Produce		json
//	@Success		200	{object}	tp.TimeUnit
//	@Router			/time-units/{timeUnitTitle} [get]
func (h *Api) getTimeUnit(c *gin.Context) {
	timeUnit, err := h.app.GetTimeUnit(c.Param("timeUnitTitle"))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnit))
}

// UpdateTimeUnit godoc
//
//	@Summary		Update a time unit
//	@Description	Update a time unit
//	@Accept			json
//	@Param			timeUnitTitle	path	string		true	"Time Unit Title"
//	@Produce		json
//	@Success		200	{object}	tp.TimeUnit
//	@Router			/time-units/{timeUnitTitle} [put]
func (h *Api) updateTimeUnit(c *gin.Context) {
	var timeUnit tp.TimeUnit
	if err := c.BindJSON(&timeUnit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeUnit, err := h.app.UpdateTimeUnit(c.Param("timeUnitTitle"), timeUnit)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnit))
}
