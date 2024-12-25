package types

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance
type Consumable struct {
	Title string    `json:"title" binding:"required"`
	Id    uuid.UUID `json:"id" swaggerignore:"true"`
}
