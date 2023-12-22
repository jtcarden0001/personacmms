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
	instance := &HttpApi{
		app:    injectedApp,
		router: ginRouter,
	}

	instance.registerEquipmentRoutes(ginRouter)
	instance.registerToolRoutes(ginRouter)

	return instance
}

func (h *HttpApi) Start() {
	h.router.Run("localhost:8080")
}

var routePrefix = "/api/v1"
