package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"

	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func ConvertApiConsumableRequestToStoreConsumable(consumableRequest apitp.ConsumableRequest) (storetp.Consumable, error) {
	return storetp.Consumable{}, ae.New(ae.CodeNotImplemented, "ConvertApiConsumableRequestToStoreConsumable not implemented")
}

func ConvertStoreConsumableListToApiConsumableResponseList(storeConsumables []storetp.Consumable) ([]apitp.ConsumableResponse, error) {
	return []apitp.ConsumableResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableListToApiConsumableResponseList not implemented")
}

func ConvertStoreConsumableQuantityListToApiConsumableQuantityResponseList(storeConsumables []storetp.ConsumableQuantity) ([]apitp.ConsumableQuantityResponse, error) {
	return []apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableQuantityListToApiConsumableQuantityResponseList not implemented")
}

func ConvertStoreConsumableQuantityToApiConsumableQuantityResponse(storeConsumable storetp.ConsumableQuantity) (apitp.ConsumableQuantityResponse, error) {
	return apitp.ConsumableQuantityResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableQuantityToApiConsumableQuantityResponse not implemented")
}

func ConvertStoreConsumableToApiConsumableResponse(storeConsumable storetp.Consumable) (apitp.ConsumableResponse, error) {
	return apitp.ConsumableResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreConsumableToApiConsumableResponse not implemented")
}
