package api

import "github.com/google/uuid"

// a task is a task that is assigned to an asset that work orders can be spawned from

// TODO: add references to tools and consumables
type TaskRequest struct {
	Id           uuid.UUID `json:"id" swaggerignore:"true"` // will be handled by service
	Title        string    `json:"title" binding:"required"`
	Instructions *string   `json:"instructions"`
	AssetId      uuid.UUID `json:"assetId" swaggerignore:"true"`
}

type TaskResponse struct {
	Id           uuid.UUID `json:"id" swaggerignore:"true"` // will be handled by service
	Title        string    `json:"title" binding:"required"`
	Instructions *string   `json:"instructions"`
	AssetId      uuid.UUID `json:"assetId" swaggerignore:"true"`
}
