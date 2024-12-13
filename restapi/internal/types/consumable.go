package types

// A Consumable is a consumable item that is available to map to a particular task if
// required by the maintenance
type Consumable struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}
