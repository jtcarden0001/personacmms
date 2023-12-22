package gin

import "fmt"

func (h *HttpApi) registerTaskConsumableRoutes() {
	baseRoute := "/v1/equipment/:equipmentId/tasks/:taskId/consumables"
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(baseRoute, h.createTaskConsumable)
	h.router.DELETE(fmt.Sprintf(individualRoute, baseRoute), h.deleteTaskConsumable)
	h.router.GET(baseRoute, h.getAllTaskConsumableByTask)
	h.router.GET(individualRoute, h.getTaskConsumable)
	h.router.PUT(individualRoute, h.taskConsumableUpdate)
}
