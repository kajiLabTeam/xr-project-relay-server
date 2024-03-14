package gateway

import (
	"encoding/json"
	"os"

	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	common_gateway "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway/common"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/object"
)

type ObjectGateway struct{}

func (og ObjectGateway) FindForSpotId(
	findForSpotIdRequest *object_record.FindForSpotIdRequest,
	application *application_models_domain.Application,
) (*object_record.FindForSpotIdResponse, error) {
	endpoint := os.Getenv("OBJECT_SERVER_URL") + "/api/object/get"

	request := common_gateway.NewRequest(application)

	responseBody, err := request.MakeApplicationJsonRequest(
		endpoint,
		findForSpotIdRequest,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var findForSpotIdResponse object_record.FindForSpotIdResponse
	err = json.Unmarshal(responseBody, &findForSpotIdResponse)
	if err != nil {
		return nil, err
	}

	return &findForSpotIdResponse, nil
}

func (og ObjectGateway) FindForSpotIds(
	FindForSpotIdAndRawDataRequest *object_record.FindForSpotIdsRequest,
	application *application_models_domain.Application,
) (*object_record.FindForObjectBySpotIAndRawDataFiledResponse, error) {
	endpoint := os.Getenv("OBJECT_SERVER_URL") + "/api/object/collection/get"

	request := common_gateway.NewRequest(application)

	responseBody, err := request.MakeApplicationJsonRequest(
		endpoint,
		FindForSpotIdAndRawDataRequest,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var findForSpotIdAndRawDataResponse object_record.FindForObjectBySpotIAndRawDataFiledResponse
	err = json.Unmarshal(responseBody, &findForSpotIdAndRawDataResponse)
	if err != nil {
		return nil, err
	}

	return &findForSpotIdAndRawDataResponse, nil
}

func (og ObjectGateway) Save(
	saveRequest *object_record.SaveRequest,
	application *application_models_domain.Application,
) (*object_record.SaveResponse, error) {
	endpoint := os.Getenv("OBJECT_SERVER_URL") + "/api/object/create"

	request := common_gateway.NewRequest(application)

	responseBody, err := request.MakeApplicationJsonRequest(
		endpoint,
		saveRequest,
	)
	if err != nil {
		return nil, err
	}
	// 404エラーの場合
	if responseBody == nil {
		return nil, nil
	}

	var createObjectResponse object_record.SaveResponse
	err = json.Unmarshal(responseBody, &createObjectResponse)
	if err != nil {
		return nil, err
	}

	return &createObjectResponse, nil
}
