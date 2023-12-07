package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jtcarden0001/personacmms/projects/webapi/pkg/handler/equipment"
)

// TODO: adding data access layer code feels very copy pasty, how can I improve the structure to reduce this?
func main() {
	router := gin.Default()
	registerEquipmentRoutes(router)

	router.Run("localhost:8080")
}

func registerEquipmentRoutes(r *gin.Engine) {
	r.GET("/equipment", equipment.GetAll)
}
