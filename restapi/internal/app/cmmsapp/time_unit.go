package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateTimeUnit(timeUnit tp.TimeUnit) (tp.TimeUnit, error) {
	return a.db.CreateTimeUnit(timeUnit)
}

func (a *App) DeleteTimeUnit(timeUnitTitle string) error {
	return a.db.DeleteTimeUnit(timeUnitTitle)
}

func (a *App) ListTimeUnits() ([]tp.TimeUnit, error) {
	return a.db.ListTimeUnits()
}

func (a *App) GetTimeUnit(timeUnitTitle string) (tp.TimeUnit, error) {
	return a.db.GetTimeUnit(timeUnitTitle)
}

func (a *App) UpdateTimeUnit(timeUnitTitle string, timeUnit tp.TimeUnit) (tp.TimeUnit, error) {
	return a.db.UpdateTimeUnit(timeUnitTitle, timeUnit)
}
