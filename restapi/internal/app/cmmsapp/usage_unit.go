package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type UsageUnit interface {
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	ListUsageUnits() ([]tp.UsageUnit, error)
	GetUsageUnit(string) (tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)
}

func (a *App) CreateUsageUnit(usageUnit tp.UsageUnit) (tp.UsageUnit, error) {
	return a.db.CreateUsageUnit(usageUnit)
}

func (a *App) DeleteUsageUnit(usageUnitTitle string) error {
	return a.db.DeleteUsageUnit(usageUnitTitle)
}

func (a *App) ListUsageUnits() ([]tp.UsageUnit, error) {
	return a.db.ListUsageUnits()
}

func (a *App) GetUsageUnit(usageUnitTitle string) (tp.UsageUnit, error) {
	return a.db.GetUsageUnit(usageUnitTitle)
}

func (a *App) UpdateUsageUnit(usageUnitTitle string, usageUnit tp.UsageUnit) (tp.UsageUnit, error) {
	return a.db.UpdateUsageUnit(usageUnitTitle, usageUnit)
}
