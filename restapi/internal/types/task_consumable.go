package types

type TaskConsumable struct {
	TaskId       UUID   `json:"taskId" binding:"required"`
	ConsumableId UUID   `json:"consumableId" binding:"required"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}

type TaskConsumableForPath struct {
	TaskId       UUID   `json:"taskId" swaggerignore:"true"`
	ConsumableId UUID   `json:"consumableId" swaggerignore:"true"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}
