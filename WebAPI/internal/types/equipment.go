package types

type Equipment struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Year        int    `json:"year"`
	Make        string `json:"make"`
	ModelNumber string `json:"modelNumber"`
	Description string `json:"description"`
	CategoryId  int    `json:"categoryId"`
}
