package types

type PreventativeTask struct {
	Title       string `json:"title" binding:"required"`
	Id          UUID   `json:"id" swaggerignore:"true"`
	Description string `json:"description"`
}
