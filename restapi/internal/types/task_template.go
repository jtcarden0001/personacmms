package types

// a task template is a predefined task that a real task can be spawned from and assigned to an asset
type TaskTemplate struct {
	Title       string  `json:"title" binding:"required"`
	Id          UUID    `json:"id" swaggerignore:"true"`
	Description *string `json:"description"`
	Type        *string `json:"type"`
}

// TaskType Enum, if you add something here, ensure to add it to ValidTaskTypes, and
// the swagger documentation
const (
	TaskTypePreventative = "preventative"
	TaskTypeCorrective   = "corrective"
)

// ValidTaskTypes is a map of valid task types
var ValidTaskTypes = map[string]bool{
	TaskTypePreventative: true,
	TaskTypeCorrective:   true,
}

func IsValidTaskType(tt string) bool {
	return ValidTaskTypes[tt]
}
