package cmmsapp

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

func (a *App) AssociateAssetWithCategory(assetId string, categoryId string) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "AssociateAssetWithCategory not implemented")
}

func (a *App) AssociateAssetWithGroup(assetId string, groupId string) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "AssociateAssetWithGroup not implemented")
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
	return ae.New(ae.CodeNotImplemented, "DeleteAsset not implemented")
}

func (a *App) DisassociateAssetWithCategory(assetId string, categoryId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateAssetWithCategory not implemented")
}

func (a *App) DisassociateAssetWithGroup(assetId string, groupId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateAssetWithGroup not implemented")
}

func (a *App) GetAsset(assetId string) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "GetAsset not implemented")
}

func (a *App) ListAssets() ([]tp.Asset, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListAssets not implemented")
}

func (a *App) ListAssetsByCategory(categoryId string) ([]tp.Asset, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListAssetsByCategory not implemented")
}

func (a *App) ListAssetsByCategoryAndGroup(categoryId string, groupId string) ([]tp.Asset, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListAssetsByCategoryAndGroup not implemented")
}

func (a *App) ListAssetsByGroup(groupId string) ([]tp.Asset, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListAssetsByGroup not implemented")
}

func (a *App) UpdateAsset(assetId string, asset tp.Asset) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "UpdateAsset not implemented")
}

func (a *App) validateAsset(asset tp.Asset) error {
	if asset.Id == uuid.Nil {
		return ae.New(ae.CodeInvalid, "asset id must not be nil")
	}

	// 255 is an arbitrary number subject to change, cannot be empty though
	if len(asset.Title) < tp.MinEntityTitleLength || len(asset.Title) > tp.MaxEntityTitleLength {
		return ae.New(ae.CodeInvalid,
			fmt.Sprintf("asset title length must be between %d and %d characters",
				tp.MinEntityTitleLength,
				tp.MaxEntityTitleLength))
	}

	return nil
}

func (a *App) assetExists(assetId uuid.UUID) (bool, error) {
	_, err := a.db.GetAsset(assetId)
	if err != nil {
		var appErr ae.AppError
		if errors.As(err, &appErr); appErr.Code == ae.CodeNotFound {
			return false, nil
		}

		return false, errors.Wrapf(err, "assetExists failed")
	}

	return true, nil
}
