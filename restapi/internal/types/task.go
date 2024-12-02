package types

type Task struct {
	Id             UUID    `json:"id" swaggerignore:"true"` // will be handled by service
	Title          string  `json:"title"`
	Instructions   *string `json:"instructions"`
	Type           *string `json:"type"`
	AssetId        UUID    `json:"assetId" swaggerignore:"true"` // will get in route path
	TaskTemplateId *UUID   `json:"taskTemplateId"`
}
