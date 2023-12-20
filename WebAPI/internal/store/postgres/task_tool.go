package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type TaskTool interface {
	CreateTaskTool(int, int) error
	DeleteTaskTool(int, int) error
	GetAllTaskTool() ([]tp.TaskTool, error)
	GetAllTaskToolByTaskId(int) ([]tp.TaskTool, error)
	GetAllTaskToolByToolId(int) ([]tp.TaskTool, error)
	// GetTaskTool(int, int) (tp.TaskTool, error)
	// UpdateTaskTool(int, int) error
}

func (pg *Store) CreateTaskTool(taskId int, toolId int) error {
	return errors.New("not implemented")
}

func (pg *Store) DeleteTaskTool(taskId int, toolId int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllTaskTool() ([]tp.TaskTool, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllTaskToolByTaskId(taskId int) ([]tp.TaskTool, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetAllTaskToolByToolId(toolId int) ([]tp.TaskTool, error) {
	return nil, errors.New("not implemented")
}
