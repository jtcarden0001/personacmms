package store

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
	"github.com/jtcarden0001/personacmms/restapi/internal/store/test/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// Store layer implements persistence for the application.
type Store interface {
	// asset
	CreateAsset(tp.Asset) (tp.Asset, error)
	DeleteAsset(string, string) error
	ListAssets() ([]tp.Asset, error)
	ListAssetsByGroup(string) ([]tp.Asset, error)
	GetAsset(string, string) (tp.Asset, error)
	UpdateAsset(string, string, tp.Asset) (tp.Asset, error)

	// asset task consumable
	CreateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)
	DeleteTaskConsumable(tp.UUID, tp.UUID) error
	ListTaskConsumables() ([]tp.TaskConsumable, error)
	GetTaskConsumable(tp.UUID, tp.UUID) (tp.TaskConsumable, error)
	UpdateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)

	// asset task tool
	CreateTaskTool(tp.TaskTool) (tp.TaskTool, error)
	DeleteTaskTool(tp.UUID, tp.UUID) error
	ListTaskTools() ([]tp.TaskTool, error)
	GetTaskTool(tp.UUID, tp.UUID) (tp.TaskTool, error)

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
	ListDateTriggersByTaskId(tp.UUID) ([]tp.DateTrigger, error)
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
	DeleteTask(tp.UUID) error
	ListTasks() ([]tp.Task, error)
	GetTask(tp.UUID) (tp.Task, error)
	GetTaskByAssetId(tp.UUID, tp.UUID) (tp.Task, error)
	UpdateTask(tp.UUID, tp.Task) (tp.Task, error)

	// task template
	CreateTaskTemplate(tp.TaskTemplate) (tp.TaskTemplate, error)
	DeleteTaskTemplate(string) error
	ListTaskTemplates() ([]tp.TaskTemplate, error)
	GetTaskTemplate(string) (tp.TaskTemplate, error)
	UpdateTaskTemplate(string, tp.TaskTemplate) (tp.TaskTemplate, error)

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
	ListTimeTriggersByTaskId(tp.UUID) ([]tp.TimeTrigger, error)
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
	ListUsageTriggersByTaskId(tp.UUID) ([]tp.UsageTrigger, error)
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

func NewMock() Store {
	return mock.New()
}
