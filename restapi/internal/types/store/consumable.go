package types

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance

// TODO: add references to tasks and work orders
type Consumable struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}

type ConsumableQuantity struct {
	Id       uuid.UUID `json:"id" swaggerignore:"true"`
	Title    string    `json:"title" swaggerignore:"true"`
	Quantity string    `json:"quantity" binding:"required"`
}
