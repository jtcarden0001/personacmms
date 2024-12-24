package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "CreateGroup not implemented")
}

func (a *App) DeleteGroup(grpId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteGroup not implemented")
}

func (a *App) ListGroups() ([]tp.Group, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListGroups not implemented")
}

func (a *App) ListGroupsByAsset(assetId string) ([]tp.Group, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListGroupsByAsset not implemented")
}

func (a *App) GetGroup(grpId string) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "GetGroup not implemented")
}

func (a *App) UpdateGroup(id string, newGroup tp.Group) (tp.Group, error) {
	return tp.Group{}, ae.New(ae.CodeNotImplemented, "UpdateGroup not implemented")
}

func (a *App) validateGroup(grp tp.Group) error {
	return ae.New(ae.CodeNotImplemented, "validateGroup not implemented")
}

func (a *App) groupExists(grpId uuid.UUID) (bool, error) {
	return false, ae.New(ae.CodeNotImplemented, "groupExists not implemented")
}
