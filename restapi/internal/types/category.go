package types

type Category struct {
	Title       string `json:"title" binding:"required"`
	Id          UUID   `json:"id" swaggerignore:"true"`
	Description string `json:"description"`
}
