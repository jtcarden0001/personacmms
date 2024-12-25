package types

import "github.com/google/uuid"

// a tool is a physical (non-disposable) item that is required to complete a task
type Tool struct {
	Title string    `json:"title" binding:"required"`
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Size  *string   `json:"size" binding:"required"`
	// TODO: might be nice to add an image of the tool
}
