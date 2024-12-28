package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateGroup(grp tp.Group) (tp.Group, error) {
	if grp.Id != uuid.Nil {
		return tp.Group{}, ae.New(ae.CodeInvalid, "group id must be nil on create, we will create an id for you")
	}
	grp.Id = uuid.New()

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
	if grp.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "group id is required")
	}

	if len(grp.Title) < tp.MinEntityTitleLength || len(grp.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("group title length must be between [%d] and [%d] characters",
			tp.MinEntityTitleLength,
			tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) groupExists(grpId uuid.UUID) (bool, error) {
	_, err := a.db.GetGroup(grpId)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
