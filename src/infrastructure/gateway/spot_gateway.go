package gateway

import (
	"strconv"

	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway/common"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

type SpotGateway struct{}

func (sg *SpotGateway) GetSpot(spotEstimationServerUrl string, getSpotRequest *spot_record.GetSpotRequest) (*spot_record.GetSpotResponse, error) {
	apiEndpoint := spotEstimationServerUrl + "/api/spot/get"

	responseBody, err := common_gateway.PostFetcher[*spot_record.GetSpotRequest, spot_record.GetSpotResponse](getSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (sg *SpotGateway) GetAreaSpots(areaEstimationServerUrl string, radius int, getAreaSpotRequest *spot_record.GetAreaSpotRequest) (*spot_record.GetAreaSpotResponse, error) {
	apiEndpoint := areaEstimationServerUrl + "/api/area/search?range=" + strconv.Itoa(radius)

	responseBody, err := common_gateway.PostFetcher[*spot_record.GetAreaSpotRequest, spot_record.GetAreaSpotResponse](getAreaSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// TODO: Multiple file upload
func (sg *SpotGateway) CreateSpot(spotEstimationServerUrl string, createSpotRequest *spot_record.CreateSpotRequest) (*spot_record.CreateSpotResponse, error) {
	apiEndpoint := spotEstimationServerUrl + "/api/spot/create"

	responseBody, err := common_gateway.PostFetcher[*spot_record.CreateSpotRequest, spot_record.CreateSpotResponse](createSpotRequest, apiEndpoint)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
