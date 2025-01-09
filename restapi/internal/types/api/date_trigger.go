package api

import (
	"time"

	"github.com/google/uuid"
)

// A DateTrigger is an event that is triggered on a specific date that results in a work order being created based on a task.
type DateTriggerRequest struct {
	ScheduledDate time.Time `json:"scheduledDate" binding:"required"`
}

type DateTriggerResponse struct {
	Id            uuid.UUID `json:"dateTriggerId"`
	ScheduledDate time.Time `json:"scheduledDate"`
	TaskId        uuid.UUID `json:"taskId"`
}
