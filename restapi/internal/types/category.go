package types

type Category struct {
	Title       string `json:"title" binding:"required"`
	Id          UUID   `json:"id"`
	Description string `json:"description"`
}
