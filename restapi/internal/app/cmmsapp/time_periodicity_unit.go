package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type TimePeriodicityUnit interface {
	CreateTimePeriodicityUnit(string) (int, error)
	DeleteTimePeriodicityUnit(int) error
	GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error)
	GetTimePeriodicityUnit(int) (tp.TimePeriodicityUnit, error)
	UpdateTimePeriodicityUnit(int, string) error
}

func (a *App) CreateTimePeriodicityUnit(title string) (int, error) {
	return a.db.CreateTimePeriodicityUnit(title)
}

func (a *App) DeleteTimePeriodicityUnit(id int) error {
	return a.db.DeleteTimePeriodicityUnit(id)
}

func (a *App) GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error) {
	return a.db.GetAllTimePeriodicityUnit()
}

func (a *App) GetTimePeriodicityUnit(id int) (tp.TimePeriodicityUnit, error) {
	return a.db.GetTimePeriodicityUnit(id)
}

func (a *App) UpdateTimePeriodicityUnit(id int, title string) error {
	return a.db.UpdateTimePeriodicityUnit(id, title)
}
