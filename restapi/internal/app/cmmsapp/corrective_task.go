package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type CorrectiveTask interface {
	CreateCorrectiveTask(tp.CorrectiveTask) (tp.CorrectiveTask, error)
	DeleteCorrectiveTask(string) error
	ListCorrectiveTasks() ([]tp.CorrectiveTask, error)
	GetCorrectiveTask(string) (tp.CorrectiveTask, error)
	UpdateCorrectiveTask(string, tp.CorrectiveTask) (tp.CorrectiveTask, error)
}

func (a *App) CreateCorrectiveTask(correctiveTask tp.CorrectiveTask) (tp.CorrectiveTask, error) {
	return a.db.CreateCorrectiveTask(correctiveTask)
}

func (a *App) DeleteCorrectiveTask(title string) error {
	return a.db.DeleteCorrectiveTask(title)
}

func (a *App) ListCorrectiveTasks() ([]tp.CorrectiveTask, error) {
	return a.db.ListCorrectiveTasks()
}

func (a *App) GetCorrectiveTask(title string) (tp.CorrectiveTask, error) {
	return a.db.GetCorrectiveTask(title)
}

func (a *App) UpdateCorrectiveTask(oldTitle string, correctiveTask tp.CorrectiveTask) (tp.CorrectiveTask, error) {
	return a.db.UpdateCorrectiveTask(oldTitle, correctiveTask)
}
