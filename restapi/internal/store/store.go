package store

import (
	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
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
	DeleteDateTrigger(uuid.UUID) error
	GetDateTrigger(uuid.UUID) (tp.DateTrigger, error)
	ListDateTriggers() ([]tp.DateTrigger, error)
	ListDateTriggersByTaskId(uuid.UUID) ([]tp.DateTrigger, error)
	UpdateDateTrigger(uuid.UUID, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	GetGroup(string) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(uuid.UUID) error
	GetTask(uuid.UUID) (tp.Task, error)
	GetTaskByAssetId(uuid.UUID, uuid.UUID) (tp.Task, error)
	ListTasks() ([]tp.Task, error)
	ListTasksByAssetId(uuid.UUID) ([]tp.Task, error)
	UpdateTask(uuid.UUID, tp.Task) (tp.Task, error)

	// task consumable
	CreateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)
	DeleteTaskConsumable(uuid.UUID, uuid.UUID) error
	GetTaskConsumable(uuid.UUID, uuid.UUID) (tp.TaskConsumable, error)
	ListTaskConsumables() ([]tp.TaskConsumable, error)
	ListTaskConsumablesByTaskId(uuid.UUID) ([]tp.TaskConsumable, error)
	UpdateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)

	// task template
	CreateTaskTemplate(tp.TaskTemplate) (tp.TaskTemplate, error)
	DeleteTaskTemplate(string) error
	GetTaskTemplate(string) (tp.TaskTemplate, error)
	ListTaskTemplates() ([]tp.TaskTemplate, error)
	UpdateTaskTemplate(string, tp.TaskTemplate) (tp.TaskTemplate, error)

	// task tool
	CreateTaskTool(tp.TaskTool) (tp.TaskTool, error)
	DeleteTaskTool(uuid.UUID, uuid.UUID) error
	GetTaskTool(uuid.UUID, uuid.UUID) (tp.TaskTool, error)
	ListTaskTools() ([]tp.TaskTool, error)
	ListTaskToolsByTaskId(uuid.UUID) ([]tp.TaskTool, error)

	// time trigger
	CreateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(uuid.UUID) error
	GetTimeTrigger(uuid.UUID) (tp.TimeTrigger, error)
	ListTimeTriggers() ([]tp.TimeTrigger, error)
	UpdateTimeTrigger(uuid.UUID, tp.TimeTrigger) (tp.TimeTrigger, error)

	// time unit
	CreateTimeUnit(tp.TimeUnit) (tp.TimeUnit, error)
	DeleteTimeUnit(string) error
	GetTimeUnit(string) (tp.TimeUnit, error)
	ListTimeUnits() ([]tp.TimeUnit, error)
	ListTimeTriggersByTaskId(uuid.UUID) ([]tp.TimeTrigger, error)
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
	DeleteUsageTrigger(uuid.UUID) error
	GetUsageTrigger(uuid.UUID) (tp.UsageTrigger, error)
	ListUsageTriggers() ([]tp.UsageTrigger, error)
	ListUsageTriggersByTaskId(uuid.UUID) ([]tp.UsageTrigger, error)
	UpdateUsageTrigger(uuid.UUID, tp.UsageTrigger) (tp.UsageTrigger, error)

	// usage unit
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	GetUsageUnit(string) (tp.UsageUnit, error)
	ListUsageUnits() ([]tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)

	// work order
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(uuid.UUID) error
	GetWorkOrder(uuid.UUID) (tp.WorkOrder, error)
	GetWorkOrderForTask(uuid.UUID, uuid.UUID) (tp.WorkOrder, error)
	ListWorkOrders() ([]tp.WorkOrder, error)
	ListWorkOrdersByTaskId(uuid.UUID) ([]tp.WorkOrder, error)
	UpdateWorkOrder(uuid.UUID, tp.WorkOrder) (tp.WorkOrder, error)

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
