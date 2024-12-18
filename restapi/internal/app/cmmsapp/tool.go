package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// Create a tool
func (a *App) CreateTool(tool tp.Tool) (tp.Tool, error) {
	if err := a.validateTool(tool); err != nil {
		return tp.Tool{}, err
	}

	return a.db.CreateTool(tool)
}

// Delete a tool
func (a *App) DeleteTool(title string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetTool(title); err != nil {
		return err
	}

	return a.db.DeleteTool(title)
}

// List all tools
func (a *App) ListTools() ([]tp.Tool, error) {
	return a.db.ListTools()
}

// Get a tool
func (a *App) GetTool(title string) (tp.Tool, error) {
	return a.db.GetTool(title)
}

// Update a tool
func (a *App) UpdateTool(oldTitle string, tool tp.Tool) (tp.Tool, error) {
	if err := a.validateTool(tool); err != nil {
		return tp.Tool{}, err
	}

	return a.db.UpdateTool(oldTitle, tool)
}

func (a *App) validateTool(tool tp.Tool) error {
	if tool.Title == "" {
		return ae.ErrToolTitleRequired
	}

	return nil
}
