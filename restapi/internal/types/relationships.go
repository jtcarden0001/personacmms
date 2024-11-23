package types

type PreventativeTaskTool struct {
	PreventativeTaskId int `json:"preventativeTaskId"`
	ToolId             int `json:"toolId"`
}

type PreventativeTaskConsumable struct {
	PreventativeTaskId int    `json:"preventativeTaskId"`
	ConsumableId       int    `json:"consumableId"`
	QuantityNote       string `json:"quantityNote" binding:"required"`
}
