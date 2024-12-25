package types

import "github.com/google/uuid"

// moved out of scope for MVP

// a task template is a predefined task that a real task can be spawned from and assigned to an asset
type TaskTemplate struct {
	Title       string    `json:"title" binding:"required"`
	Id          uuid.UUID `json:"id" swaggerignore:"true"`
	Description *string   `json:"description"`
	Type        *string   `json:"type"`
}
