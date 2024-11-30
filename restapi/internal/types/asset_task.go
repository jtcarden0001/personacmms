package types

type AssetTask struct {
	Id                 UUID   `json:"id" swaggerignore:"true"` // will be handled by service
	Title              string `json:"title"`
	UniqueInstructions string `json:"uniqueInstructions"`
	AssetId            UUID   `json:"assetId" swaggerignore:"true"` // will get in route path
	TaskTemplateId     UUID   `json:"taskTemplateId"`
}
