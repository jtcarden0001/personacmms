package app

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/app/cmmsapp"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
)

// App layer hosts the business logic (To be implemented) and forwards requests to the Store layer.
type App interface {
	imp.Asset
	imp.Category
	imp.Consumable
	imp.Group
	imp.TaskConsumable
	imp.TaskTool
	imp.Task
	imp.TimeUnit
	imp.Tool
	imp.UsageUnit
	imp.WorkOrderStatus
	imp.WorkOrder
}

type AppTest interface {
	App
}

func New(injectedStore st.Store) App {
	return imp.New(injectedStore)
}
