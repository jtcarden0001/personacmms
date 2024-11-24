package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateTask(task tp.Task) (tp.Task, error) {
	return a.db.CreateTask(task)
}

func (a *App) DeleteTask(title string) error {
	return a.db.DeleteTask(title)
}

func (a *App) ListTasks() ([]tp.Task, error) {
	return a.db.ListTasks()
}

func (a *App) GetTask(title string) (tp.Task, error) {
	return a.db.GetTask(title)
}

func (a *App) UpdateTask(oldTitle string, task tp.Task) (tp.Task, error) {
	return a.db.UpdateTask(oldTitle, task)
}
