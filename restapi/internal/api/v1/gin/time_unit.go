package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerTimeUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/time-units", routePrefix)

	h.router.GET(baseRoute, h.listTimeUnits)
}

// ListTimeUnits godoc
//
//	@Summary		List time units
//	@Description	List all time units
//	@Tags			time-units
//	@Produce		json
//	@Success		200	{object}	[]types.TimeUnit
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/time-units [get]
func (h *Api) listTimeUnits(c *gin.Context) {
	timeUnits, err := h.app.ListTimeUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnits))
}
