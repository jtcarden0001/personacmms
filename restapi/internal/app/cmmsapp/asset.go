package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

// TODO: ensure the returned Asset hsa a list of references for the associated entities (categories and groups)

func (a *App) AssociateAssetWithCategory(assetId string, categoryId string) (tp.Asset, error) {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.AssociateAssetWithCategory(aUuid, cUuid)
}

func (a *App) AssociateAssetWithGroup(assetId string, groupId string) (tp.Asset, error) {
	aUuid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	gUuid, err := uuid.Parse(groupId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.AssociateAssetWithGroup(aUuid, gUuid)
}

func (a *App) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	if asset.Id != uuid.Nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "asset id must be nil on create, we will create an id for you")
	}
	asset.Id = uuid.New()
	err := a.validateAsset(asset)
	if err != nil {
		return tp.Asset{}, errors.Wrapf(err, "CreateAsset validation failed")
	}

	return a.db.CreateAsset(asset)
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
		return ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return ae.New(ae.CodeInvalid, "category id must be a valid uuid")
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

func (a *App) GetAsset(assetId string) (tp.Asset, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	return a.db.GetAsset(assetUuid)
}

func (a *App) ListAssets() ([]tp.Asset, error) {
	return a.db.ListAssets()
}

func (a *App) ListAssetsByCategory(categoryId string) ([]tp.Asset, error) {
	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	return a.db.ListAssetsByCategory(cUuid)
}

func (a *App) ListAssetsByCategoryAndGroup(categoryId string, groupId string) ([]tp.Asset, error) {
	cUuid, err := uuid.Parse(categoryId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "category id must be a valid uuid")
	}

	gUuid, err := uuid.Parse(groupId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	return a.db.ListAssetsByCategoryAndGroup(cUuid, gUuid)
}

func (a *App) ListAssetsByGroup(groupId string) ([]tp.Asset, error) {
	gUuid, err := uuid.Parse(groupId)
	if err != nil {
		return nil, ae.New(ae.CodeInvalid, "group id must be a valid uuid")
	}

	return a.db.ListAssetsByGroup(gUuid)
}

func (a *App) UpdateAsset(assetId string, asset tp.Asset) (tp.Asset, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return tp.Asset{}, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
	}

	if asset.Id != uuid.Nil && asset.Id != assetUuid {
		return tp.Asset{}, ae.New(ae.CodeInvalid,
			fmt.Sprintf("asset id mismatch between [%s] and [%s]",
				asset.Id.String(), assetUuid.String()))
	}

	asset.Id = assetUuid
	err = a.validateAsset(asset)
	if err != nil {
		return tp.Asset{}, errors.Wrapf(err, "UpdateAsset - asset validation failed")
	}

	return a.db.UpdateAsset(asset)
}

func (a *App) validateAsset(asset tp.Asset) error {
	if asset.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "asset id must not be nil")
	}

	// 255 is an arbitrary number subject to change, cannot be empty though
	if len(asset.Title) < tp.MinEntityTitleLength || len(asset.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("asset title length must be between [%d] and [%d] characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) assetExists(assetId string) (uuid.UUID, bool, error) {
	assetUuid, err := uuid.Parse(assetId)
	if err != nil {
		return uuid.Nil, false, ae.New(ae.CodeInvalid, "asset id must be a valid uuid")
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
