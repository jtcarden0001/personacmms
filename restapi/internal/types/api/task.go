package api

import "github.com/google/uuid"

// a task is a task that is assigned to an asset that work orders can be spawned from
type TaskRequest struct {
	Title        string  `json:"title" binding:"required"`
	Instructions *string `json:"instructions"`
}

type TaskResponse struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Instructions *string   `json:"instructions"`
	AssetId      uuid.UUID `json:"-"`
	// api reference /api/v1/assets/{id}
	AssetReference string `json:"parentAsset"`
	// TODO:
	// array string of api references "/api/v1/assets/{assetId}/tasks/{id}/tools/{id}"
	// Tools []string `json:"associatedTools"`
	// array string of api references "/api/v1/assets/{assetId}/tasks/{id}/consumables/{id}
	// Consumables []string `json:"associatedConsumables"`
}
