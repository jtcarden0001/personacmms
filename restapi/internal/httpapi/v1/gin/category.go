package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *HttpApi) registerCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryTitle", baseRoute)

	h.router.POST(baseRoute, h.createCategory)
	h.router.DELETE(individualRoute, h.deleteCategory)
	h.router.GET(baseRoute, h.listCategory)
	h.router.GET(individualRoute, h.getCategory)
	h.router.PUT(individualRoute, h.updateCategory)
}

func (h *HttpApi) createCategory(c *gin.Context) {
	var cat tp.Category
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := h.app.CreateCategory(cat.Title, cat.Description)
	if err != nil {
		processAppError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, cat) // switch to .JSON() for performance

}

func (h *HttpApi) deleteCategory(c *gin.Context) {
	title := c.Param("categoryTitle")
	err := h.app.DeleteCategory(title)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{}) // switch to .JSON() for performance
}

func (h *HttpApi) getCategory(c *gin.Context) {
	title := c.Param("categoryTitle")
	cat, err := h.app.GetCategory(title)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusOK, cat) // switch to .JSON() for performance
}

func (h *HttpApi) listCategory(c *gin.Context) {
	cats, err := h.app.ListCategory()
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusOK, cats) // switch to .JSON() for performance
}

func (h *HttpApi) updateCategory(c *gin.Context) {
	var cat tp.Category
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldTitle := c.Param("categoryTitle")
	newCat, err := h.app.UpdateCategory(oldTitle, cat.Title, cat.Description)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusOK, newCat) // switch to .JSON() for performance
}
