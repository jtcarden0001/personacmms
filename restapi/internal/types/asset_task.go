package types

type AssetTask struct {
	Id                 UUID   `json:"id" swaggerignore:"true"` // will be handled by service
	UniqueInstructions string `json:"uniqueInstructions"`
	AssetId            UUID   `json:"assetId" swaggerignore:"true"` // will get in route path
	TaskId             UUID   `json:"taskId"`
}
