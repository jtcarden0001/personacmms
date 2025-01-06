package convert

import (
	apitp "github.com/jtcarden0001/personacmms/restapi/internal/types/api"
	storetp "github.com/jtcarden0001/personacmms/restapi/internal/types/store"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func ConvertApiGroupRequestToStoreGroup(groupRequest apitp.GroupRequest) (storetp.Group, error) {
	return storetp.Group{}, ae.New(ae.CodeNotImplemented, "ConvertApiGroupRequestToStoreGroup not implemented")
}

func ConvertStoreGroupListToApiGroupResponseList(storeGroups []storetp.Group) ([]apitp.GroupResponse, error) {
	return []apitp.GroupResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreGroupListToApiGroupResponseList not implemented")
}

func ConvertStoreGroupToApiGroupResponse(storeGroup storetp.Group) (apitp.GroupResponse, error) {
	return apitp.GroupResponse{}, ae.New(ae.CodeNotImplemented, "ConvertStoreGroupToApiGroupResponse not implemented")
}
