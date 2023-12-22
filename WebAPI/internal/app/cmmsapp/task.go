package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type Task interface {
	CreateTask(string, string, *int, *int, *int, *int, int) (int, error)
	DeleteTask(int) error
	GetAllTask() ([]tp.Task, error)
	GetAllTaskByEquipmentId(int) ([]tp.Task, error)
	GetTask(int) (tp.Task, error)
	UpdateTask(int, string, string, *int, *int, *int, *int, int) error
}
