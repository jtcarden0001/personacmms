package api

import "github.com/google/uuid"

// a tool is a physical (non-disposable) item that is required to complete a task

// TODO: add references to tasks and work orders
type ToolRequest struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
	// TODO: might be nice to add an image of the tool
}

type ToolResponse struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
	// TODO: might be nice to add an image of the tool
}

type ToolSizeRequest struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"tool" swaggerignore:"true"`
	Size  *string   `json:"size"`
}

type ToolSizeResponse struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"tool" swaggerignore:"true"`
	Size  *string   `json:"size"`
}
