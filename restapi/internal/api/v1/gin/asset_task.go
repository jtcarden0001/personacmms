package gin

import "fmt"

func (h *Api) registerAssetTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/tasks", baseAssetRoute)
	individualRoute := fmt.Sprintf("%s/:taskId", baseRoute)
	_ = individualRoute
}
