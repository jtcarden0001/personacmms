package cmmsapp

import (
	"fmt"

	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// creates a task template
func (a *App) CreateTaskTemplate(taskTemplate tp.TaskTemplate) (tp.TaskTemplate, error) {
	err := a.validateTaskTemplate(taskTemplate)
	if err != nil {
		return tp.TaskTemplate{}, err
	}

	return a.db.CreateTaskTemplate(taskTemplate)
}

func (a *App) DeleteTaskTemplate(title string) error {
	// Get before delete so we can return a not found error.
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

func (a *App) validateTaskTemplate(tt tp.TaskTemplate) error {
	if tt.Title == "" {
		return ae.New(ae.CodeInvalid, "Task Template title is required")
	}

	if tt.Type != nil && !tp.IsValidTaskType(*tt.Type) {
		return ae.New(
			ae.CodeInvalid,
			fmt.Sprintf("Task Template type is invalid, %s", ae.CreateInvalidTaskTypeMessage()),
		)
	}

	return nil
}
