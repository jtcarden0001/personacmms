package types

type TimeTrigger struct {
	Id          UUID   `json:"id" swaggerignore:"true"`
	Quantity    int    `json:"quantity" binding:"required"`
	TimeUnit    string `json:"time_unit" binding:"required"`
	AssetTaskId UUID   `json:"asset_task_id" swaggerignore:"true"`
}
