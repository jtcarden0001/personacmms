package types

type Consumable struct {
	Title string `json:"title" binding:"required"`
	Id    UUID   `json:"id" swaggerignore:"true"`
}
