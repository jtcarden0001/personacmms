package gin

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	a "github.com/jtcarden0001/personacmms/restapi/internal/app"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jtcarden0001/personacmms/restapi/internal/api/v1/docs" // This line is required for swaggo to find docs
)

//	@title			PersonaCMMS API
//	@version		1.0
//	@description	This is the Personal Computer Maintenance Management System REST API.

//	@contact.email	greenrivercodelabs@gmail.com

//	@license.name	GNU Affero General Public License v3.0
//	@license.url	https://spdx.org/licenses/AGPL-3.0-or-later.html

//	@host		127.0.0.1:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

type Api struct {
	app    a.App
	router *gin.Engine
}

func New(injectedApp a.App) *Api {
	ginRouter := gin.Default()
	ginApi := &Api{
		app:    injectedApp,
		router: ginRouter,
	}

	ginApi.registerRoutes()
	ginApi.configureCORS()
	return ginApi
}

func (h *Api) Start() {
	h.router.Run(":8080") //TODO: source this from configuration
}

var routePrefix = "/api/v1"

func (h *Api) registerRoutes() {
	h.registerAssetRoutes()
	h.registerCategoryRoutes()
	h.registerConsumableRoutes()
	h.registerDateTriggerRoutes()
	h.registerGroupRoutes()
	h.registerTaskRoutes()
	// h.registerTaskTemplateRoutes() // cut out of MVP scope
	h.registerTimeTriggerRoutes()
	h.registerTimeUnitRoutes()
	h.registerToolRoutes()
	h.registerUsageTriggerRoutes()
	h.registerUsageUnitRoutes()
	h.registerWorkOrderRoutes()
	h.registerWorkOrderStatusRoutes()
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (h *Api) configureCORS() {
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - All the methods
	// - Origin header
	// - Credentials shareT ODO: what is this?
	// - Preflight requests cached for 12 hours TODO: what is this?
	h.router.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		// AllowOrigins:  []string{"http://localhost:5173"},
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
