package store

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// Store layer implements persistence for the application.
type Store interface {
	// asset
	CreateAsset(tp.Asset) (tp.Asset, error)
	DeleteAsset(string, string) error
	ListAssets() ([]tp.Asset, error)
	GetAsset(string, string) (tp.Asset, error)
	UpdateAsset(string, string, tp.Asset) (tp.Asset, error)

	// asset task
	CreateAssetTask(tp.AssetTask) (tp.AssetTask, error)
	DeleteAssetTask(tp.UUID) error
	ListAssetTasks() ([]tp.AssetTask, error)
	GetAssetTask(tp.UUID) (tp.AssetTask, error)
	UpdateAssetTask(tp.UUID, tp.AssetTask) (tp.AssetTask, error)

	// asset task consumable
	CreateAssetTaskConsumable(tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error)
	DeleteAssetTaskConsumable(tp.UUID, tp.UUID) error
	ListAssetTaskConsumables() ([]tp.AssetTaskConsumable, error)
	GetAssetTaskConsumable(tp.UUID, tp.UUID) (tp.AssetTaskConsumable, error)
	UpdateAssetTaskConsumable(tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error)

	// asset task tool
	CreateAssetTaskTool(tp.AssetTaskTool) (tp.AssetTaskTool, error)
	DeleteAssetTaskTool(tp.UUID, tp.UUID) error
	ListAssetTaskTools() ([]tp.AssetTaskTool, error)
	GetAssetTaskTool(tp.UUID, tp.UUID) (tp.AssetTaskTool, error)

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
	CreateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(tp.UUID) error
	ListDateTriggers() ([]tp.DateTrigger, error)
	GetDateTrigger(tp.UUID) (tp.DateTrigger, error)
	UpdateDateTrigger(tp.UUID, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	ListGroups() ([]tp.Group, error)
	GetGroup(string) (tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(string) error
	ListTasks() ([]tp.Task, error)
	GetTask(string) (tp.Task, error)
	UpdateTask(string, tp.Task) (tp.Task, error)

	// time trigger
	CreateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(tp.UUID) error
	ListTimeTriggers() ([]tp.TimeTrigger, error)
	GetTimeTrigger(tp.UUID) (tp.TimeTrigger, error)
	UpdateTimeTrigger(tp.UUID, tp.TimeTrigger) (tp.TimeTrigger, error)

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
	CreateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(tp.UUID) error
	ListUsageTriggers() ([]tp.UsageTrigger, error)
	GetUsageTrigger(tp.UUID) (tp.UsageTrigger, error)
	UpdateUsageTrigger(tp.UUID, tp.UsageTrigger) (tp.UsageTrigger, error)

	// usage unit
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	ListUsageUnits() ([]tp.UsageUnit, error)
	GetUsageUnit(string) (tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)

	// work order
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(tp.UUID) error
	ListWorkOrders() ([]tp.WorkOrder, error)
	GetWorkOrder(tp.UUID) (tp.WorkOrder, error)
	UpdateWorkOrder(tp.UUID, tp.WorkOrder) (tp.WorkOrder, error)

	// work order status
	CreateWorkOrderStatus(tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(string) error
	ListWorkOrderStatuses() ([]tp.WorkOrderStatus, error)
	GetWorkOrderStatus(string) (tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(string, tp.WorkOrderStatus) (tp.WorkOrderStatus, error)

	// utilities
	Exec(string) error
	Close() error
}

func New() Store {
	return imp.New()
}

// used for testing
func NewWithDb(dbName string) Store {
	return imp.NewWithDb(dbName)
}
