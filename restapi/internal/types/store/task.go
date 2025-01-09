package store

import "github.com/google/uuid"

// a task is a task that is assigned to an asset that work orders can be spawned from
type Task struct {
	Id                     uuid.UUID
	Title                  string
	Instructions           *string
	AssetId                uuid.UUID
	ToolSizeList           []ToolSize
	ConsumableQuantityList []ConsumableQuantity
}
