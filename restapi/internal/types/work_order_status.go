package types

type WorkOrderStatus struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}

const (
	WorkOrderStatusNew        = "new"
	WorkOrderStatusInProgress = "in progress"
	WorkOrderStatusComplete   = "complete"
	WorkOrderStatusClosed     = "closed" // not complete, not active
)
