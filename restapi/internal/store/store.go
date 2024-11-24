package store

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// these embedded interfaces are not implementation specific but it's easier keeping related code (interface + implementation) together.
// I will move these to a common location (out of imp) if I field the idea of another implementation or find a need.
// there is probably a way to greatly reduce the code quantity and generalize the interface implementation, just making the query strings
// implementation specific but tbd....
type Store interface {
	imp.Asset
	imp.AssetTask

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

	imp.Task
	imp.TimeUnit

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

	imp.WorkOrderStatus
	imp.Exec
}

func New() Store {
	return imp.New()
}

// used for testing
func NewWithDb(dbName string) Store {
	return imp.NewWithDb(dbName)
}
