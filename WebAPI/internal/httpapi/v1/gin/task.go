package gin

import "fmt"

func (h *HttpApi) registerTaskRoutes() {
	baseRoute := "/v1/equipment/:equipmentId/tasks"
	individualRoute := fmt.Sprintf("%s/:taskId", baseRoute)

	h.router.POST(baseRoute, h.createTask)
	h.router.DELETE(individualRoute, h.deleteTask)
	h.router.GET(baseRoute, h.getAllTaskByEquipment)
	h.router.GET(individualRoute, h.getTask)
	h.router.PUT(individualRoute, h.updateTask)
}
