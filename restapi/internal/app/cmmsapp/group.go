package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	return a.db.CreateGroup(grp)
}

func (a *App) DeleteGroup(title string) error {
	return a.db.DeleteGroup(title)
}

func (a *App) ListGroups() ([]tp.Group, error) {
	return a.db.ListGroups()
}

func (a *App) GetGroup(title string) (tp.Group, error) {
	return a.db.GetGroup(title)
}

func (a *App) UpdateGroup(oldTitle string, newGroup tp.Group) (tp.Group, error) {
	return a.db.UpdateGroup(oldTitle, newGroup)
}
