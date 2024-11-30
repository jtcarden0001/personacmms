package types

type TaskTool struct {
	TaskId UUID `json:"taskId" binding:"required"`
	ToolId UUID `json:"toolId" binding:"required"`
}
