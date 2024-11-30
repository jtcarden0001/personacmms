package types

type UsageTrigger struct {
	Id        UUID   `json:"id" swaggerignore:"true"` // will get from route when needed
	Quantity  int    `json:"quantity" binding:"required"`
	UsageUnit string `json:"usage_unit" binding:"required"`
	TaskId    UUID   `json:"asset_task_id" swaggerignore:"true"` // will get from route
}
