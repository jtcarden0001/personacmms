package store

import imp "github.com/jtcarden0001/personacmms/restapi/internal/store/postgres"

// these embedded interfaces are not implementation specific but it's easier keeping related code (interface + implementation) together.
// I will move these to a common location (out of imp) if I field the idea of another implementation or find a need.
// there is probably a way to greatly reduce the code quantity and generalize the interface implementation, just making the query strings
// implementation specific but tbd....
type Store interface {
	imp.Asset
	imp.Category
	imp.Consumable
	imp.Group
	imp.Task
	imp.TaskConsumable
	imp.TaskTool
	imp.TimePeriodicityUnit
	imp.Tool
	imp.UsagePeriodicityUnit
	imp.WorkOrder
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
