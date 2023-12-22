package gin

import "fmt"

func (h *HttpApi) registerEquipmentCategoryRoutes() {
	baseRoute := "/v1/categories"
	individualRoute := fmt.Sprintf("%s/:categoryId", baseRoute)

	h.router.POST(baseRoute, h.createEquipmentCategory)
	h.router.DELETE(individualRoute, h.deleteEquipmentCategory)
	h.router.GET(baseRoute, h.getAllEquipmentCategory)
	h.router.GET(individualRoute, h.getEquipmentCategory)
	h.router.PUT(individualRoute, h.updateEquipmentCategory)
}
