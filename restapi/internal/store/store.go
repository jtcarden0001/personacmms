package store

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

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
}

func New() Store {
	return imp.New()
}

// used for testing
func NewWithDb(dbName string) Store {
	return imp.NewWithDb(dbName)
}
