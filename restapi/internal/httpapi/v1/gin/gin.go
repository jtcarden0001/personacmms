package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jtcarden0001/personacmms/webapi/internal/app"
)

type HttpApi struct {
	app    app.App
	router *gin.Engine
}

func New(injectedApp app.App) *HttpApi {
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
