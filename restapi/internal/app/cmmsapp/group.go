package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Group interface {
	CreateGroup(tp.Group) (tp.Group, error)
	ListGroups() ([]tp.Group, error)
}

func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	return a.db.CreateGroup(grp)
}

func (a *App) ListGroups() ([]tp.Group, error) {
	return a.db.ListGroups()
}
