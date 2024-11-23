package gin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerPreventativeTaskConsumableRoutes() {
	baseRoute := fmt.Sprintf("%s/asset/:assetId/preventativeTasks/:preventativeTaskId/consumables", routePrefix)
	individualRoute := fmt.Sprintf("%s/:consumableId", baseRoute)

	h.router.POST(individualRoute, h.createPreventativeTaskConsumable)
	h.router.DELETE(individualRoute, h.deletePreventativeTaskConsumable)
	h.router.GET(baseRoute, h.getAllPreventativeTaskConsumableByPreventativeTask)
	h.router.GET(individualRoute, h.getPreventativeTaskConsumable)
	h.router.PUT(individualRoute, h.updatePreventativeTaskConsumable)
}

func (h *Api) createPreventativeTaskConsumable(c *gin.Context) {
	tc := tp.PreventativeTaskConsumable{}
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.PreventativeTaskId = preventativeTaskId

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.ConsumableId = consumableId

	err = h.app.CreatePreventativeTaskConsumable(tc.PreventativeTaskId, tc.ConsumableId, tc.QuantityNote)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(201, tc) // switch to .JSON() for performance
	}
}

func (h *Api) deletePreventativeTaskConsumable(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeletePreventativeTaskConsumable(preventativeTaskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllPreventativeTaskConsumableByPreventativeTask(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskConsumables, err := h.app.GetAllPreventativeTaskConsumableByPreventativeTaskId(preventativeTaskId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskConsumables) // switch to .JSON() for performance
	}
}

func (h *Api) getPreventativeTaskConsumable(c *gin.Context) {
	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskConsumable, err := h.app.GetPreventativeTaskConsumable(preventativeTaskId, consumableId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, preventativeTaskConsumable) // switch to .JSON() for performance
	}
}

func (h *Api) updatePreventativeTaskConsumable(c *gin.Context) {
	var tc tp.PreventativeTaskConsumable
	if err := c.BindJSON(&tc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preventativeTaskId, err := strconv.Atoi(c.Param("preventativeTaskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.PreventativeTaskId = preventativeTaskId

	consumableId, err := strconv.Atoi(c.Param("consumableId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tc.ConsumableId = consumableId

	err = h.app.UpdatePreventativeTaskConsumable(tc.PreventativeTaskId, tc.ConsumableId, tc.QuantityNote)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}
