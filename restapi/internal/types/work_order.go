package types

import (
	tm "time"

	"github.com/google/uuid"
)

// a worker order is either a record of work done and a record of work to be done depending on the
// status of the work order
type WorkOrder struct {
	Id              uuid.UUID `json:"id" swaggerignore:"true"`
	CreatedDate     tm.Time   `json:"createdDate" swaggerignore:"true"`
	CompletedDate   *tm.Time  `json:"completedDate"`
	Instructions    *string   `json:"instructions"`
	Notes           *string   `json:"notes"`
	CumulativeMiles *int      `json:"cumulativeMiles"`
	CumulativeHours *int      `json:"cumulativeHours"`
	Status          string    `json:"status" binding:"required"`
}

// WorkOrderStatus Enum, if you add something here, ensure to add it to ValidWorkOrderStatuses,
// and the swagger documentation and to the db since this a db backed enum
const (
	WorkOrderStatusNew        = "new"
	WorkOrderStatusInProgress = "in progress"
	WorkOrderStatusComplete   = "complete"
	WorkOrderStatusCancelled  = "cancelled" // not complete, not active
)

var ValidWorkOrderStatuses = map[string]bool{
	WorkOrderStatusNew:        true,
	WorkOrderStatusInProgress: true,
	WorkOrderStatusComplete:   true,
	WorkOrderStatusCancelled:  true,
}

func IsValidWorkOrderStatus(ws string) bool {
	return ValidWorkOrderStatuses[ws]
}
