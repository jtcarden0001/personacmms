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
	imp.CorrectiveTask
	imp.Group
	imp.PreventativeTaskConsumable
	imp.PreventativeTaskTool
	imp.PreventativeTask
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
