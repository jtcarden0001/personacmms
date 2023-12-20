package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type UsagePeriodicityUnit interface {
	CreateUsagePeriodicityUnit(string) (int, error)
	DeleteUsagePeriodicityUnit(int) error
	GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error)
	GetUsagePeriodicityUnit(int) (tp.UsagePeriodicityUnit, error)
	UpdateUsagePeriodicityUnit(int, string) error
}

type UsagePeriodicityUnitTest interface {
	ResetSequenceUsagePeriodicityUnit(int) error
}
