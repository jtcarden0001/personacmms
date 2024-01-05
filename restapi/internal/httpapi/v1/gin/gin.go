package gin

import (
	"github.com/gin-gonic/gin"
	a "github.com/jtcarden0001/personacmms/restapi/internal/app"
)

type HttpApi struct {
	app    a.App
	router *gin.Engine
}

func New(injectedApp a.App) *HttpApi {
	ginRouter := gin.Default()
	ginApi := &HttpApi{
		app:    injectedApp,
		router: ginRouter,
	}

	ginApi.registerRoutes()
	return ginApi
}

func (h *HttpApi) Start() {
	h.router.Run("localhost:8080")
}

var routePrefix = "/api/v1"

func (h *HttpApi) registerRoutes() {
	h.registerConsumableRoutes()
	h.registerEquipmentCategoryRoutes()
	h.registerEquipmentRoutes()
	h.registerTaskConsumableRoutes()
	h.registerTaskToolRoutes()
	h.registerTaskRoutes()
	h.registerTimePeriodicityUnitRoutes()
	h.registerToolRoutes()
	h.registerUsagePeriodicityUnitRoutes()
	h.registerWorkOrderStatusRoutes()
	h.registerWorkOrderRoutes()
}
