package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

type PreventativeTask interface {
	CreatePreventativeTask(tp.PreventativeTask) (tp.PreventativeTask, error)
	DeletePreventativeTask(string) error
	ListPreventativeTasks() ([]tp.PreventativeTask, error)
	GetPreventativeTask(string) (tp.PreventativeTask, error)
	UpdatePreventativeTask(string, tp.PreventativeTask) (tp.PreventativeTask, error)
}

func (a *App) CreatePreventativeTask(preventativeTask tp.PreventativeTask) (tp.PreventativeTask, error) {
	return a.db.CreatePreventativeTask(preventativeTask)
}

func (a *App) DeletePreventativeTask(title string) error {
	return a.db.DeletePreventativeTask(title)
}

func (a *App) ListPreventativeTasks() ([]tp.PreventativeTask, error) {
	return a.db.ListPreventativeTasks()
}

func (a *App) GetPreventativeTask(title string) (tp.PreventativeTask, error) {
	return a.db.GetPreventativeTask(title)
}

func (a *App) UpdatePreventativeTask(oldTitle string, preventativeTask tp.PreventativeTask) (tp.PreventativeTask, error) {
	return a.db.UpdatePreventativeTask(oldTitle, preventativeTask)
}
