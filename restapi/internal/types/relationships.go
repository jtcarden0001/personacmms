package types

type TaskTool struct {
	TaskId int `json:"preventativeTaskId"`
	ToolId int `json:"toolId"`
}

type TaskConsumable struct {
	TaskId       int    `json:"preventativeTaskId"`
	ConsumableId int    `json:"consumableId"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}
