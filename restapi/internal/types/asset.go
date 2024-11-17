package types

type Asset struct {
	GroupTitle    string `json:"groupTitle" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Id            string `json:"id" swaggerignore:"true"`
	Year          int    `json:"year"`
	Make          string `json:"make"`
	ModelNumber   string `json:"modelNumber"`
	Description   string `json:"description"`
	CategoryTitle string `json:"categoryTitle"`
}
