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
	AssociateAssetWithCategory(uuid.UUID, uuid.UUID) (tp.Asset, error) //TODO
	AssociateAssetWithGroup(uuid.UUID, uuid.UUID) (tp.Asset, error)    //TODO
	CreateAsset(tp.Asset) (tp.Asset, error)
	DeleteAsset(uuid.UUID) error
	DisassociateAssetWithCategory(uuid.UUID, uuid.UUID) error //TODO
	DisassociateAssetWithGroup(uuid.UUID, uuid.UUID) error    //TODO
	GetAsset(uuid.UUID) (tp.Asset, error)
	ListAssets() ([]tp.Asset, error)
	ListAssetsByCategory(uuid.UUID) ([]tp.Asset, error)                    //TODO
	ListAssetsByCategoryAndGroup(uuid.UUID, uuid.UUID) ([]tp.Asset, error) //TODO
	ListAssetsByGroup(uuid.UUID) ([]tp.Asset, error)                       //TODO
	UpdateAsset(asset tp.Asset) (tp.Asset, error)

	// category
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(uuid.UUID) error
	GetCategory(uuid.UUID) (tp.Category, error)
	ListCategories() ([]tp.Category, error)
	ListCategoriesByAsset(uuid.UUID) ([]tp.Category, error) //TODO
	UpdateCategory(tp.Category) (tp.Category, error)

	// consumable
	AssociateConsumableWithTask(uuid.UUID, uuid.UUID) (tp.Consumable, error)      //TODO
	AssociateConsumableWithWorkOrder(uuid.UUID, uuid.UUID) (tp.Consumable, error) //TODO
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(uuid.UUID) error
	DisassociateConsumableWithTask(uuid.UUID, uuid.UUID) error      //TODO
	DisassociateConsumableWithWorkOrder(uuid.UUID, uuid.UUID) error //TODO
	GetConsumable(uuid.UUID) (tp.Consumable, error)
	ListConsumables() ([]tp.Consumable, error)
	ListConsumablesByTask(uuid.UUID) ([]tp.Consumable, error)      //TODO
	ListConsumablesByWorkOrder(uuid.UUID) ([]tp.Consumable, error) //TODO
	UpdateConsumable(tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(uuid.UUID) error
	DeleteDateTriggerFromTask(uuid.UUID, uuid.UUID) error //TODO
	GetDateTrigger(uuid.UUID) (tp.DateTrigger, error)
	ListDateTriggers() ([]tp.DateTrigger, error)
	ListDateTriggersByTask(uuid.UUID) ([]tp.DateTrigger, error) //TODO
	UpdateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(uuid.UUID) error
	GetGroup(uuid.UUID) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	ListGroupsByAsset(uuid.UUID) ([]tp.Group, error) //TODO
	UpdateGroup(tp.Group) (tp.Group, error)

	// task
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(uuid.UUID) error
	DeleteTaskFromAsset(uuid.UUID, uuid.UUID) error //TODO
	GetTask(uuid.UUID) (tp.Task, error)
	ListTasks() ([]tp.Task, error)
	ListTasksByAsset(uuid.UUID) ([]tp.Task, error) //TODO
	UpdateTask(tp.Task) (tp.Task, error)

	// time trigger
	CreateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(uuid.UUID) error
	DeleteTimeTriggerFromTask(uuid.UUID, uuid.UUID) error //TODO
	GetTimeTrigger(uuid.UUID) (tp.TimeTrigger, error)
	ListTimeTriggers() ([]tp.TimeTrigger, error)
	ListTimeTriggersByTask(uuid.UUID) ([]tp.TimeTrigger, error) //TODO
	UpdateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)

	// tool
	AssociateToolWithTask(uuid.UUID, uuid.UUID) (tp.Tool, error)      //TODO
	AssociateToolWithWorkOrder(uuid.UUID, uuid.UUID) (tp.Tool, error) //TODO
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(uuid.UUID) error
	DisassociateToolWithTask(uuid.UUID, uuid.UUID) error      //TODO
	DisassociateToolWithWorkOrder(uuid.UUID, uuid.UUID) error //TODO
	GetTool(uuid.UUID) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	ListToolsByTask(uuid.UUID) ([]tp.Tool, error)      //TODO
	ListToolsByWorkOrder(uuid.UUID) ([]tp.Tool, error) //TODO
	UpdateTool(tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(uuid.UUID) error
	DeleteUsageTriggerFromTask(uuid.UUID, uuid.UUID) error
	GetUsageTrigger(uuid.UUID) (tp.UsageTrigger, error)
	ListUsageTriggers() ([]tp.UsageTrigger, error)
	ListUsageTriggersByTask(uuid.UUID) ([]tp.UsageTrigger, error)
	UpdateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)

	// work order
	AssociateWorkOrderWithTask(uuid.UUID, uuid.UUID) (tp.WorkOrder, error) // TODO
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(uuid.UUID) error
	DeleteWorkOrderFromAsset(uuid.UUID, uuid.UUID) error      // TODO
	DisassociateWorkOrderWithTask(uuid.UUID, uuid.UUID) error // TODO
	GetWorkOrder(uuid.UUID) (tp.WorkOrder, error)
	ListWorkOrders() ([]tp.WorkOrder, error)
	ListWorkOrdersByAsset(uuid.UUID) ([]tp.WorkOrder, error) // TODO
	UpdateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)

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
