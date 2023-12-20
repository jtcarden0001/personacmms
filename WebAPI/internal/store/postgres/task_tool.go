package postgres

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type TaskTool interface {
	CreateTaskTool(int, int) error
	DeleteTaskTool(int, int) error
	GetAllTaskTool() ([]tp.TaskTool, error)
	GetAllTaskToolByTaskId(int) ([]tp.TaskTool, error)
	GetAllTaskToolByToolId(int) ([]tp.TaskTool, error)
	// GetTaskTool(int, int) (tp.TaskTool, error)
	// UpdateTaskTool(int, int) error
}
