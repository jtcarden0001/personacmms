package types

type Task struct {
	Title       string `json:"title" binding:"required"`
	Id          UUID   `json:"id" swaggerignore:"true"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

// TaskType Enum
const (
	_ = iota
	TaskTypePreventative
	TaskTypeCorrective
)
