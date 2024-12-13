package types

// a task consumable is the logical mapping of a consumable to a task (with quantity details)
type TaskConsumable struct {
	TaskId       UUID   `json:"taskId" binding:"required"`
	ConsumableId UUID   `json:"consumableId" binding:"required"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}

// a task consumable for path is the logical mapping of a consumable to a task (with quantity details)
// with a slightly different swagger annotations for pulling details from the REST URI path
type TaskConsumableForPath struct {
	TaskId       UUID   `json:"taskId" swaggerignore:"true"`
	ConsumableId UUID   `json:"consumableId" swaggerignore:"true"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}
