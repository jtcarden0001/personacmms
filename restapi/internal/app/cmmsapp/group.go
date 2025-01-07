package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) CreateGroup(grp apitp.GroupRequest) (apitp.GroupResponse, error) {
	if grp.Id != uuid.Nil {
		return apitp.GroupResponse{}, ae.New(ae.CodeInvalid, "group id must be nil on create, we will create an id for you")
	}
	grp.Id = uuid.New()

	err := a.validateGroup(grp)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "CreateGroup validation failed")
	}

	stGrpRequest, err := convertApiGroupRequestToStoreGroup(grp)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "CreateGroup - error converting to store type")
	}

	stGrpResponse, err := a.db.CreateGroup(stGrpRequest)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "CreateGroup - error creating group")
	}

	return convertStoreGroupToApiGroupResponse(stGrpResponse)
}

func (a *App) DeleteGroup(grpId string) error {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	// TODO: block group deletion if in use

	return a.db.DeleteGroup(grpUuid)
}

func (a *App) ListGroups() ([]apitp.GroupResponse, error) {
	stGrpResponses, err := a.db.ListGroups()
	if err != nil {
		return []apitp.GroupResponse{}, errors.Wrapf(err, "ListGroups - error listing groups")
	}

	return convertStoreGroupListToApiGroupResponseList(stGrpResponses)
}

func (a *App) ListGroupsByAsset(assetId string) ([]apitp.GroupResponse, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, printInvalidUuidErrorMessage("asset", assetId))
	}

	stGrpResponses, err := a.db.ListGroupsByAsset(assetUuid)
	if err != nil {
		return []apitp.GroupResponse{}, errors.Wrapf(err, "ListGroupsByAsset - error listing groups")
	}

	return convertStoreGroupListToApiGroupResponseList(stGrpResponses)
}

func (a *App) GetGroup(grpId string) (apitp.GroupResponse, error) {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil {
		return apitp.GroupResponse{}, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	stGrpResponse, err := a.db.GetGroup(grpUuid)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "GetGroup - error getting group")
	}

	return convertStoreGroupToApiGroupResponse(stGrpResponse)
}

func (a *App) UpdateGroup(id string, newGroup apitp.GroupRequest) (apitp.GroupResponse, error) {
	grpUuid, err := uuid.Parse(id)
	if err != nil {
		return apitp.GroupResponse{}, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	if newGroup.Id != uuid.Nil && newGroup.Id != grpUuid {
		return apitp.GroupResponse{}, ae.New(ae.CodeInvalid, fmt.Sprintf("group id mismatch [%s] and [%s]", newGroup.Id, grpUuid))
	}

	newGroup.Id = grpUuid
	err = a.validateGroup(newGroup)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "UpdateGroup validation failed")
	}

	stGrpRequest, err := convertApiGroupRequestToStoreGroup(newGroup)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "UpdateGroup - error converting to store type")
	}
	stGrpResponse, err := a.db.UpdateGroup(stGrpRequest)
	if err != nil {
		return apitp.GroupResponse{}, errors.Wrapf(err, "UpdateGroup - error updating group")
	}

	return convertStoreGroupToApiGroupResponse(stGrpResponse)
}

func (a *App) validateGroup(grp apitp.GroupRequest) error {
	if grp.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "group id is required")
	}

	if len(grp.Title) < apitp.MinEntityTitleLength || len(grp.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid, fmt.Sprintf("group title length must be between [%d] and [%d] characters",
			apitp.MinEntityTitleLength,
			apitp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) groupExists(grpId string) (uuid.UUID, bool, error) {
	grpUuid, err := uuid.Parse(grpId)
	if err != nil || grpUuid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "group id must be a valid, non-nil uuid")
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

func printInvalidUuidErrorMessage(resource, id string) string {
	return fmt.Sprintf("%s id '%s' is not a valid uuid.  Uuid must follow the format '%s'", resource, id, uuid.Nil)
}

func convertApiGroupRequestToStoreGroup(groupRequest apitp.GroupRequest) (storetp.Group, error) {
	return storetp.Group{}, ae.New(ae.CodeNotImplemented, "ConvertApiGroupRequestToStoreGroup not implemented")
}

func convertStoreGroupListToApiGroupResponseList(storeGroups []storetp.Group) ([]apitp.GroupResponse, error) {
	return []apitp.GroupResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreGroupListToApiGroupResponseList not implemented")
}

func convertStoreGroupToApiGroupResponse(storeGroup storetp.Group) (apitp.GroupResponse, error) {
	return apitp.GroupResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreGroupToApiGroupResponse not implemented")
}
