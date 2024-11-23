package gin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// external so it is accessible to AssetPreventativeTask routes
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
	var e tp.Asset
	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// id, err := h.app.CreateAsset(e.Title, e.Year, e.Make, e.ModelNumber, e.Description, e.CategoryTitle)
	// if err != nil {
	// 	processAppError(c, err)
	// } else {
	// 	e.Id = id
	// 	c.IndentedJSON(http.StatusCreated, e) // switch to .JSON() for performance
	// }
}

func (h *Api) deleteAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(assetTitle))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteAsset(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusNoContent, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllAsset(c *gin.Context) {
	asset, err := h.app.GetAllAsset()
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, asset) // switch to .JSON() for performance
	}
}

func (h *Api) getAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(assetTitle))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	asset, err := h.app.GetAsset(id)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, asset) // switch to .JSON() for performance
	}
}

func (h *Api) updateAsset(c *gin.Context) {
	var e tp.Asset

	if err := c.BindJSON(&e); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// e.Id = id // ignoring the id in the body and using the id in the url
	// err = h.app.UpdateAsset(e.Id, e.Title, e.Year, e.Make, e.ModelNumber, e.Description, e.CategoryId)
	// if err != nil {
	// 	processAppError(c, err)
	// } else {
	// 	c.IndentedJSON(http.StatusOK, e) // switch to .JSON() for performance
	// }
}
