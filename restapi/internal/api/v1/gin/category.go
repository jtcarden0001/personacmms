package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
)

// - POST /categories (JSON) done
// - GET  /categories/{categoryId} done
// - GET  /categories done
// - PUT  /categories/{categoryId} (JSON) done
// - DEL  /categories/{categoryId} done
//
// - GET /assets/{assetId}/categories done

var categoryId = "categoryId"
var categoryGp = "categories"
var categoryResource = fmt.Sprintf("%s/:%s", categoryGp, categoryId)

var baseCategoryRoute = fmt.Sprintf("%s/%s", routePrefix, categoryGp)
var indCategoryRoute = fmt.Sprintf("%s/%s", routePrefix, categoryResource)

func (h *Api) registerCategoryRoutes() {
	h.router.POST(baseCategoryRoute, h.createCategory)

	h.router.DELETE(indCategoryRoute, h.deleteCategory)

	h.router.GET(indCategoryRoute, h.getCategory)
	h.router.GET(baseCategoryRoute, h.listCategories)
	h.router.GET(fmt.Sprintf("%s/%s", indAssetRoute, categoryGp), h.listCategoriesByAsset)

	h.router.PUT(indCategoryRoute, h.updateCategory)
}

// createCategory godoc
//
//	@Summary		Create an asset category
//	@Description	Create an asset category
//	@Tags			categories
//	@Accept			json
//	@Param			category	body	apitp.CategoryRequest	true	"Category object"
//	@Produce		json
//	@Success		201	{object}	apitp.CategoryResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories [post]
func (h *Api) createCategory(c *gin.Context) {
	var cat apitp.CategoryRequest
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	cat, err := h.app.CreateCategory(cat)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, cat))
}

// deleteCategory godoc
//
//	@Summary		Delete an asset category
//	@Description	Delete an asset category
//	@Tags			categories
//	@Param			categoryId	path	string	true	"Category Id"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId} [delete]
func (h *Api) deleteCategory(c *gin.Context) {
	err := h.app.DeleteCategory(c.Param(categoryId))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

// getCategory godoc
//
//	@Summary		Get an asset category
//	@Description	Get a category
//	@Tags			categories
//	@Param			categoryId	path	string	true	"Category Id"
//	@Produce		json
//	@Success		200	{object}	apitp.CategoryResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId} [get]
func (h *Api) getCategory(c *gin.Context) {
	cat, err := h.app.GetCategory(c.Param(categoryId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, cat))
}

// listCategories godoc
//
//	@Summary		List asset categories
//	@Description	List asset categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	[]apitp.CategoryResponse
//	@Failure		500	{object}	map[string]any
//	@Router			/categories [get]
func (h *Api) listCategories(c *gin.Context) {
	cats, err := h.app.ListCategories()
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, cats))
}

// listCategoriesByAsset godoc
//
//	@Summary		List asset categories
//	@Description	List asset categories
//	@Tags			categories
//	@Param			assetId	path	string	true	"Asset Id"
//	@Produce		json
//	@Success		200	{object}	[]apitp.CategoryResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/assets/{assetId}/categories [get]
func (h *Api) listCategoriesByAsset(c *gin.Context) {
	cats, err := h.app.ListCategoriesByAsset(c.Param(assetId))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, cats))
}

// updateCategory godoc
//
//	@Summary		Update an asset category
//	@Description	Update an asset category
//	@Tags			categories
//	@Accept			json
//	@Param			categoryId	path	string					true	"Category Id"
//	@Param			category	body	apitp.CategoryRequest	true	"Category object"
//	@Produce		json
//	@Success		200	{object}	apitp.CategoryResponse
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/categories/{categoryId} [put]
func (h *Api) updateCategory(c *gin.Context) {
	var cat apitp.CategoryRequest
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	newCat, err := h.app.UpdateCategory(c.Param(categoryId), cat)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, newCat))
}
