package app

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
)

// App layer hosts the business logic and forwards simple requests to the Store layer.
type App interface {
	// asset
	AssociateAssetWithCategory(assetId string, categoryId string) (apitp.AssetResponse, error)
	AssociateAssetWithGroup(assetId string, groupId string) (apitp.AssetResponse, error)
	CreateAsset(asset apitp.AssetRequest) (apitp.AssetResponse, error)
	DeleteAsset(assetId string) error
	DisassociateAssetWithCategory(assetId string, categoryId string) error
	DisassociateAssetWithGroup(assetId string, groupId string) error
	GetAsset(assetId string) (apitp.AssetResponse, error)
	ListAssets() ([]apitp.AssetResponse, error)
	ListAssetsByCategory(categoryId string) ([]apitp.AssetResponse, error)
	ListAssetsByCategoryAndGroup(categoryId string, groupId string) ([]apitp.AssetResponse, error)
	ListAssetsByGroup(groupId string) ([]apitp.AssetResponse, error)
	UpdateAsset(assetId string, asset apitp.AssetRequest) (apitp.AssetResponse, error)

	// category
	CreateCategory(category apitp.CategoryRequest) (apitp.CategoryResponse, error)
	DeleteCategory(categoryId string) error
	GetCategory(categoryId string) (apitp.CategoryResponse, error)
	ListCategories() ([]apitp.CategoryResponse, error)
	ListCategoriesByAsset(assetId string) ([]apitp.CategoryResponse, error)
	UpdateCategory(categoryId string, category apitp.CategoryRequest) (apitp.CategoryResponse, error)

	// consumable
	AssociateConsumableWithTask(assetId string, taskId string, consumableId string, consumableQuantity apitp.ConsumableQuantityRequest) (apitp.ConsumableQuantityResponse, error)
	AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string, consumableQuantity apitp.ConsumableQuantityRequest) (apitp.ConsumableQuantityResponse, error)
	CreateConsumable(consumable apitp.ConsumableRequest) (apitp.ConsumableResponse, error)
	DeleteConsumable(consumableId string) error
	DisassociateConsumableWithTask(assetId string, taskId string, consumableId string) error
	DisassociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) error
	GetConsumable(consumableId string) (apitp.ConsumableResponse, error)
	ListConsumables() ([]apitp.ConsumableRequest, error)
	UpdateConsumable(consumableId string, consumable apitp.ConsumableRequest) (apitp.ConsumableResponse, error)

	// date trigger
	CreateDateTrigger(assetId string, taskId string, dateTrigger apitp.DateTriggerRequest) (apitp.DateTriggerResponse, error)
	DeleteDateTrigger(assetId string, taskId string, dateTriggerId string) error
	GetDateTrigger(assetId string, taskId string, dateTriggerId string) (apitp.DateTriggerResponse, error)
	ListDateTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.DateTriggerResponse, error)
	UpdateDateTrigger(assetId string, taskId string, dateTriggerId string, dateTrigger apitp.DateTriggerRequest) (apitp.DateTriggerResponse, error)

	// group
	CreateGroup(group apitp.GroupRequest) (apitp.GroupResponse, error)
	DeleteGroup(groupId string) error
	GetGroup(groupId string) (apitp.GroupResponse, error)
	ListGroups() ([]apitp.GroupRequest, error)
	ListGroupsByAsset(assetId string) ([]apitp.GroupResponse, error)
	UpdateGroup(groupId string, group apitp.GroupRequest) (apitp.GroupResponse, error)

	// task
	CreateTask(assetId string, task apitp.TaskRequest) (apitp.TaskResponse, error)
	DeleteTask(assetId string, taskId string) error
	DisassociateTaskWithWorkOrder(assetId string, taskId string, workOrderId string) error
	GetTask(assetId string, taskId string) (apitp.TaskResponse, error)
	ListTasksByAsset(assetId string) ([]apitp.TaskResponse, error)
	UpdateTask(assetId string, taskId string, task apitp.TaskRequest) (apitp.TaskResponse, error)

	// time trigger
	CreateTimeTrigger(assetId string, taskId string, timeTrigger apitp.TimeTriggerRequest) (apitp.TimeTriggerResponse, error)
	DeleteTimeTrigger(assetId string, taskId string, timeTriggerId string) error
	GetTimeTrigger(assetId string, taskId string, timeTriggerId string) (apitp.TimeTriggerResponse, error)
	ListTimeTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.TimeTriggerResponse, error)
	ListTimeTriggerUnits() ([]string, error)
	UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger apitp.TimeTriggerRequest) (apitp.TimeTriggerResponse, error)

	// tool
	AssociateToolWithTask(assetId string, taskId string, toolId string, toolSize apitp.ToolSizeRequest) (apitp.ToolSizeResponse, error)
	AssociateToolWithWorkOrder(assetId string, workOrderId string, toolId string, toolSize apitp.ToolSizeRequest) (apitp.ToolSizeResponse, error)
	CreateTool(tool apitp.ToolRequest) (apitp.ToolResponse, error)
	DeleteTool(toolId string) error
	DisassociateToolWithTask(assetId string, taskId string, toolId string) error
	DisassociateToolWithWorkOrder(assetId string, workOrderId string, toolId string) error
	GetTool(toolId string) (apitp.ToolResponse, error)
	ListTools() ([]apitp.ToolResponse, error)
	UpdateTool(toolId string, tool apitp.ToolRequest) (apitp.ToolResponse, error)

	// usage trigger
	CreateUsageTrigger(assetId string, taskId string, usageTrigger apitp.UsageTriggerRequest) (apitp.UsageTriggerResponse, error)
	DeleteUsageTrigger(assetId string, taskId string, usageTriggerId string) error
	GetUsageTrigger(assetId string, taskId string, usageTriggerId string) (apitp.UsageTriggerResponse, error)
	ListUsageTriggersByAssetAndTask(assetId string, taskId string) ([]apitp.UsageTriggerResponse, error)
	ListUsageTriggerUnits() ([]string, error)
	UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger apitp.UsageTriggerRequest) (apitp.UsageTriggerResponse, error)

	// work order
	AssociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) (apitp.WorkOrderResponse, error)
	CreateWorkOrder(assetId string, workOrder apitp.WorkOrderRequest) (apitp.WorkOrderResponse, error)
	DeleteWorkOrder(assetId string, workOrderId string) error
	DisassociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) error
	GetWorkOrder(assetId string, workOrderId string) (apitp.WorkOrderResponse, error)
	ListWorkOrdersByAsset(assetId string) ([]apitp.WorkOrderResponse, error)
	ListWorkOrderStatus() ([]string, error)
	UpdateWorkOrder(assetId string, workOrderId string, workOrder apitp.WorkOrderRequest) (apitp.WorkOrderResponse, error)
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
