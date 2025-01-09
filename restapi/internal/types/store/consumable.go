package store

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance
type Consumable struct {
	Id    uuid.UUID
	Title string
}

type ConsumableQuantity struct {
	ConsumableId uuid.UUID
	Title        string
	Quantity     string
	TaskId       uuid.UUID
	WorkOrderId  uuid.UUID
}
