package types

type AssetTaskConsumable struct {
	AssetTaskId  UUID   `json:"assetTaskId" binding:"required"`
	ConsumableId UUID   `json:"consumableId" binding:"required"`
	QuantityNote string `json:"quantityNote"`
}
