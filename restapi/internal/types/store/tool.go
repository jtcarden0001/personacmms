package store

import "github.com/google/uuid"

// a tool is a physical (non-disposable) item that is required to complete a task
type Tool struct {
	Id    uuid.UUID
	Title string
	// TODO: might be nice to add an image of the tool
}

type ToolSize struct {
	ToolId    uuid.UUID
	Title     string
	Size      *string
	Task      *Task
	WorkOrder *WorkOrder
}
