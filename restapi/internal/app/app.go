package app

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// App layer hosts the business logic and forwards simple requests to the Store layer.
type App interface {
	// asset
	CreateAsset(string, tp.Asset) (tp.Asset, error)
	DeleteAsset(string, string) error
	ListAssets(string) ([]tp.Asset, error)
	GetAsset(string, string) (tp.Asset, error)
	UpdateAsset(string, string, tp.Asset) (tp.Asset, error)

	// asset task
	CreateAssetTask(string, string, tp.AssetTask) (tp.AssetTask, error)
	DeleteAssetTask(string, string, string) error
	ListAssetTasks(string, string) ([]tp.AssetTask, error)
	GetAssetTask(string, string, string) (tp.AssetTask, error)
	UpdateAssetTask(string, string, string, tp.AssetTask) (tp.AssetTask, error)

	// asset task consumable
	CreateAssetTaskConsumable(tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error)
	CreateAssetTaskConsumableWithValidation(string, string, string, string, string) (tp.AssetTaskConsumable, error)
	DeleteAssetTaskConsumable(string, string, string, string) error
	ListAssetTaskConsumables(string, string, string) ([]tp.AssetTaskConsumable, error)
	GetAssetTaskConsumable(string, string, string, string) (tp.AssetTaskConsumable, error)
	UpdateAssetTaskConsumable(tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error)
	UpdateAssetTaskConsumableWithValidation(string, string, string, string, string) (tp.AssetTaskConsumable, error)

	// asset task tools
	CreateAssetTaskTool(tp.AssetTaskTool) (tp.AssetTaskTool, error)
	CreateAssetTaskToolWithValidation(string, string, string, string) (tp.AssetTaskTool, error)
	DeleteAssetTaskTool(string, string, string, string) error
	ListAssetTaskTools(string, string, string) ([]tp.AssetTaskTool, error)
	GetAssetTaskTool(string, string, string, string) (tp.AssetTaskTool, error)

	// category
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(string) error
	ListCategories() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, tp.Category) (tp.Category, error)

	// consumable
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(string) error
	ListConsumables() ([]tp.Consumable, error)
	GetConsumable(string) (tp.Consumable, error)
	UpdateConsumable(string, tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(string, string, string, tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(string, string, string, string) error
	ListDateTriggers(string, string, string) ([]tp.DateTrigger, error)
	GetDateTrigger(string, string, string, string) (tp.DateTrigger, error)
	UpdateDateTrigger(string, string, string, string, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	ListGroups() ([]tp.Group, error)
	GetGroup(string) (tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task template
	CreateTaskTemplate(tp.TaskTemplate) (tp.TaskTemplate, error)
	DeleteTaskTemplate(string) error
	ListTaskTemplates() ([]tp.TaskTemplate, error)
	GetTaskTemplate(string) (tp.TaskTemplate, error)
	UpdateTaskTemplate(string, tp.TaskTemplate) (tp.TaskTemplate, error)

	// time trigger
	CreateTimeTrigger(string, string, string, tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(string, string, string, string) error
	ListTimeTriggers(string, string, string) ([]tp.TimeTrigger, error)
	GetTimeTrigger(string, string, string, string) (tp.TimeTrigger, error)
	UpdateTimeTrigger(string, string, string, string, tp.TimeTrigger) (tp.TimeTrigger, error)

	// time unit
	CreateTimeUnit(tp.TimeUnit) (tp.TimeUnit, error)
	DeleteTimeUnit(string) error
	ListTimeUnits() ([]tp.TimeUnit, error)
	GetTimeUnit(string) (tp.TimeUnit, error)
	UpdateTimeUnit(string, tp.TimeUnit) (tp.TimeUnit, error)

	// tool
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	ListTools() ([]tp.Tool, error)
	GetTool(string) (tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(string, string, string, tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(string, string, string, string) error
	ListUsageTriggers(string, string, string) ([]tp.UsageTrigger, error)
	GetUsageTrigger(string, string, string, string) (tp.UsageTrigger, error)
	UpdateUsageTrigger(string, string, string, string, tp.UsageTrigger) (tp.UsageTrigger, error)

	// usage unit
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	ListUsageUnits() ([]tp.UsageUnit, error)
	GetUsageUnit(string) (tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)

	// work order status
	CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(title string) error
	ListWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)

	// work order
	CreateWorkOrder(string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
	DeleteAssetTaskWorkOrder(string, string, string, string) error
	ListAssetTaskWorkOrders(string, string, string) ([]tp.WorkOrder, error)
	GetAssetTaskWorkOrder(string, string, string, string) (tp.WorkOrder, error)
	UpdateAssetTaskWorkOrder(string, string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
