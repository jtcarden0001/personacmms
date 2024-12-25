package types

import "github.com/google/uuid"

// a tool is a physical (non-disposable) item that is required to complete a task
type Tool struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
	// TODO: might be nice to add an image of the tool
}

type ToolSize struct {
	Size string `json:"size" binding:"required"`
}
