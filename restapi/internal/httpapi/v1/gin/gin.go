package gin

import (
	"time"

	"github.com/gin-contrib/cors"
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
	ginApi.configureCORS()
	return ginApi
}

func (h *HttpApi) Start() {
	h.router.Run("localhost:8080")
}

var routePrefix = "/api/v1"

func (h *HttpApi) registerRoutes() {
	h.registerConsumableRoutes()
	h.registerCategoryRoutes()
	h.registerAssetRoutes()
	h.registerTaskConsumableRoutes()
	h.registerTaskToolRoutes()
	h.registerTaskRoutes()
	h.registerTimePeriodicityUnitRoutes()
	h.registerToolRoutes()
	h.registerUsagePeriodicityUnitRoutes()
	h.registerWorkOrderStatusRoutes()
	h.registerWorkOrderRoutes()
}

func (h *HttpApi) configureCORS() {
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - All the methods
	// - Origin header
	// - Credentials shareT ODO: what is this?
	// - Preflight requests cached for 12 hours TODO: what is this?
	h.router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin, Content-Type, Authorization, X-Requested-With"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
