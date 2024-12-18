package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// Create a group
func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	return a.db.CreateGroup(grp)
}

// Delete a group
func (a *App) DeleteGroup(title string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetGroup(title); err != nil {
		return err
	}

	return a.db.DeleteGroup(title)
}

// List all groups
func (a *App) ListGroups() ([]tp.Group, error) {
	return a.db.ListGroups()
}

// Get a group
func (a *App) GetGroup(title string) (tp.Group, error) {
	return a.db.GetGroup(title)
}

// Update a group
func (a *App) UpdateGroup(oldTitle string, newGroup tp.Group) (tp.Group, error) {
	if newGroup.Title == "" {
		return tp.Group{}, ae.ErrGroupTitleRequired
	}

	return a.db.UpdateGroup(oldTitle, newGroup)
}
