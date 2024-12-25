package types

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance
type Consumable struct {
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
	Title string    `json:"title" binding:"required"`
}
