package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type TimeUnit interface {
	CreateTimeUnit(string) (int, error)
	DeleteTimeUnit(int) error
	GetAllTimeUnit() ([]tp.TimeUnit, error)
	GetTimeUnit(int) (tp.TimeUnit, error)
	UpdateTimeUnit(int, string) error
}

func (a *App) CreateTimeUnit(title string) (int, error) {
	return a.db.CreateTimeUnit(title)
}

func (a *App) DeleteTimeUnit(id int) error {
	return a.db.DeleteTimeUnit(id)
}

func (a *App) GetAllTimeUnit() ([]tp.TimeUnit, error) {
	return a.db.GetAllTimeUnit()
}

func (a *App) GetTimeUnit(id int) (tp.TimeUnit, error) {
	return a.db.GetTimeUnit(id)
}

func (a *App) UpdateTimeUnit(id int, title string) error {
	return a.db.UpdateTimeUnit(id, title)
}
