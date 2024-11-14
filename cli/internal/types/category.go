package types

type Category struct {
	Title       string `json:"name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}
