package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Task interface {
	CreateTask(string, string, *int, *int, *int, *int) (int, error)
	DeleteTask(int) error
	GetAllTask() ([]tp.Task, error)
	GetTask(int) (tp.Task, error)
	UpdateTask(int, string, string, *int, *int, *int, *int) error
}

type TaskTest interface {
	ResetSequenceTask(int) error
}

func (pg *Store) CreateTask(title string, instructions string, timeQuant *int, timeUnit *int, usageQuant *int, usageUnit *int) (int, error) {
	return 0, errors.New("not implemented")
}

func (pg *Store) DeleteTask(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllTask() ([]tp.Task, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetTask(id int) (tp.Task, error) {
	return tp.Task{}, errors.New("not implemented")
}

func (pg *Store) UpdateTask(id int, title string, instructions string, timeQuant *int, timeUnit *int, usageQuant *int, usageUnit *int) error {
	return errors.New("not implemented")
}
