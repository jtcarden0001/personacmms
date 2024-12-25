package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var usageUnitGp = "usage-units"

var baseUsageUnitRoute = fmt.Sprintf("%s/%s", routePrefix, usageUnitGp)

func (h *Api) registerUsageUnitRoutes() {
	h.router.GET(baseUsageUnitRoute, h.listUsageUnits)
}

// ListUsageUnits godoc
//
//	@Summary		List usage units
//	@Description	List all usage units tht can be used with usage triggers
//	@Tags			usage-units
//	@Produce		json
//	@Success		200	{object}	[]types.UsageUnit
//	@Failure		500	{object}	map[string]any
//	@Router			/usage-units [get]
func (h *Api) listUsageUnits(c *gin.Context) {
	usageUnits, err := h.app.ListUsageTriggerUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageUnits))
}
