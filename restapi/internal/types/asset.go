package types

type Asset struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Make        string `json:"make" binding:"required"`
	ModelNumber string `json:"modelNumber" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryId  int    `json:"categoryId"`
}
