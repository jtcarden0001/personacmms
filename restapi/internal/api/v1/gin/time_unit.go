package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var timeUnitGp = "time-units"

var baseTimeUnitRoute = fmt.Sprintf("%s/%s", routePrefix, timeUnitGp)

func (h *Api) registerTimeUnitRoutes() {
	h.router.GET(baseTimeUnitRoute, h.listTimeUnits)
}

// ListTimeUnits godoc
//
//	@Summary		List time units
//	@Description	List all time units that can be used with time-triggers and usage-triggers
//	@Tags			time-units
//	@Produce		json
//	@Success		200	{object}	[]types.TimeUnit
//	@Failure		500	{object}	map[string]any
//	@Router			/time-units [get]
func (h *Api) listTimeUnits(c *gin.Context) {
	timeUnits, err := h.app.ListTimeTriggerUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnits))
}
