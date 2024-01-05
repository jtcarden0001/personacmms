package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type UsagePeriodicityUnit interface {
	CreateUsagePeriodicityUnit(string) (int, error)
	DeleteUsagePeriodicityUnit(int) error
	GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error)
	GetUsagePeriodicityUnit(int) (tp.UsagePeriodicityUnit, error)
	UpdateUsagePeriodicityUnit(int, string) error
}

func (a *App) CreateUsagePeriodicityUnit(title string) (int, error) {
	return a.db.CreateUsagePeriodicityUnit(title)
}

func (a *App) DeleteUsagePeriodicityUnit(id int) error {
	return a.db.DeleteUsagePeriodicityUnit(id)
}

func (a *App) GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error) {
	return a.db.GetAllUsagePeriodicityUnit()
}

func (a *App) GetUsagePeriodicityUnit(id int) (tp.UsagePeriodicityUnit, error) {
	return a.db.GetUsagePeriodicityUnit(id)
}

func (a *App) UpdateUsagePeriodicityUnit(id int, title string) error {
	return a.db.UpdateUsagePeriodicityUnit(id, title)
}
