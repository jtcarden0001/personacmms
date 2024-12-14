package store

import (
	"github.com/google/uuid"
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
	"github.com/jtcarden0001/personacmms/restapi/internal/store/test/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// Store layer implements persistence for the application.
type Store interface {
	// asset
	CreateAsset(tp.Asset) (tp.Asset, error)
	DeleteAsset(string, string) error
	GetAsset(string, string) (tp.Asset, error)
	ListAssets() ([]tp.Asset, error)
	ListAssetsByGroup(string) ([]tp.Asset, error)
	UpdateAsset(string, string, tp.Asset) (tp.Asset, error)

	// category
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(string) error
	GetCategory(string) (tp.Category, error)
	ListCategories() ([]tp.Category, error)
	UpdateCategory(string, tp.Category) (tp.Category, error)

	// consumable
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(string) error
	GetConsumableById(uuid.UUID) (tp.Consumable, error)
	GetConsumableByTitle(string) (tp.Consumable, error)
	ListConsumables() ([]tp.Consumable, error)
	UpdateConsumable(string, tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(tp.UUID) error
	GetDateTrigger(tp.UUID) (tp.DateTrigger, error)
	ListDateTriggers() ([]tp.DateTrigger, error)
	ListDateTriggersByTaskId(tp.UUID) ([]tp.DateTrigger, error)
	UpdateDateTrigger(tp.UUID, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	GetGroup(string) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(tp.UUID) error
	GetTask(tp.UUID) (tp.Task, error)
	GetTaskByAssetId(tp.UUID, tp.UUID) (tp.Task, error)
	ListTasks() ([]tp.Task, error)
	ListTasksByAssetId(tp.UUID) ([]tp.Task, error)
	UpdateTask(tp.UUID, tp.Task) (tp.Task, error)

	// task consumable
	CreateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)
	DeleteTaskConsumable(tp.UUID, tp.UUID) error
	GetTaskConsumable(tp.UUID, tp.UUID) (tp.TaskConsumable, error)
	ListTaskConsumables() ([]tp.TaskConsumable, error)
	ListTaskConsumablesByTaskId(tp.UUID) ([]tp.TaskConsumable, error)
	UpdateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)

	// task template
	CreateTaskTemplate(tp.TaskTemplate) (tp.TaskTemplate, error)
	DeleteTaskTemplate(string) error
	GetTaskTemplate(string) (tp.TaskTemplate, error)
	ListTaskTemplates() ([]tp.TaskTemplate, error)
	UpdateTaskTemplate(string, tp.TaskTemplate) (tp.TaskTemplate, error)

	// task tool
	CreateTaskTool(tp.TaskTool) (tp.TaskTool, error)
	DeleteTaskTool(tp.UUID, tp.UUID) error
	GetTaskTool(tp.UUID, tp.UUID) (tp.TaskTool, error)
	ListTaskTools() ([]tp.TaskTool, error)
	ListTaskToolsByTaskId(tp.UUID) ([]tp.TaskTool, error)

	// time trigger
	CreateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(tp.UUID) error
	GetTimeTrigger(tp.UUID) (tp.TimeTrigger, error)
	ListTimeTriggers() ([]tp.TimeTrigger, error)
	UpdateTimeTrigger(tp.UUID, tp.TimeTrigger) (tp.TimeTrigger, error)

	// time unit
	CreateTimeUnit(tp.TimeUnit) (tp.TimeUnit, error)
	DeleteTimeUnit(string) error
	GetTimeUnit(string) (tp.TimeUnit, error)
	ListTimeUnits() ([]tp.TimeUnit, error)
	ListTimeTriggersByTaskId(tp.UUID) ([]tp.TimeTrigger, error)
	UpdateTimeUnit(string, tp.TimeUnit) (tp.TimeUnit, error)

	// tool
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	GetTool(string) (tp.Tool, error)
	GetToolById(uuid.UUID) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(tp.UUID) error
	GetUsageTrigger(tp.UUID) (tp.UsageTrigger, error)
	ListUsageTriggers() ([]tp.UsageTrigger, error)
	ListUsageTriggersByTaskId(tp.UUID) ([]tp.UsageTrigger, error)
	UpdateUsageTrigger(tp.UUID, tp.UsageTrigger) (tp.UsageTrigger, error)

	// usage unit
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	GetUsageUnit(string) (tp.UsageUnit, error)
	ListUsageUnits() ([]tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)

	// work order
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(tp.UUID) error
	GetWorkOrder(tp.UUID) (tp.WorkOrder, error)
	ListWorkOrders() ([]tp.WorkOrder, error)
	UpdateWorkOrder(tp.UUID, tp.WorkOrder) (tp.WorkOrder, error)

	// work order status
	CreateWorkOrderStatus(tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(string) error
	GetWorkOrderStatus(string) (tp.WorkOrderStatus, error)
	ListWorkOrderStatuses() ([]tp.WorkOrderStatus, error)
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
