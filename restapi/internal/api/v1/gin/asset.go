package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var baseAssetRoute = fmt.Sprintf("%s/assets", indGroupRoute)
var assetTitle = "assetTitle"
var indAssetRoute = fmt.Sprintf("%s/:%s", baseAssetRoute, assetTitle)

func (h *Api) registerAssetRoutes() {
	h.router.POST(baseAssetRoute, h.createAsset)
	h.router.DELETE(indAssetRoute, h.deleteAsset)
	h.router.GET(baseAssetRoute, h.getAllAsset)
	h.router.GET(indAssetRoute, h.getAsset)
	h.router.PUT(indAssetRoute, h.updateAsset)
}

func (h *Api) createAsset(c *gin.Context) {
	var a tp.Asset
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.app.CreateAsset(c.Param(groupTitle), a)
	c.JSON(getStatus(err, http.StatusCreated), getResponse(err, a))
}

func (h *Api) deleteAsset(c *gin.Context) {
	err := h.app.DeleteAsset(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusNoContent), getResponse(err, nil))
}

func (h *Api) getAllAsset(c *gin.Context) {
	asset, err := h.app.ListAssets(c.Param(groupTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

func (h *Api) getAsset(c *gin.Context) {
	asset, err := h.app.GetAsset(c.Param(groupTitle), c.Param(assetTitle))
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, asset))
}

func (h *Api) updateAsset(c *gin.Context) {
	var a tp.Asset
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.app.UpdateAsset(c.Param(groupTitle), c.Param(assetTitle), a)
	c.JSON(getStatus(err, http.StatusOK), getResponse(err, a))
}
