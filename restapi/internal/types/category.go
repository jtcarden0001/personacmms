package types

// A Category is a logical grouping of asset types
type Category struct {
	Title       string  `json:"title" binding:"required"`
	Id          UUID    `json:"id" swaggerignore:"true"`
	Description *string `json:"description"`
}
