package gateway

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/kajiLabTeam/xr-project-relay-server/config"
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway/common"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

type SpotGateway struct{}

func (sg *SpotGateway) FindForIdsAndRawDataFile(
	spotId []string,
	rawDataFile []byte,
	application *application_models_domain.Application,
) (*spot_record.FindForIdsAndRawDataFileResponse, error) {
	var query string
	for _, id := range spotId {
		query += "spotIds=" + id + "&"
	}

	endpoint := os.Getenv("SPOT_ESTIMATION_SERVER_URL") + "/api/spot/search" + "?" + query

	request := common_gateway.NewRequest(application)
	if err := request.AddFindSpotForIdsAndRawDataFileRequest(
		config.RAW_DATA_FILE_FIELD,
		config.RAW_DATA_FILE_NAME,
		rawDataFile,
	); err != nil {
		return nil, err
	}

	responseBody, err := request.MakeMultipartRequest(
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var findForIdsAndRawDataFileResponse spot_record.FindForIdsAndRawDataFileResponse
	err = json.Unmarshal(responseBody, &findForIdsAndRawDataFileResponse)
	if err != nil {
		return nil, err
	}

	return &findForIdsAndRawDataFileResponse, nil
}

func (sg *SpotGateway) FindForCoordinateAndRadius(
	radius int,
	findForCoordinateAndRadiusRequest *spot_record.FindForCoordinateAndRadiusRequest,
	application *application_models_domain.Application,
) (*spot_record.FindForCoordinateAndRadiusResponse, error) {
	endpoint := os.Getenv("AREA_ESTIMATION_SERVER_URL") + "/api/area/search?radius=" + strconv.Itoa(radius)

	request := common_gateway.NewRequest(application)
	responseBody, err := request.MakeApplicationJsonRequest(
		endpoint,
		findForCoordinateAndRadiusRequest,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var findForCoordinateAndRadiusResponse spot_record.FindForCoordinateAndRadiusResponse
	err = json.Unmarshal(responseBody, &findForCoordinateAndRadiusResponse)
	if err != nil {
		return nil, err
	}

	return &findForCoordinateAndRadiusResponse, nil
}

func (sg *SpotGateway) Save(
	rawDataFile []byte,
	saveRequest *spot_record.SaveRequest,
	application *application_models_domain.Application,
) (*spot_record.SaveResponse, error) {
	endpoint := os.Getenv("SPOT_ESTIMATION_SERVER_URL") + "/api/spot/create"

	request := common_gateway.NewRequest(application)
	if err := request.AddSaveSpotRequest(
		config.RAW_DATA_FILE_FIELD,
		config.RAW_DATA_FILE_NAME,
		rawDataFile,
		saveRequest,
	); err != nil {
		return nil, err
	}

	responseBody, err := request.MakeMultipartRequest(
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var saveResponse spot_record.SaveResponse
	err = json.Unmarshal(responseBody, &saveResponse)
	if err != nil {
		return nil, err
	}

	return &saveResponse, nil
}
