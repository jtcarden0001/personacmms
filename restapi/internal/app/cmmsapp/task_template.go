package cmmsapp

import tp "github.com/jtcarden0001/personacmms/restapi/internal/types"

func (a *App) CreateTaskTemplate(taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	return a.db.CreateTaskTemplate(taskTemplate)
}

func (a *App) DeleteTaskTemplate(title string) error {
	if _, err := a.GetTaskTemplate(title); err != nil {
		return err
	}

	return a.db.DeleteTaskTemplate(title)
}

func (a *App) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	return a.db.ListTaskTemplates()
}

func (a *App) GetTaskTemplate(title string) (tp.TaskTemplate, error) {
	return a.db.GetTaskTemplate(title)
}

func (a *App) UpdateTaskTemplate(oldTitle string, taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	return a.db.UpdateTaskTemplate(oldTitle, taskTemplate)
}
