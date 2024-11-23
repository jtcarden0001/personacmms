package gin

import "fmt"

func (h *Api) registerAssetPreventativeTaskRoutes() {
	baseRoute := fmt.Sprintf("%s/preventativeTasks", baseAssetRoute)
	individualRoute := fmt.Sprintf("%s/:preventativeTaskId", baseRoute)
	_ = individualRoute
}
