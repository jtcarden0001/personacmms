package types

import "github.com/google/uuid"

// A Category is a logical grouping of asset types
type Category struct {
	Id          uuid.UUID `json:"id" swaggerignore:"true"`
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
}
