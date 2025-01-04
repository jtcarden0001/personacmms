package api

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance

// TODO: add references to tasks and work orders
type ConsumableRequest struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}

type ConsumableResponse struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}

type ConsumableQuantityRequest struct {
	Id       uuid.UUID `json:"id" swaggerignore:"true"`
	Title    string    `json:"title" swaggerignore:"true"`
	Quantity string    `json:"quantity" binding:"required"`
}

type ConsumableQuantityResponse struct {
	Id       uuid.UUID `json:"id" swaggerignore:"true"`
	Title    string    `json:"title" swaggerignore:"true"`
	Quantity string    `json:"quantity" binding:"required"`
}
