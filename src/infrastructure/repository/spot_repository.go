package repository

import (
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/env"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	input_dto_infrastructure "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/dto/input"
	"github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/gateway"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/record/spot"
)

var sg = gateway.SpotGateway{}

func GetSpotBySpotIdsAndRawDataFileRepository(
	functionServerEnv *env.FunctionServerEnv,
	spotIds []string,
	rawData []byte,
) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getSpotRequest := spot_record.
		GetSpotBySpotIdsAndRawDataFileRequest{
		SpotIds:     spotIds,
		RawDataFile: rawData,
	}

	getSpotResponse, err := sg.
		GetSpotBySpotIdsAndRawDataFileGateway(
			spotEstimationServerUrl,
			&getSpotRequest,
		)
	if err != nil {
		return nil, err
	}

	resSpot, err := getSpotResponse.ToDomainSpot()
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}

func GetSpotCollectionByCoordinateAndRadiusRepository(
	functionServerEnv *env.FunctionServerEnv,
	radius int,
	latitude float64,
	longitude float64,
) (spot_model_domain.SpotCollection, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getAreaSpotRequest := spot_record.
		GetSpotCollectionByCoordinateAndRadiusRequest{
		Latitude:  latitude,
		Longitude: longitude,
	}

	getAreaSpotResponse, err := sg.
		GetSpotCollectionByCoordinateAndRadiusGateway(
			spotEstimationServerUrl,
			radius,
			&getAreaSpotRequest,
		)
	if err != nil {
		return nil, err
	}

	resSpotCollection, err := getAreaSpotResponse.ToDomainSpotCollection()
	if err != nil {
		return nil, err
	}

	return resSpotCollection, nil
}

func CreateSpotRepository(
	csidto *input_dto_infrastructure.CreateSpotRepositoryInputDTO,
) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := csidto.FunctionServerEnv.GetSpotEstimationServiceUrl()
	createSpotRequest := spot_record.CreateSpotRequest{
		Name:         csidto.Name,
		Floors:       csidto.Floors,
		LocationType: csidto.LocationType,
		Latitude:     csidto.Latitude,
		Longitude:    csidto.Longitude,
		RawDataFile:  csidto.RawDataFile,
	}

	createSpotResponse, err := sg.
		CreateSpotGateway(
			spotEstimationServerUrl,
			&createSpotRequest,
		)
	if err != nil {
		return nil, err
	}

	resSpot, err := createSpotResponse.ToDomainSpot()
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}
