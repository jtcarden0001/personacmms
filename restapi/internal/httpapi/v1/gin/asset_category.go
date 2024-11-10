package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *HttpApi) registerAssetCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryId", baseRoute)

	h.router.POST(baseRoute, h.createAssetCategory)
	h.router.DELETE(individualRoute, h.deleteAssetCategory)
	h.router.GET(baseRoute, h.getAllAssetCategory)
	h.router.GET(individualRoute, h.getAssetCategory)
	h.router.PUT(individualRoute, h.updateAssetCategory) // accepts object id in url, disregards id in body, may revisit this design
}

func (h *HttpApi) createAssetCategory(c *gin.Context) {
	var ec tp.AssetCategory
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateAssetCategory(ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		ec.Id = id
		c.IndentedJSON(201, ec) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteAssetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteAssetCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAllAssetCategory(c *gin.Context) {
	assetCategories, err := h.app.GetAllAssetCategory()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, assetCategories) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getAssetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	assetCategory, err := h.app.GetAssetCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, assetCategory) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateAssetCategory(c *gin.Context) {
	var ec tp.AssetCategory
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateAssetCategory(id, ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
