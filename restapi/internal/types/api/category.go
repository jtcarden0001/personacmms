package api

import "github.com/google/uuid"

// A Category is a logical grouping of asset types
type CategoryRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
}

type CategoryResponse struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
	// TODO:
	// array string of api references "/api/v1/assets/{id}"
	// Assets     []string `json:"assets"`
}
