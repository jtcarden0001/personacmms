package types

import (
	"strings"
	tm "time"

	"github.com/google/uuid"
)

// a worker order is either a record of work done and a record of work to be done on an asset depending
// on the status of the work order

// TODO: add references to tasks, workers, and consumables
type WorkOrder struct {
	Id              uuid.UUID `json:"id" swaggerignore:"true"`
	Title           string    `json:"title" binding:"required"`
	CreatedDate     tm.Time   `json:"createdDate" swaggerignore:"true"`
	CompletedDate   *tm.Time  `json:"completedDate"`
	Instructions    *string   `json:"instructions"`
	Notes           *string   `json:"notes"`
	CumulativeMiles *int      `json:"cumulativeMiles"`
	CumulativeHours *int      `json:"cumulativeHours"`
	AssetId         uuid.UUID `json:"assetId" swaggerignore:"true"`
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

func PrintValidWorkOrderStatuses() string {
	statuses := []string{}
	for status := range ValidWorkOrderStatuses {
		statuses = append(statuses, status)
	}
	return strings.Join(statuses, ", ")
}
