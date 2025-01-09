package api

import "github.com/google/uuid"

// group is a logial grouping (in other words, a container) of related assets
type GroupRequest struct {
	Title string `json:"title" binding:"required"`
}

type GroupResponse struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	// TODO:
	// array string of api references "/api/v1/assets/{id}"
	// Assets []string `json:"associatedAssets"`
}
