package types

import "time"

// A DateTrigger is an event that is triggered on a specific date that results in a work order being created based on a task.
type DateTrigger struct {
	Id     UUID      `json:"id" swaggerignore:"true"`
	Date   time.Time `json:"date"`
	TaskId UUID      `json:"asset_task_id" swaggerignore:"true"` // will pull from route
}
