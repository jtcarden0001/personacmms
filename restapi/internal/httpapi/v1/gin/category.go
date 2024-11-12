package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	uid "github.com/google/uuid"
)

func (h *HttpApi) registerCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryId", baseRoute)

	h.router.POST(baseRoute, h.createCategory)
	h.router.DELETE(individualRoute, h.deleteCategory)
	h.router.GET(baseRoute, h.listCategory)
	h.router.GET(individualRoute, h.getCategory)
	h.router.PUT(individualRoute, h.updateCategory) // accepts object id in url, disregards id in body, may revisit this design
}

func (h *HttpApi) createCategory(c *gin.Context) {
	var ec tp.Category
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.app.CreateCategory(ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		ec.Id = id
		c.IndentedJSON(201, ec) // switch to .JSON() for performance
	}
}

func (h *HttpApi) deleteCategory(c *gin.Context) {
	id, err :=  // strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *HttpApi) listCategory(c *gin.Context) {
	assetCategories, err := h.app.ListCategory()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, assetCategories) // switch to .JSON() for performance
	}
}

func (h *HttpApi) getCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	category, err := h.app.GetCategory(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, category) // switch to .JSON() for performance
	}
}

func (h *HttpApi) updateCategory(c *gin.Context) {
	var ec tp.Category
	if err := c.BindJSON(&ec); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.UpdateCategory(id, ec.Title)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
