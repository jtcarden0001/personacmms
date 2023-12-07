package types

type Equipment struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Tool struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
