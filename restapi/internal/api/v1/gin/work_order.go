package gin

import (
	"fmt"
	"strconv"
	tm "time"

	"github.com/gin-gonic/gin"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (h *Api) registerWorkOrderRoutes() {
	baseRoute := fmt.Sprintf("%s/work-orders", routePrefix)
	baseRouteByTask := fmt.Sprintf("%s/asset/:assetId/tasks/:taskId/work-orders", routePrefix)
	individualRouteByTask := fmt.Sprintf("%s/:workOrderId", baseRouteByTask)

	h.router.POST(baseRouteByTask, h.createWorkOrderByTask)
	h.router.DELETE(individualRouteByTask, h.deleteWorkOrderByTask)
	h.router.GET(baseRoute, h.getAllWorkOrder)
	h.router.GET(baseRouteByTask, h.getAllWorkOrderByTask)
	h.router.GET(individualRouteByTask, h.getWorkOrderByTask)
	h.router.PUT(individualRouteByTask, h.updateWorkOrderByTask)
}

func (h *Api) createWorkOrderByTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var wo tp.WorkOrder
	if err := c.BindJSON(&wo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo.TaskId = taskId
	id, err := h.app.CreateWorkOrder(wo.TaskId, wo.StatusId, wo.CreatedDate, wo.CompletedDate)
	if err != nil {
		processAppError(c, err)
	} else {
		wo.Id = id
		c.IndentedJSON(201, wo) // switch to .JSON() for performance
	}
}

func (h *Api) deleteWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	workOrderId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.app.DeleteWorkOrder(workOrderId)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

func (h *Api) getAllWorkOrder(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woss, err := h.app.GetAllWorkOrder()
	if err != nil {
		processAppError(c, err)
	} else {
		iwoss, err := h.interpolateWorkOrders(woss)
		if err != nil {
			processAppError(c, err)
		} else {
			c.IndentedJSON(200, iwoss) // switch to .JSON() for performance
		}
	}
}

func (h *Api) getAllWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woss, err := h.app.GetAllWorkOrder() // TODO: BUG: this is getting all work orders and not filtering by task
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(200, woss) // switch to .JSON() for performance
	}
}

func (h *Api) getWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo, err := h.app.GetWorkOrder(woId)
	if err != nil {
		processAppError(c, err)
		return
	}

	iwo, err := h.interpolateWorkOrder(wo)
	if err != nil {
		processAppError(c, err)
		return
	}

	c.IndentedJSON(200, iwo) // switch to .JSON() for performance
}

func (h *Api) updateWorkOrderByTask(c *gin.Context) {
	// while we don't use the taskId, a work order is always associated with a task
	// should we change the route to exclude the taskId? or should we do any validation on the taskid?
	woId, err := strconv.Atoi(c.Param("workOrderId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var wo tp.WorkOrder
	if err := c.BindJSON(&wo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	wo.Id = woId
	err = h.app.UpdateWorkOrder(wo.Id, wo.StatusId, wo.CreatedDate, wo.CompletedDate)
	if err != nil {
		processAppError(c, err)
	} else {
		c.IndentedJSON(204, gin.H{}) // switch to .JSON() for performance
	}
}

type interpolatedWorkOrder struct {
	Id            int      `json:"id"`
	TaskId        int      `json:"taskId"`
	TaskTitle     string   `json:"taskTitle" binding:"required"`
	StatusTitle   string   `json:"statusTitle" binding:"required"`
	AssetId       int      `json:"assetId" binding:"required"`
	AssetTitle    string   `json:"assetTitle" binding:"required"`
	CreatedDate   tm.Time  `json:"createdDate" binding:"required"`
	CompletedDate *tm.Time `json:"completedDate"`
}

func (h *Api) interpolateWorkOrders(woss []tp.WorkOrder) ([]interpolatedWorkOrder, error) {
	var iwoss []interpolatedWorkOrder
	var err error
	for _, wo := range woss {
		iwo, err := h.interpolateWorkOrder(wo)
		if err != nil {
			return nil, err
		}

		iwoss = append(iwoss, iwo)
	}

	return iwoss, err
}

// may want to move this interpolation logic down the stack to reduce db calls and allow db joins to do the work
// or do some memoization to reduce the db calls. Will do if performance suffers, no need for premature optimization.
func (h *Api) interpolateWorkOrder(wo tp.WorkOrder) (interpolatedWorkOrder, error) {
	t, err := h.app.GetTask(wo.TaskId)
	if err != nil {
		return interpolatedWorkOrder{}, err
	}

	e, err := h.app.GetAsset(t.AssetId)
	if err != nil {
		return interpolatedWorkOrder{}, err
	}

	s, err := h.app.GetWorkOrderStatus(wo.StatusId)
	if err != nil {
		return interpolatedWorkOrder{}, err
	}

	return interpolatedWorkOrder{
		Id:            wo.Id,
		TaskId:        wo.TaskId,
		TaskTitle:     t.Title,
		StatusTitle:   s.Title,
		AssetId:       t.AssetId,
		AssetTitle:    e.Title,
		CreatedDate:   wo.CreatedDate,
		CompletedDate: wo.CompletedDate,
	}, nil
}
