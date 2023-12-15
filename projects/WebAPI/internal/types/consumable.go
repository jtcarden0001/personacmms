package types

type Consumable struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}
