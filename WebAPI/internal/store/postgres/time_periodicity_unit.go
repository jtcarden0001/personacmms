package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type TimePeriodicityUnit interface {
	CreateTimePeriodicityUnit(string) (int, error)
	DeleteTimePeriodicityUnit(int) error
	GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error)
	GetTimePeriodicityUnit(int) (tp.TimePeriodicityUnit, error)
	UpdateTimePeriodicityUnit(int, string) error
}

type TimePeriodicityUnitTest interface {
	ResetSequenceTimePeriodicityUnit(int) error
}

func (pg *Store) CreateTimePeriodicityUnit(title string) (int, error) {
	return 0, errors.New("not implemented")
}

func (pg *Store) DeleteTimePeriodicityUnit(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetTimePeriodicityUnit(id int) (tp.TimePeriodicityUnit, error) {
	return tp.TimePeriodicityUnit{}, errors.New("not implemented")
}

func (pg *Store) UpdateTimePeriodicityUnit(id int, title string) error {
	return errors.New("not implemented")
}

func (pg *Store) ResetSequenceTimePeriodicityUnit(id int) error {
	return errors.New("not implemented")
}
