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
	ListDateTriggersByAssetAndTask(string, string) ([]tp.DateTrigger, error)
	UpdateDateTrigger(string, string, string, tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	GetGroup(string) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	ListGroupsByAsset(string) ([]tp.Group, error)
	UpdateGroup(string, tp.Group) (tp.Group, error)

	// task
	CreateTask(string, tp.Task) (tp.Task, error)
	DeleteTask(string, string) error
	DisassociateTaskWithWorkOrder(string, string, string) error
	GetTask(string, string) (tp.Task, error)
	ListTasksByAsset(string) ([]tp.Task, error)
	UpdateTask(string, string, tp.Task) (tp.Task, error)

	// time trigger
	CreateTimeTrigger(string, string, tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(string, string, string) error
	GetTimeTrigger(string, string, string) (tp.TimeTrigger, error)
	ListTimeTriggersByAssetAndTask(string, string) ([]tp.TimeTrigger, error)
	ListTimeTriggerUnits() ([]string, error)
	UpdateTimeTrigger(string, string, string, tp.TimeTrigger) (tp.TimeTrigger, error)

	// tool
	AssociateToolWithTask(string, string, string) (tp.Tool, error)
	AssociateToolWithWorkOrder(string, string, string) (tp.Tool, error)
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	DisassociateToolWithTask(string, string, string) error
	DisassociateToolWithWorkOrder(string, string, string) error
	GetTool(string) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(string, string, tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(string, string, string) error
	GetUsageTrigger(string, string, string) (tp.UsageTrigger, error)
	ListUsageTriggersByAssetAndTask(string, string) ([]tp.UsageTrigger, error)
	ListUsageTriggerUnits() ([]string, error)
	UpdateUsageTrigger(string, string, string, tp.UsageTrigger) (tp.UsageTrigger, error)

	// work order
	AssociateWorkOrderWithTask(string, string, string) (tp.WorkOrder, error)
	CreateWorkOrder(string, tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(string, string) error
	DisassociateWorkOrderWithTask(string, string, string) error
	GetWorkOrder(string, string) (tp.WorkOrder, error)
	ListWorkOrdersByAsset(string) ([]tp.WorkOrder, error)
	ListWorkOrderStatus() ([]string, error)
	UpdateWorkOrder(string, string, tp.WorkOrder) (tp.WorkOrder, error)
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
