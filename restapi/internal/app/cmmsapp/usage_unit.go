package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateUsageUnit(usageUnit tp.UsageUnit) (tp.UsageUnit, error) {
	return a.db.CreateUsageUnit(usageUnit)
}

func (a *App) DeleteUsageUnit(usageUnitTitle string) error {
	if _, err := a.GetUsageUnit(usageUnitTitle); err != nil {
		return err
	}

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
