package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerCategoryRoutes() {
	baseRoute := fmt.Sprintf("%s/categories", routePrefix)
	individualRoute := fmt.Sprintf("%s/:categoryTitle", baseRoute)

	h.router.POST(baseRoute, h.createCategory)
	h.router.DELETE(individualRoute, h.deleteCategory)
	h.router.GET(baseRoute, h.listCategories)
	h.router.GET(individualRoute, h.getCategory)
	h.router.PUT(individualRoute, h.updateCategory)
}

// CreateCategory godoc
//
//	@Summary		Create an asset category
//	@Description	Create a category
//	@Accept			json
//	@Param			category	body	tp.Category	true	"Category object"
//	@Produce		json
//	@Success		201	{object}	tp.Category
//	@Router			/categories [post]
func (h *Api) createCategory(c *gin.Context) {
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

// DeleteCategory godoc
//
//	@Summary		Delete an asset category
//	@Description	Delete a category
//	@Param			categoryTitle	path	string	true	"Category Title"
//	@Success		204
//	@Failure		404
//	@Router			/categories/{categoryTitle} [delete]
func (h *Api) deleteCategory(c *gin.Context) {
	title := c.Param("categoryTitle")
	err := h.app.DeleteCategory(title)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{}) // switch to .JSON() for performance
}

// GetCategory godoc
//
//	@Summary		Get an asset category
//	@Description	Get a category
//	@Param			categoryTitle	path	string	true	"Category Title"
//	@Produce		json
//	@Success		200	{object}	tp.Category
//	@Router			/categories/{categoryTitle} [get]
func (h *Api) getCategory(c *gin.Context) {
	title := c.Param("categoryTitle")
	cat, err := h.app.GetCategory(title)
	if err != nil {
		processAppError(c, err)
	}

	c.IndentedJSON(http.StatusOK, cat) // switch to .JSON() for performance
}

// ListCategory godoc
//
//	@Summary		List asset categories
//	@Description	List all categories
//	@Produce		json
//	@Success		200	{object}	[]tp.Category
//	@Router			/categories [get]
func (h *Api) listCategories(c *gin.Context) {
	cats, err := h.app.ListCategory()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, cats) // switch to .JSON() for performance
	}
}

// UpdateCategory godoc
//
//	@Summary		Update an asset category
//	@Description	Update a category
//	@Accept			json
//	@Param			categoryTitle	path	string	true	"Category Title"
//	@Param			category	body	tp.Category	true	"Category object"
//	@Produce		json
//	@Success		200	{object}	tp.Category
//	@Failure		500
//	@Router			/categories/{categoryTitle} [put]
func (h *Api) updateCategory(c *gin.Context) {
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
