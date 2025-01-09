package api

import "github.com/google/uuid"

// a tool is a physical (non-disposable) item that is required to complete a task
type ToolRequest struct {
	Title string `json:"title" binding:"required"`
	// TODO: might be nice to add an image of the tool
}

type ToolResponse struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	// TODO:
	// api references "/assets/{id}/tasks/{id}"
	// TaskReferences []string `json:"taskReferences"`
	// active work order references "/assets/{id}workorders/{id}"
	// WorkOrderReferences []string `json:"activeWorkOrderReferences"`
	// TODO: might be nice to add an image of the tool
}

type ToolSizeRequest struct {
	Size *string `json:"size"`
}

type ToolSizeResponse struct {
	Title  string    `json:"title" swaggerignore:"true"`
	Size   *string   `json:"size"`
	ToolId uuid.UUID `json:"-"`
	// api references "/tools/{id}"
	ToolReference string `json:"toolReference"`
	// api references "/assets/{id}/tasks/{id}"
	TaskReference string `json:"taskReference,omitempty"`
	// active work order references "/assets/{id}workorders/{id}"
	WorkOrderReference string `json:"workOrderReference,omitempty"`
}
