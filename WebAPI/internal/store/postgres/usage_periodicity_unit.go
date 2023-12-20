package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type UsagePeriodicityUnit interface {
	CreateUsagePeriodicityUnit(string) (int, error)
	DeleteUsagePeriodicityUnit(int) error
	GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error)
	GetUsagePeriodicityUnit(int) (tp.UsagePeriodicityUnit, error)
	UpdateUsagePeriodicityUnit(int, string) error
}

func (pg *Store) CreateUsagePeriodicityUnit(name string) (int, error) {
	return 0, errors.New("not implemented")
}

func (pg *Store) DeleteUsagePeriodicityUnit(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetUsagePeriodicityUnit(id int) (tp.UsagePeriodicityUnit, error) {
	return tp.UsagePeriodicityUnit{}, errors.New("not implemented")
}

func (pg *Store) UpdateUsagePeriodicityUnit(id int, name string) error {
	return errors.New("not implemented")
}
