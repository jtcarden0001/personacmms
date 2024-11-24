package types

type Asset struct {
	GroupTitle    string `json:"groupTitle" swaggerignore:"true"`
	Title         string `json:"title" binding:"required"`
	Id            UUID   `json:"id" swaggerignore:"true"`
	Year          int    `json:"year"`
	Make          string `json:"make"`
	ModelNumber   string `json:"modelNumber"`
	SerialNumber  string `json:"serialNumber"`
	Description   string `json:"description"`
	CategoryTitle string `json:"categoryTitle"`
}
