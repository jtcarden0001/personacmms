package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) AssociateAssetWithCategory(AssetId string, CategoryId string) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "AssociateAssetWithCategory not implemented")
}

func (a *App) AssociateAssetWithGroup(AssetId string, GroupId string) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "AssociateAssetWithGroup not implemented")
}

func (a *App) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	return tp.Asset{}, ae.New(ae.CodeNotImplemented, "CreateAsset not implemented")
}

func (a *App) DeleteAsset(assetId string) error {
	return ae.New(ae.CodeNotImplemented, "DeleteAsset not implemented")
}

func (a *App) DisassociateAssetWithCategory(AssetId string, CategoryId string) error {
	return ae.New(ae.CodeNotImplemented, "DisassociateAssetWithCategory not implemented")
}

func (a *App) DisassociateAssetWithGroup(AssetId string, GroupId string) error {
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
