package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type TimePeriodicityUnit interface {
	CreateTimePeriodicityUnit(string) (int, error)
	DeleteTimePeriodicityUnit(int) error
	GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error)
	GetTimePeriodicityUnit(int) (tp.TimePeriodicityUnit, error)
	UpdateTimePeriodicityUnit(int, string) error
}
