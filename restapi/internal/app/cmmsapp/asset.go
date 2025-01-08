package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

// TODO: ensure the returned Asset hsa a list of references for the associated entities (categories and groups)

func (a *App) AssociateAssetWithCategory(assetId string, categoryId string) (apitp.AssetResponse, error) {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	stAsset, err := a.db.AssociateAssetWithCategory(aUuid, cUuid)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "AssociateAssetWithCategory failed")
	}

	return a.convertStoreAssetToApiAssetResponse(stAsset)
}

func (a *App) AssociateAssetWithGroup(assetId string, groupId string) (apitp.AssetResponse, error) {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	gUuid, err := uuid.Parse(groupId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	stAsset, err := a.db.AssociateAssetWithGroup(aUuid, gUuid)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "AssociateAssetWithGroup failed")
	}

	return a.convertStoreAssetToApiAssetResponse(stAsset)
}

func (a *App) CreateAsset(asset apitp.AssetRequest) (apitp.AssetResponse, error) {
	err := a.validateAsset(asset)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "CreateAsset validation failed")
	}

	assetId := uuid.New()
	stAssetRequest, err := a.convertApiAssetRequestToStoreAsset(assetId, asset)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "CreateAsset - convertApiAssetRequestToStoreAsset failed")
	}

	stAssetResponse, err := a.db.CreateAsset(stAssetRequest)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "CreateAsset failed")
	}

	return a.convertStoreAssetToApiAssetResponse(stAssetResponse)
}

func (a *App) DeleteAsset(assetId string) error {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	// TODO: ensure cascading deletes to delete relationships with groups/categories, tasks, work orders,
	// tool and consumable relationships, and triggers for the tasks
	return a.db.DeleteAsset(assetUuid)
}

func (a *App) DisassociateAssetWithCategory(assetId string, categoryId string) error {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "asset id must be a valid, non-nil uuid")
	}

	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "category id must be a valid, non-nil uuid")
	}

	return a.db.DisassociateAssetWithCategory(aUuid, cUuid)
}

func (a *App) DisassociateAssetWithGroup(assetId string, groupId string) error {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	gUuid, err := uuid.Parse(groupId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.DisassociateAssetWithGroup(aUuid, gUuid)
}

func (a *App) GetAsset(assetId string) (apitp.AssetResponse, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	stAsset, err := a.db.GetAsset(assetUuid)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "GetAsset failed")
	}

	return a.convertStoreAssetToApiAssetResponse(stAsset)
}

func (a *App) ListAssets() ([]apitp.AssetResponse, error) {
	stAssets, err := a.db.ListAssets()
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssets failed")
	}

	return a.convertStoreAssetListToApiAssetResponseList(stAssets)
}

func (a *App) ListAssetsByCategory(categoryId string) ([]apitp.AssetResponse, error) {
	cUuid, cFound, err := a.categoryExists(categoryId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByCategory - categoryExists failed")
	}

	if !cFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("category with id [%s] not found", categoryId))
	}

	stAssets, err := a.db.ListAssetsByCategory(cUuid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByCategory failed")
	}

	return a.convertStoreAssetListToApiAssetResponseList(stAssets)
}

func (a *App) ListAssetsByCategoryAndGroup(categoryId string, groupId string) ([]apitp.AssetResponse, error) {
	cUuid, cFound, err := a.categoryExists(categoryId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByCategoryAndGroup - categoryExists failed")
	}

	if !cFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("category with id [%s] not found", categoryId))
	}

	gUuid, gFound, err := a.groupExists(groupId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByCategoryAndGroup - groupExists failed")
	}

	if !gFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("group with id [%s] not found", groupId))
	}

	stAssets, err := a.db.ListAssetsByCategoryAndGroup(cUuid, gUuid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByCategoryAndGroup failed")
	}

	return a.convertStoreAssetListToApiAssetResponseList(stAssets)
}

func (a *App) ListAssetsByGroup(groupId string) ([]apitp.AssetResponse, error) {
	gUuid, gFound, err := a.groupExists(groupId)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByGroup - groupExists failed")
	}

	if !gFound {
		return nil, ae.New(ae.CodeNotFound, fmt.Sprintf("group with id [%s] not found", groupId))
	}

	stAssets, err := a.db.ListAssetsByGroup(gUuid)
	if err != nil {
		return nil, errors.Wrapf(err, "ListAssetsByGroup failed")
	}

	return a.convertStoreAssetListToApiAssetResponseList(stAssets)
}

func (a *App) UpdateAsset(assetId string, asset apitp.AssetRequest) (apitp.AssetResponse, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return apitp.AssetResponse{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	err = a.validateAsset(asset)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "UpdateAsset - asset validation failed")
	}

	stAsset, err := a.convertApiAssetRequestToStoreAsset(assetUuid, asset)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "UpdateAsset - convertApiAssetRequestToStoreAsset failed")
	}

	stAsset, err = a.db.UpdateAsset(stAsset)
	if err != nil {
		return apitp.AssetResponse{}, errors.Wrapf(err, "UpdateAsset failed")
	}

	return a.convertStoreAssetToApiAssetResponse(stAsset)
}

func (a *App) validateAsset(asset apitp.AssetRequest) error {
	if len(asset.Title) < apitp.MinEntityTitleLength || len(asset.Title) > apitp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("asset title length must be between [%d] and [%d] characters",
				apitp.MinEntityTitleLength,
				apitp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) assetExists(assetId string) (uuid.UUID, bool, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil || assetUuid == uuid.Nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "asset id must be a valid and not nil uuid")
	}

	_, err = a.db.GetAsset(assetUuid)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return assetUuid, false, nil
		}

		return assetUuid, false, errors.Wrapf(err, "assetExists failed")
	}

	return assetUuid, true, nil
}

func (a *App) convertApiAssetRequestToStoreAsset(Id uuid.UUID, assetRequest apitp.AssetRequest) (storetp.Asset, error) {
	return storetp.Asset{
		Id:           Id,
		Title:        assetRequest.Title,
		Year:         assetRequest.Year,
		Manufacturer: assetRequest.Manufacturer,
		Make:         assetRequest.Make,
		ModelNumber:  assetRequest.ModelNumber,
		SerialNumber: assetRequest.SerialNumber,
		Description:  assetRequest.Description,
	}, nil
}

func (a *App) convertStoreAssetListToApiAssetResponseList(storeAssets []storetp.Asset) ([]apitp.AssetResponse, error) {
	apiAssets := make([]apitp.AssetResponse, len(storeAssets))
	for i, storeAsset := range storeAssets {
		apiAsset, err := a.convertStoreAssetToApiAssetResponse(storeAsset)
		if err != nil {
			return nil, errors.Wrapf(err, "convertStoreAssetListToApiAssetResponseList failed")
		}

		apiAssets[i] = apiAsset
	}

	return apiAssets, nil
}

func (a *App) convertStoreAssetToApiAssetResponse(storeAsset storetp.Asset) (apitp.AssetResponse, error) {
	apiAsset := apitp.AssetResponse{
		Id:           storeAsset.Id,
		Title:        storeAsset.Title,
		Year:         storeAsset.Year,
		Manufacturer: storeAsset.Manufacturer,
		Make:         storeAsset.Make,
		ModelNumber:  storeAsset.ModelNumber,
		SerialNumber: storeAsset.SerialNumber,
		Description:  storeAsset.Description,
	}

	return apiAsset, nil
}
