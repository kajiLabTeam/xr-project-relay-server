package gateway

import (
	"strconv"

	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway/common"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

type SpotGateway struct{}

func (sg *SpotGateway) GetSpotBySpotIdsAndRawDataFileGateway(spotEstimationServerUrl string, getSpotRequest *spot_record.GetSpotBySpotIdsAndRawDataFileRequest) (*spot_record.GetSpotBySpotIdsAndRawDataFileResponse, error) {
	apiEndpoint := spotEstimationServerUrl + "/api/spot/get"

	responseBody, err := common_gateway.PostFetcher[*spot_record.GetSpotBySpotIdsAndRawDataFileRequest, spot_record.GetSpotBySpotIdsAndRawDataFileResponse](getSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (sg *SpotGateway) GetSpotCollectionByCoordinateAndRadiusGateway(areaEstimationServerUrl string, radius int, getAreaSpotRequest *spot_record.GetSpotCollectionByCoordinateAndRadiusRequest) (*spot_record.GetSpotCollectionByCoordinateAndRadiusResponse, error) {
	apiEndpoint := areaEstimationServerUrl + "/api/area/search?range=" + strconv.Itoa(radius)

	responseBody, err := common_gateway.PostFetcher[*spot_record.GetSpotCollectionByCoordinateAndRadiusRequest, spot_record.GetSpotCollectionByCoordinateAndRadiusResponse](getAreaSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (sg *SpotGateway) CreateSpotGateway(spotEstimationServerUrl string, createSpotRequest *spot_record.CreateSpotRequest) (*spot_record.CreateSpotResponse, error) {
	apiEndpoint := spotEstimationServerUrl + "/api/spot/create"

	responseBody, err := common_gateway.PostFetcher[*spot_record.CreateSpotRequest, spot_record.CreateSpotResponse](createSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
