package cmmsapp

import tp "github.com/jtcarden0001/personacmms/webapi/internal/types"

type TaskConsumable interface {
	CreateTaskConsumable(int, int, string) error
	DeleteTaskConsumable(int, int) error
	GetAllTaskConsumable() ([]tp.TaskConsumable, error)
	GetAllTaskConsumableByTaskId(int) ([]tp.TaskConsumable, error)
	GetTaskConsumable(int, int) (tp.TaskConsumable, error)
	UpdateTaskConsumable(int, int, string) error
}

func (a *App) CreateTaskConsumable(taskId int, consumableId int, quantity string) error {
	return a.db.CreateTaskConsumable(taskId, consumableId, quantity)
}

func (a *App) DeleteTaskConsumable(taskId int, consumableId int) error {
	return a.db.DeleteTaskConsumable(taskId, consumableId)
}

func (a *App) GetAllTaskConsumable() ([]tp.TaskConsumable, error) {
	return a.db.GetAllTaskConsumable()
}

func (a *App) GetAllTaskConsumableByTaskId(taskId int) ([]tp.TaskConsumable, error) {
	return a.db.GetAllTaskConsumableByTaskId(taskId)
}

func (a *App) GetTaskConsumable(taskId int, consumableId int) (tp.TaskConsumable, error) {
	return a.db.GetTaskConsumable(taskId, consumableId)
}

func (a *App) UpdateTaskConsumable(taskId int, consumableId int, quantity string) error {
	return a.db.UpdateTaskConsumable(taskId, consumableId, quantity)
}
