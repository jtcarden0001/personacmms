package types

import "github.com/google/uuid"

// A usage trigger is an even that is triggered after a specific usage threshold has been reached
// since the last time a task was completed.  Like a generator being used for 30 hours, etc.
type UsageTrigger struct {
	Id        uuid.UUID `json:"id" swaggerignore:"true"` // will get from route when needed
	Quantity  int       `json:"quantity" binding:"required"`
	UsageUnit string    `json:"usage_unit" binding:"required"`
	TaskId    uuid.UUID `json:"asset_task_id" swaggerignore:"true"` // will get from route
}
