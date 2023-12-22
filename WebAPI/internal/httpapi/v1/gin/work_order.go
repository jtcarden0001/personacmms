package gin

import "fmt"

func (h *HttpApi) registerWorkOrderRoutes() {
	baseRoute := "/v1/work-orders"
	baseRouteByEquipment := "/v1/equipment/:equipmentId/work-orders"
	baseRouteByTask := "/v1/equipment/:equipmentId/tasks/:taskId/work-orders"

	h.router.POST(baseRoute, h.createWorkOrder)
	h.router.POST(baseRouteByEquipment, h.createWorkOrderByEquipment)
	h.router.POST(baseRouteByTask, h.createWorkOrderByTask)
	h.router.DELETE(fmt.Sprintf("%s/:workOrderId", baseRoute), h.deleteWorkOrder)
	h.router.DELETE(fmt.Sprintf("%s/:workOrderId", baseRouteByEquipment), h.deleteWorkOrderByEquipment)
	h.router.DELETE(fmt.Sprintf("%s/:workOrderId", baseRouteByTask), h.deleteWorkOrderByTask)
	h.router.GET(baseRoute, h.getAllWorkOrder)
	h.router.GET(fmt.Sprintf("%s/:workOrderId", baseRoute), h.getWorkOrder)
	h.router.GET(baseRouteByEquipment, h.getAllWorkOrderByEquipment)
	h.router.GET(fmt.Sprintf("%s/:workOrderId", baseRouteByEquipment), h.getWorkOrderByEquipment)
	h.router.GET(baseRouteByTask, h.getAllWorkOrderByTask)
	h.router.GET(fmt.Sprintf("%s/:workOrderId", baseRouteByTask), h.getWorkOrderByTask)
	h.router.PUT(fmt.Sprintf("%s/:workOrderId", baseRoute), h.updateWorkOrder)
	h.router.PUT(fmt.Sprintf("%s/:workOrderId", baseRouteByEquipment), h.updateWorkOrderByEquipment)
	h.router.PUT(fmt.Sprintf("%s/:workOrderId", baseRouteByTask), h.updateWorkOrderByTask)
}
