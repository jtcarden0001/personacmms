package app

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// App layer hosts the business logic and forwards simple requests to the Store layer.
type App interface {
	// asset
	AssociateAssetWithCategory(assetId string, categoryId string) (tp.Asset, error)
	AssociateAssetWithGroup(assetId string, groupId string) (tp.Asset, error)
	CreateAsset(asset tp.Asset) (tp.Asset, error)
	DeleteAsset(assetId string) error
	DisassociateAssetWithCategory(assetId string, categoryId string) error
	DisassociateAssetWithGroup(assetId string, groupId string) error
	GetAsset(assetId string) (tp.Asset, error)
	ListAssets() ([]tp.Asset, error)
	ListAssetsByCategory(categoryId string) ([]tp.Asset, error)
	ListAssetsByCategoryAndGroup(categoryId string, groupId string) ([]tp.Asset, error)
	ListAssetsByGroup(groupId string) ([]tp.Asset, error)
	UpdateAsset(assetId string, asset tp.Asset) (tp.Asset, error)

	// category
	CreateCategory(category tp.Category) (tp.Category, error)
	DeleteCategory(categoryId string) error
	GetCategory(categoryId string) (tp.Category, error)
	ListCategories() ([]tp.Category, error)
	ListCategoriesByAsset(assetId string) ([]tp.Category, error)
	UpdateCategory(categoryId string, category tp.Category) (tp.Category, error)

	// consumable
	AssociateConsumableWithTask(assetId string, taskId string, consumableId string, consumableQuantity tp.ConsumableQuantity) (tp.ConsumableQuantity, error)
	AssociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string, consumableQuantity tp.ConsumableQuantity) (tp.ConsumableQuantity, error)
	CreateConsumable(consumable tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(consumableId string) error
	DisassociateConsumableWithTask(assetId string, taskId string, consumableId string) error
	DisassociateConsumableWithWorkOrder(assetId string, workOrderId string, consumableId string) error
	GetConsumable(consumableId string) (tp.Consumable, error)
	ListConsumables() ([]tp.Consumable, error)
	UpdateConsumable(consumableId string, consumable tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(assetId string, taskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(assetId string, taskId string, dateTriggerId string) error
	GetDateTrigger(assetId string, taskId string, dateTriggerId string) (tp.DateTrigger, error)
	ListDateTriggersByAssetAndTask(assetId string, taskId string) ([]tp.DateTrigger, error)
	UpdateDateTrigger(assetId string, taskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(group tp.Group) (tp.Group, error)
	DeleteGroup(groupId string) error
	GetGroup(groupId string) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	ListGroupsByAsset(assetId string) ([]tp.Group, error)
	UpdateGroup(groupId string, group tp.Group) (tp.Group, error)

	// task
	CreateTask(assetId string, task tp.Task) (tp.Task, error)
	DeleteTask(assetId string, taskId string) error
	DisassociateTaskWithWorkOrder(assetId string, taskId string, workOrderId string) error
	GetTask(assetId string, taskId string) (tp.Task, error)
	ListTasksByAsset(assetId string) ([]tp.Task, error)
	UpdateTask(assetId string, taskId string, task tp.Task) (tp.Task, error)

	// time trigger
	CreateTimeTrigger(assetId string, taskId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(assetId string, taskId string, timeTriggerId string) error
	GetTimeTrigger(assetId string, taskId string, timeTriggerId string) (tp.TimeTrigger, error)
	ListTimeTriggersByAssetAndTask(assetId string, taskId string) ([]tp.TimeTrigger, error)
	ListTimeTriggerUnits() ([]string, error)
	UpdateTimeTrigger(assetId string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error)

	// tool
	AssociateToolWithTask(assetId string, taskId string, toolId string, toolSize tp.ToolSize) (tp.ToolSize, error)
	AssociateToolWithWorkOrder(assetId string, workOrderId string, toolId string, toolSize tp.ToolSize) (tp.ToolSize, error)
	CreateTool(tool tp.Tool) (tp.Tool, error)
	DeleteTool(toolId string) error
	DisassociateToolWithTask(assetId string, taskId string, toolId string) error
	DisassociateToolWithWorkOrder(assetId string, workOrderId string, toolId string) error
	GetTool(toolId string) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	UpdateTool(toolId string, tool tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(assetId string, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(assetId string, taskId string, usageTriggerId string) error
	GetUsageTrigger(assetId string, taskId string, usageTriggerId string) (tp.UsageTrigger, error)
	ListUsageTriggersByAssetAndTask(assetId string, taskId string) ([]tp.UsageTrigger, error)
	ListUsageTriggerUnits() ([]string, error)
	UpdateUsageTrigger(assetId string, taskId string, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error)

	// work order
	AssociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) (tp.WorkOrder, error)
	CreateWorkOrder(assetId string, workOrder tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(assetId string, workOrderId string) error
	DisassociateWorkOrderWithTask(assetId string, taskId string, workOrderId string) error
	GetWorkOrder(assetId string, workOrderId string) (tp.WorkOrder, error)
	ListWorkOrdersByAsset(assetId string) ([]tp.WorkOrder, error)
	ListWorkOrderStatus() ([]string, error)
	UpdateWorkOrder(assetId string, workOrderId string, workOrder tp.WorkOrder) (tp.WorkOrder, error)
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
