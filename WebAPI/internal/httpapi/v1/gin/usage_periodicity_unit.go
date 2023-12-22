package gin

import "fmt"

func (h *HttpApi) registerUsagePeriodicityUnitRoutes() {
	baseRoute := "/v1/usage_periodicity_units"
	individualRoute := fmt.Sprintf("%s/:usagePeriodicityUnitId", baseRoute)

	h.router.POST(baseRoute, h.createUsagePeriodicityUnit)
	h.router.DELETE(individualRoute, h.deleteUsagePeriodicityUnit)
	h.router.GET(baseRoute, h.getAllUsagePeriodicityUnit)
	h.router.GET(individualRoute, h.getUsagePeriodicityUnit)
	h.router.PUT(individualRoute, h.updateUsagePeriodicityUnit)
}
