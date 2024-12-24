package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateTaskTemplate(taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	return tp.TaskTemplate{}, ae.New(ae.CodeNotImplemented, "CreateTaskTemplate not implemented")
}

func (a *App) DeleteTaskTemplate(title string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteTaskTemplate not implemented")
}

func (a *App) GetTaskTemplate(title string) (tp.TaskTemplate, error) {
	return tp.TaskTemplate{}, ae.New(ae.CodeNotImplemented, "GetTaskTemplate not implemented")
}

func (a *App) ListTaskTemplates() ([]tp.TaskTemplate, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTaskTemplates not implemented")
}

func (a *App) UpdateTaskTemplate(oldTitle string, taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	return tp.TaskTemplate{}, ae.New(ae.CodeNotImplemented, "UpdateTaskTemplate not implemented")
}
