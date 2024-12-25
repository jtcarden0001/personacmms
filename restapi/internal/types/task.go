package types

import "github.com/google/uuid"

// a task is a task that is assigned to an asset that work orders can be spawned from
type Task struct {
	Id             uuid.UUID  `json:"id" swaggerignore:"true"` // will be handled by service
	Title          string     `json:"title"`
	Instructions   *string    `json:"instructions"`
	Type           *string    `json:"type"`
	AssetId        uuid.UUID  `json:"assetId" swaggerignore:"true"` // will get in route path
	TaskTemplateId *uuid.UUID `json:"taskTemplateId"`
}
