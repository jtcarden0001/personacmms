package types

type TaskTool struct {
	TaskId int `json:"taskId"`
	ToolId int `json:"toolId"`
}

type TaskConsumable struct {
	TaskId       int    `json:"taskId"`
	ConsumableId int    `json:"consumableId"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}
