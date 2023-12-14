package types

type Equipment struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Year        int    `json:"year"`
	Make        string `json:"make"`
	ModelNumber string `json:"model_number"`
	Description string `json:"description"`
}
