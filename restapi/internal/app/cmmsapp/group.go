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

	err := a.validateGroup(grp)
	if err != nil {
		return tp.Group{}, errors.Wrapf(err, "CreateGroup validation failed")
	}

	return a.db.CreateGroup(grp)
}

func (a *App) DeleteGroup(grpId string) error {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	return a.db.DeleteGroup(grpUuid)
}

func (a *App) ListGroups() ([]tp.Group, error) {
	return a.db.ListGroups()
}

func (a *App) ListGroupsByAsset(assetId string) ([]tp.Group, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	return a.db.ListGroupsByAsset(assetUuid)
}

func (a *App) GetGroup(grpId string) (tp.Group, error) {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil {
		return tp.Group{}, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	return a.db.GetGroup(grpUuid)
}

func (a *App) UpdateGroup(id string, newGroup tp.Group) (tp.Group, error) {
	grpUuid, err := uuid.Parse(id)
	if err != nil {
		return tp.Group{}, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	if newGroup.Id != uuid.Nil && newGroup.Id != grpUuid {
		return tp.Group{}, ae.New(ae.CodeInvalid, fmt.Sprintf("group id mismatch [%s] and [%s]", newGroup.Id, grpUuid))
	}

	newGroup.Id = grpUuid
	err = a.validateGroup(newGroup)
	if err != nil {
		return tp.Group{}, errors.Wrapf(err, "UpdateGroup validation failed")
	}

	return a.db.UpdateGroup(newGroup)
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

func (a *App) groupExists(grpId string) (uuid.UUID, bool, error) {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	_, err = a.db.GetGroup(grpUuid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return grpUuid, false, nil
		}
		return grpUuid, false, err
	}
	return grpUuid, true, nil
}
