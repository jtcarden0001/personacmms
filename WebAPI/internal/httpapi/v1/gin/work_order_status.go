package gin

import "fmt"

func (h *HttpApi) registerWorkOrderStatusRoutes() {
	baseRoute := "/v1/work-order-statuses"
	individualRoute := fmt.Sprintf("%s/:workOrderStatusId", baseRoute)

	h.router.POST(baseRoute, h.createWorkOrderStatus)
	h.router.DELETE(individualRoute, h.deleteWorkOrderStatus)
	h.router.GET(baseRoute, h.getAllWorkOrderStatus)
	h.router.GET(individualRoute, h.getWorkOrderStatus)
	h.router.PUT(individualRoute, h.updateWorkOrderStatus)
}
