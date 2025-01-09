package store

import (
	"time"

	"github.com/google/uuid"
)

// A DateTrigger is an event that is triggered on a specific date that results in a work order being created based on a task.
type DateTrigger struct {
	Id            uuid.UUID
	ScheduledDate time.Time
	TaskId        uuid.UUID
}
