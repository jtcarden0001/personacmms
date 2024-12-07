package types

import "time"

type DateTrigger struct {
	Id     UUID      `json:"id" swaggerignore:"true"`
	Date   time.Time `json:"date"`
	TaskId UUID      `json:"asset_task_id" swaggerignore:"true"` // will pull from route
}
