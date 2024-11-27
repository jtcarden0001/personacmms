package types

type AssetTaskConsumable struct {
	AssetTaskId  UUID   `json:"assetTaskId" binding:"required"`
	ConsumableId UUID   `json:"consumableId" binding:"required"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}

type AssetTaskConsumableForPath struct {
	AssetTaskId  UUID   `json:"assetTaskId" swaggerignore:"true"`
	ConsumableId UUID   `json:"consumableId" swaggerignore:"true"`
	QuantityNote string `json:"quantityNote" binding:"required"`
}
