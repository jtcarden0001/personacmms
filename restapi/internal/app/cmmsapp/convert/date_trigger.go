package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func ConvertApiDateTriggerRequestToStoreDateTrigger(dateTriggerRequest apitp.DateTriggerRequest) (storetp.DateTrigger, error) {
	return storetp.DateTrigger{}, ae.New(ae.CodeNotImplemented, "ConvertApiDateTriggerRequestToStoreDateTrigger not implemented")
}

func ConvertStoreDateTriggerListToApiDateTriggerResponseList(storeDateTriggers []storetp.DateTrigger) ([]apitp.DateTriggerResponse, error) {
	return []apitp.DateTriggerResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreDateTriggerListToApiDateTriggerResponseList not implemented")
}

func ConvertStoreDateTriggerToApiDateTriggerResponse(storeDateTrigger storetp.DateTrigger) (apitp.DateTriggerResponse, error) {
	return apitp.DateTriggerResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreDateTriggerToApiDateTriggerResponse not implemented")
}
