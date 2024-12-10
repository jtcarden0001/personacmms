package types

// a time trigger is an event that is triggered after a specific time has elapsed since the last time a
// work order was completed for a task.
type TimeTrigger struct {
	Id       UUID   `json:"id" swaggerignore:"true"`
	Quantity int    `json:"quantity" binding:"required"`
	TimeUnit string `json:"time_unit" binding:"required"`
	TaskId   UUID   `json:"asset_task_id" swaggerignore:"true"`
}
