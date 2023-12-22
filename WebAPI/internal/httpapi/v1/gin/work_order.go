package gin

import "fmt"

func (h *HttpApi) registerWorkOrderRoutes() {
	baseRouteByTask := "/v1/equipment/:equipmentId/tasks/:taskId/work-orders"
	individualRouteByTask := fmt.Sprintf("%s/:workOrderId", baseRouteByTask)

	h.router.POST(baseRouteByTask, h.createWorkOrderByTask)
	h.router.DELETE(individualRouteByTask, h.deleteWorkOrderByTask)
	h.router.GET(baseRouteByTask, h.getAllWorkOrderByTask)
	h.router.GET(individualRouteByTask, h.getWorkOrderByTask)
	h.router.PUT(individualRouteByTask, h.updateWorkOrderByTask)
}
