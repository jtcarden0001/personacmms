package types

import "github.com/google/uuid"

// bit self explanatory, discrete definition of a work order status
type WorkOrderStatus struct {
	Title string    `json:"title" binding:"required"`
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
}

// WorkOrderStatus Enum, if you add something here, ensure to add it to ValidWorkOrderStatuses,
// and the swagger documentation and to the db since this a db backed enum
const (
	WorkOrderStatusNew        = "new"
	WorkOrderStatusInProgress = "in progress"
	WorkOrderStatusComplete   = "complete"
	WorkOrderStatusClosed     = "closed" // not complete, not active
)

var ValidWorkOrderStatuses = map[string]bool{
	WorkOrderStatusNew:        true,
	WorkOrderStatusInProgress: true,
	WorkOrderStatusComplete:   true,
	WorkOrderStatusClosed:     true,
}

func IsValidWorkOrderStatus(ws string) bool {
	return ValidWorkOrderStatuses[ws]
}
