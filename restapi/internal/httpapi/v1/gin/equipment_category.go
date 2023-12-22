package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *HttpApi) registerEquipmentCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryId", baseRoute)

	h.router.POST(baseRoute, h.createEquipmentCategory)
	h.router.DELETE(individualRoute, h.deleteEquipmentCategory)
	h.router.GET(baseRoute, h.getAllEquipmentCategory)
	h.router.GET(individualRoute, h.getEquipmentCategory)
	h.router.PUT(individualRoute, h.updateEquipmentCategory)
}

func (h *HttpApi) createEquipmentCategory(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) deleteEquipmentCategory(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getAllEquipmentCategory(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) getEquipmentCategory(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}

func (h *HttpApi) updateEquipmentCategory(c *gin.Context) {
	c.JSON(503, gin.H{"error": fmt.Errorf("not implemented")})
	return
}
