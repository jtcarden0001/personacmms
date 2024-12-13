package types

// a task tool is the logical mapping of a tool to a task
type TaskTool struct {
	TaskId UUID `json:"taskId" binding:"required"`
	ToolId UUID `json:"toolId" binding:"required"`
}
