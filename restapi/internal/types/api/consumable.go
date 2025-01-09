package api

import "github.com/google/uuid"

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance

// TODO: add references to tasks and work orders
type ConsumableRequest struct {
	Title string `json:"title" binding:"required"`
}

type ConsumableResponse struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	// TODO: maybe add references to tasks and work orders
	// array string of api references "/api/v1/tasks/{id}"
	// Tasks []string `json:"associatedTasks"`
	// array string of api references "/api/v1/workorders/{id}"
	// WorkOrders []string `json:"associatedWorkOrders"`
}

type ConsumableQuantityRequest struct {
	Quantity string `json:"quantity" binding:"required"`
}

type ConsumableQuantityResponse struct {
	ConsumableId uuid.UUID `json:"consumableId"`
	Quantity     string    `json:"quantity"`
	TaskId       uuid.UUID `json:"taskId,omitempty"`
	WorkOrderId  uuid.UUID `json:"workOrderId,omitempty"`
}
