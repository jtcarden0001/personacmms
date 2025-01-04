package api

import "github.com/google/uuid"

// A Category is a logical grouping of asset types

// TODO: add references to assets
type CategoryRequest struct {
	Id          uuid.UUID `json:"id" swaggerignore:"true"`
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
}

type CategoryResponse struct {
	Id          uuid.UUID `json:"id" swaggerignore:"true"`
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
}
