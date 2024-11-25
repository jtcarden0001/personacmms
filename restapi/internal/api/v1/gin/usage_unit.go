package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerUsageUnitRoutes() {
	baseRoute := fmt.Sprintf("%s/usage-units", routePrefix)
	h.router.GET(baseRoute, h.listUsageUnits)
}

// ListUsageUnits godoc
//
//	@Summary		List usage units
//	@Description	List all usage units
//	@Produce		json
//	@Success		200	{object}	[]types.UsageUnit
//	@Router			/usage-units [get]
func (h *Api) listUsageUnits(c *gin.Context) {
	usageUnits, err := h.app.ListUsageUnits()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, usageUnits))
}
