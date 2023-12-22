package gin

import "fmt"

func (h *httpApi) registerTimePeriodicityUnitRoutes() {
	baseRoute := "/v1/time-periodicity-units"
	individualRoute := fmt.Sprintf("%s/:timePeriodicityUnitId", baseRoute)

	h.router.POST(baseRoute, h.createTimePeriodicityUnit)
	h.router.DELETE(individualRoute, h.deleteTimePeriodicityUnit)
	h.router.GET(baseRoute, h.getAllTimePeriodicityUnit)
	h.router.GET(individualRoute, h.getTimePeriodicityUnit)
	h.router.PUT(individualRoute, h.updateTimePeriodicityUnit)
}
