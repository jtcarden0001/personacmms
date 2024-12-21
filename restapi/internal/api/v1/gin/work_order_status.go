package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerWorkOrderStatusRoutes() {
	baseRoute := fmt.Sprintf("%s/work-order-statuses", routePrefix)

	h.router.GET(baseRoute, h.listWorkOrderStatus)
}

// ListWorkOrderStatus godoc
//
//	@Summary		List work order statuses
//	@Description	List all work order statuses
//	@Tags			work-order-statuses
//	@Produce		json
//	@Success		200	{object}	[]types.WorkOrderStatus
//	@Failure		500	{object}	map[string]any
//	@Router			/work-order-statuses [get]
func (h *Api) listWorkOrderStatus(c *gin.Context) {
	wos, err := h.app.ListWorkOrderStatus()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, wos))
}
