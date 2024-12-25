package types

import "github.com/google/uuid"

// a time trigger is an event that is triggered after a specific time has elapsed since the last time a
// work order was completed for a task.
type TimeTrigger struct {
	Id       uuid.UUID `json:"id" swaggerignore:"true"`
	Quantity int       `json:"quantity" binding:"required"`
	TimeUnit string    `json:"time_unit" binding:"required"`
	TaskId   uuid.UUID `json:"asset_task_id" swaggerignore:"true"`
}
