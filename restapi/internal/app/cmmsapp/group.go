package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "CreateGroup not implemented")
}

func (a *App) DeleteGroup(title string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteGroup not implemented")
}

func (a *App) ListGroups() ([]tp.Group, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListGroups not implemented")
}

func (a *App) ListGroupsByAsset(assetId string) ([]tp.Group, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListGroupsByAsset not implemented")
}

func (a *App) GetGroup(title string) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "GetGroup not implemented")
}

func (a *App) UpdateGroup(oldTitle string, newGroup tp.Group) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "UpdateGroup not implemented")
}
