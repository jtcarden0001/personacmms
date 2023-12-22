package gin

import "fmt"

func (h *HttpApi) registerTaskToolRoutes() {
	baseRoute := "/v1/equipment/:equipmentId/tasks/:taskId/tools"
	individualRoute := fmt.Sprintf("%s/:toolId", baseRoute)

	h.router.POST(baseRoute, h.createTaskTool)
	h.router.DELETE(individualRoute, h.deleteTaskTool)
	h.router.GET(baseRoute, h.getAllTaskToolByTask)
	h.router.GET(individualRoute, h.getTaskTool)
}
