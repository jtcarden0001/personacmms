package types

import "github.com/google/uuid"

// a task tool is the logical mapping of a tool to a task
type TaskTool struct {
	TaskId uuid.UUID `json:"taskId" binding:"required"`
	ToolId uuid.UUID `json:"toolId" binding:"required"`
}
