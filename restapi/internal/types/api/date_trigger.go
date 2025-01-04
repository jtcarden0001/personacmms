package api

import (
	"time"

	"github.com/google/uuid"
)

// A DateTrigger is an event that is triggered on a specific date that results in a work order being created based on a task.
type DateTriggerRequest struct {
	Id            uuid.UUID `json:"id" swaggerignore:"true"`
	ScheduledDate time.Time `json:"scheduled_date"`
	TaskId        uuid.UUID `json:"asset_task_id" swaggerignore:"true"` // will pull from route
}

type DateTriggerResponse struct {
	Id            uuid.UUID `json:"id" swaggerignore:"true"`
	ScheduledDate time.Time `json:"scheduled_date"`
	TaskId        uuid.UUID `json:"asset_task_id" swaggerignore:"true"` // will pull from route
}
