package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type Task interface {
	CreateTask(string, string, *int, *int, *int, *int, int) (int, error)
	DeleteTask(int) error
	GetAllTask() ([]tp.Task, error)
	GetAllTaskByAssetId(int) ([]tp.Task, error)
	GetTask(int) (tp.Task, error)
	UpdateTask(int, string, string, *int, *int, *int, *int, int) error
}

func (a *App) CreateTask(title string, instructions string, timePeriodicityQuant *int, timePeriodicityUnitId *int, usagePeriodicityQuant *int, usagePeriodicityUnitId *int, assetId int) (int, error) {
	return a.db.CreateTask(title, instructions, timePeriodicityQuant, timePeriodicityUnitId, usagePeriodicityQuant, usagePeriodicityUnitId, assetId)
}

func (a *App) DeleteTask(id int) error {
	return a.db.DeleteTask(id)
}

func (a *App) GetAllTask() ([]tp.Task, error) {
	return a.db.GetAllTask()
}

func (a *App) GetAllTaskByAssetId(assetId int) ([]tp.Task, error) {
	return a.db.GetAllTaskByAssetId(assetId)
}

func (a *App) GetTask(id int) (tp.Task, error) {
	return a.db.GetTask(id)
}

func (a *App) UpdateTask(id int, title string, instructions string, timePeriodicityQuant *int, timePeriodicityUnitId *int, usagePeriodicityQuant *int, usagePeriodicityUnitId *int, assetId int) error {
	return a.db.UpdateTask(id, title, instructions, timePeriodicityQuant, timePeriodicityUnitId, usagePeriodicityQuant, usagePeriodicityUnitId, assetId)
}
