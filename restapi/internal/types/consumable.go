package types

// A Consumable is a consumable item that is required to complete a maintenance task for an asset.
type Consumable struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}
