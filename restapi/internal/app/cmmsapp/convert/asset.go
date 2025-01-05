package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func ConvertApiAssetRequestToStoreAsset(assetRequest apitp.AssetRequest) (storetp.Asset, error) {
	return storetp.Asset{}, ae.New(ae.CodeNotImplemented, "ConvertApiAssetRequestToStoreAsset not implemented")
}

func ConvertStoreAssetListToApiAssetResponseList(storeAssets []storetp.Asset) ([]apitp.AssetResponse, error) {
	return []apitp.AssetResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreAssetListToApiAssetResponseList not implemented")
}

func ConvertStoreAssetToApiAssetResponse(storeAsset storetp.Asset) (apitp.AssetResponse, error) {
	return apitp.AssetResponse{}, ae.New(ae.CodeNotImplemented, "convertStoreAssetToApiAssetResponse not implemented")
}
