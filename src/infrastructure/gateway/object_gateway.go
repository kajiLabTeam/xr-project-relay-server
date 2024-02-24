package gateway

import (
	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway/common"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/object"
)

type ObjectGateway struct{}

func (og ObjectGateway) GetObjectsBySpotIdsGateway(objectServerUrl string, getObjectRequest *object_record.GetObjectsBySpotIdsRequest) (*object_record.GetObjectsBySpotIdsResponse, error) {
	apiEndpoint := objectServerUrl + "/api/object/get"

	getObjectResponse, err := common_gateway.PostFetcher[*object_record.GetObjectsBySpotIdsRequest, object_record.GetObjectsBySpotIdsResponse](getObjectRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return getObjectResponse, nil
}

func (og ObjectGateway) CreateObjectGateway(objectServerUrl string, createObjectRequest *object_record.CreateObjectRequest) (*object_record.CreateObjectResponse, error) {
	apiEndpoint := objectServerUrl + "/api/object/upload"

	createObjectResponse, err := common_gateway.PostFetcher[*object_record.CreateObjectRequest, object_record.CreateObjectResponse](createObjectRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return createObjectResponse, nil
}
