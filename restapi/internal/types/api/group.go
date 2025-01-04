package api

import "github.com/google/uuid"

// group is a logial grouping (in other words, a container) of related assets

// TODO: add references to assets
type GroupRequest struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}

type GroupResponse struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}
