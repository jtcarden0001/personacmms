package cmmsapp

import (
	"github.com/google/uuid"
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

func (a *App) validateTaskTemplate(taskTemplate tp.TaskTemplate) error {
	return ae.New(ae.CodeNotImplemented, "validateTaskTemplate not implemented")
}

func (a *App) taskTemplateExists(id uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "taskTemplateExists not implemented")
}
