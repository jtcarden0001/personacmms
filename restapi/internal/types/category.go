package types

import "github.com/google/uuid"

// A Category is a logical grouping of asset types
type Category struct {
	Title       string    `json:"title" binding:"required"`
	Id          uuid.UUID `json:"id" swaggerignore:"true"`
	Description *string   `json:"description"`
}
