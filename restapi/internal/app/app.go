package app

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// App layer hosts the business logic and forwards simple requests to the Store layer.
type App interface {
	// asset
	AssociateAssetWithCategory(string, string) (tp.Asset, error)
	AssociateAssetWithGroup(string, string) (tp.Asset, error)
	CreateAsset(tp.Asset) (tp.Asset, error)
	DeleteAsset(string) error
	DisassociateAssetWithCategory(string, string) error
	DisassociateAssetWithGroup(string, string) error
	GetAsset(string) (tp.Asset, error)
	ListAssets() ([]tp.Asset, error)
	ListAssetsByCategory(string) ([]tp.Asset, error)
	ListAssetsByCategoryAndGroup(string, string) ([]tp.Asset, error)
	ListAssetsByGroup(string) ([]tp.Asset, error)
	UpdateAsset(string, tp.Asset) (tp.Asset, error)

	// category
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(string) error
	GetCategory(string) (tp.Category, error)
	ListCategories() ([]tp.Category, error)
	ListCategoriesByAsset(string) ([]tp.Category, error)
	UpdateCategory(string, tp.Category) (tp.Category, error)

	// consumable
	AssociateConsumableWithTask(string, string, string) (tp.Consumable, error)
	AssociateConsumableWithWorkOrder(string, string, string) (tp.Consumable, error)
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(string) error
	DisassociateConsumableWithTask(string, string, string) error
	DisassociateConsumableWithWorkOrder(string, string, string) error
	GetConsumable(string) (tp.Consumable, error)
	ListConsumables() ([]tp.Consumable, error)
	UpdateConsumable(string, tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(string, string, tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(string, string, string) error
	GetDateTrigger(string, string, string) (tp.DateTrigger, error)
	ListDateTriggers(string, string) ([]tp.DateTrigger, error)
	UpdateDateTrigger(string, string, string, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	GetGroup(string) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task
	CreateTask(string, string, tp.Task) (tp.Task, error)
	DeleteTask(string, string, string) error
	GetTask(string, string, string) (tp.Task, error)
	ListTasks(string, string) ([]tp.Task, error)
	UpdateTask(string, string, string, tp.Task) (tp.Task, error)

	// task consumable
	CreateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)
	CreateTaskConsumableWithValidation(string, string, string, string, string) (tp.TaskConsumable, error)
	DeleteTaskConsumable(string, string, string, string) error
	GetTaskConsumable(string, string, string, string) (tp.TaskConsumable, error)
	ListTaskConsumables(string, string, string) ([]tp.TaskConsumable, error)
	UpdateTaskConsumable(tp.TaskConsumable) (tp.TaskConsumable, error)
	UpdateTaskConsumableWithValidation(string, string, string, string, string) (tp.TaskConsumable, error)

	// task template
	CreateTaskTemplate(tp.TaskTemplate) (tp.TaskTemplate, error)
	DeleteTaskTemplate(string) error
	GetTaskTemplate(string) (tp.TaskTemplate, error)
	ListTaskTemplates() ([]tp.TaskTemplate, error)
	UpdateTaskTemplate(string, tp.TaskTemplate) (tp.TaskTemplate, error)

	// task tools
	CreateTaskTool(tp.TaskTool) (tp.TaskTool, error)
	CreateTaskToolWithValidation(string, string, string, string) (tp.TaskTool, error)
	DeleteTaskTool(string, string, string, string) error
	GetTaskTool(string, string, string, string) (tp.TaskTool, error)
	ListTaskTools(string, string, string) ([]tp.TaskTool, error)
	UpdateTaskTool(tp.TaskTool) (tp.TaskTool, error)

	// time trigger
	CreateTimeTrigger(string, string, string, tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(string, string, string, string) error
	GetTimeTrigger(string, string, string, string) (tp.TimeTrigger, error)
	ListTimeTriggers(string, string, string) ([]tp.TimeTrigger, error)
	UpdateTimeTrigger(string, string, string, string, tp.TimeTrigger) (tp.TimeTrigger, error)

	// time unit
	CreateTimeUnit(tp.TimeUnit) (tp.TimeUnit, error)
	DeleteTimeUnit(string) error
	GetTimeUnit(string) (tp.TimeUnit, error)
	ListTimeUnits() ([]tp.TimeUnit, error)
	UpdateTimeUnit(string, tp.TimeUnit) (tp.TimeUnit, error)

	// tool
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	GetTool(string) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(string, string, string, tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(string, string, string, string) error
	GetUsageTrigger(string, string, string, string) (tp.UsageTrigger, error)
	ListUsageTriggers(string, string, string) ([]tp.UsageTrigger, error)
	UpdateUsageTrigger(string, string, string, string, tp.UsageTrigger) (tp.UsageTrigger, error)

	// usage unit
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	GetUsageUnit(string) (tp.UsageUnit, error)
	ListUsageUnits() ([]tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)

	// work order status
	CreateWorkOrderStatus(wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)
	DeleteWorkOrderStatus(title string) error
	GetWorkOrderStatus(title string) (tp.WorkOrderStatus, error)
	ListWorkOrderStatus() ([]tp.WorkOrderStatus, error)
	UpdateWorkOrderStatus(title string, wos tp.WorkOrderStatus) (tp.WorkOrderStatus, error)

	// work order
	CreateWorkOrder(string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(string, string, string, string) error
	GetWorkOrder(string, string, string, string) (tp.WorkOrder, error)
	ListWorkOrders(string, string, string) ([]tp.WorkOrder, error)
	UpdateWorkOrder(string, string, string, string, tp.WorkOrder) (tp.WorkOrder, error)
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
