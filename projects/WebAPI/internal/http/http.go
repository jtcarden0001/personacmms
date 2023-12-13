package http

router := gin.Default()
	registerEquipmentRoutes(router)

	router.Run("localhost:8080")
}

func registerEquipmentRoutes(r *gin.Engine) {
	r.GET("/equipment", equipment.GetAll)
}