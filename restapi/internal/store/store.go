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
	DeleteAsset(uuid.UUID) error
	GetAsset(uuid.UUID) (tp.Asset, error)
	ListAssets() ([]tp.Asset, error)
	UpdateAsset(tp.Asset) (tp.Asset, error)

	// category
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(uuid.UUID) error
	GetCategory(uuid.UUID) (tp.Category, error)
	ListCategories() ([]tp.Category, error)
	UpdateCategory(tp.Category) (tp.Category, error)

	// consumable
	CreateConsumable(tp.Consumable) (tp.Consumable, error)
	DeleteConsumable(uuid.UUID) error
	GetConsumable(uuid.UUID) (tp.Consumable, error)
	ListConsumables() ([]tp.Consumable, error)
	UpdateConsumable(tp.Consumable) (tp.Consumable, error)

	// date trigger
	CreateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)
	DeleteDateTrigger(uuid.UUID) error
	GetDateTrigger(uuid.UUID) (tp.DateTrigger, error)
	ListDateTriggers() ([]tp.DateTrigger, error)
	UpdateDateTrigger(tp.DateTrigger) (tp.DateTrigger, error)

	// group
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(uuid.UUID) error
	GetGroup(uuid.UUID) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
	UpdateGroup(tp.Group) (tp.Group, error)

	// task
	CreateTask(tp.Task) (tp.Task, error)
	DeleteTask(uuid.UUID) error
	GetTask(uuid.UUID) (tp.Task, error)
	ListTasks() ([]tp.Task, error)
	UpdateTask(tp.Task) (tp.Task, error)

	// time trigger
	CreateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)
	DeleteTimeTrigger(uuid.UUID) error
	GetTimeTrigger(uuid.UUID) (tp.TimeTrigger, error)
	ListTimeTriggers() ([]tp.TimeTrigger, error)
	UpdateTimeTrigger(tp.TimeTrigger) (tp.TimeTrigger, error)

	// tool
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(uuid.UUID) error
	GetTool(uuid.UUID) (tp.Tool, error)
	ListTools() ([]tp.Tool, error)
	UpdateTool(tp.Tool) (tp.Tool, error)

	// usage trigger
	CreateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)
	DeleteUsageTrigger(uuid.UUID) error
	GetUsageTrigger(uuid.UUID) (tp.UsageTrigger, error)
	ListUsageTriggers() ([]tp.UsageTrigger, error)
	UpdateUsageTrigger(tp.UsageTrigger) (tp.UsageTrigger, error)

	// work order
	CreateWorkOrder(tp.WorkOrder) (tp.WorkOrder, error)
	DeleteWorkOrder(uuid.UUID) error
	GetWorkOrder(uuid.UUID) (tp.WorkOrder, error)
	ListWorkOrders() ([]tp.WorkOrder, error)
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
