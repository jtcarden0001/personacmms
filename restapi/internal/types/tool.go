package types

type Tool struct {
	Title string  `json:"title" binding:"required"`
	Id    UUID    `json:"id" swaggerignore:"true"`
	Size  *string `json:"size" binding:"required"`
	// TODO: might be nice to add an image of the tool
}
