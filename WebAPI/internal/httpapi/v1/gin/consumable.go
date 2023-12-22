package gin

func (h *HttpApi) registerConsumableRoutes() {
	h.router.POST("/v1/consumables", h.createConsumable)
	h.router.DELETE("/v1/consumables/:id", h.deleteConsumable)
	h.router.GET("/v1/consumables", h.getAllConsumable)
	h.router.GET("/v1/consumables/:id", h.getConsumable)
	h.router.PUT("/v1/consumables/:id", h.updateConsumable)
}
