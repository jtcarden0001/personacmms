package types

type AssetTaskTool struct {
	AssetTaskId UUID `json:"assetTaskId" binding:"required"`
	ToolId      UUID `json:"toolId" binding:"required"`
}
