package types

type TaskTool struct {
	TaskId int `json:"taskId" binding:"required"`
	ToolId int `json:"toolId" binding:"required"`
}

type TaskConsumable struct {
	TaskId       int `json:"taskId" binding:"required"`
	ConsumableId int `json:"consumableId" binding:"required"`
	QuantityNote int `json:"quantityNote" binding:"required"`
}

type EquipmentTask struct {
	EquipmentId int `json:"equipmentId" binding:"required"`
	TaskId      int `json:"taskId" binding:"required"`
}
