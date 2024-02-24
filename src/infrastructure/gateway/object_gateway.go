package gateway

import (
	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/gateway/common"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/record/object"
)

type ObjectGateway struct{}

func (og ObjectGateway) GetObjectBySpotIdGateway(objectServerUrl string, getObjectRequest *object_record.GetObjectBySpotIdRequest) (*object_record.GetObjectBySpotIdResponse, error) {
	apiEndpoint := objectServerUrl + "/api/object/get"

	getObjectBySpotIdResponse, err := common_gateway.PostFetcher[*object_record.GetObjectBySpotIdRequest, object_record.GetObjectBySpotIdResponse](getObjectRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return getObjectBySpotIdResponse, nil
}

func (og ObjectGateway) GetObjectCollectionBySpotIdsGateway(objectServerUrl string, getObjectRequest *object_record.GetObjectCollectionBySpotIdsRequest) (*object_record.GetObjectCollectionBySpotIdsResponse, error) {
	apiEndpoint := objectServerUrl + "/api/object/collection/get"

	getObjectCollectionBySpotIdsResponse, err := common_gateway.PostFetcher[*object_record.GetObjectCollectionBySpotIdsRequest, object_record.GetObjectCollectionBySpotIdsResponse](getObjectRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return getObjectCollectionBySpotIdsResponse, nil
}

func (og ObjectGateway) CreateObjectGateway(objectServerUrl string, createObjectRequest *object_record.CreateObjectRequest) (*object_record.CreateObjectResponse, error) {
	apiEndpoint := objectServerUrl + "/api/object/upload"

	createObjectResponse, err := common_gateway.PostFetcher[*object_record.CreateObjectRequest, object_record.CreateObjectResponse](createObjectRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return createObjectResponse, nil
}
