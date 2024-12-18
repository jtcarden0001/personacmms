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
//	@Produce		json
//	@Success		200	{object}	[]types.TimeUnit
//	@Router			/time-units [get]
func (h *Api) listTimeUnits(c *gin.Context) {
	timeUnits, err := h.app.ListTimeUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, timeUnits))
}
