package repository

import (
	"os"

	"github.com/kajiLabTeam/xr-project-relay-server/config/env"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/spot"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

var sg = gateway.SpotGateway{}

func GetSpotBySpotIdsAndRawDataFile(functionServerEnv *env.FunctionServerEnv, spotIds []string, rawData *os.File) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getSpotRequest := spot_record.GetSpotBySpotIdsAndRawDataFileRequest{
		SpotIds:     spotIds,
		RawDataFile: rawData,
	}

	getSpotResponse, err := sg.GetSpotBySpotIdsAndRawDataFileGateway(spotEstimationServerUrl, &getSpotRequest)
	if err != nil {
		return nil, err
	}

	resSpot, err := getSpotResponse.ToDomainSpot()
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}

func GetSpotCollectionByCoordinateAndRadius(functionServerEnv *env.FunctionServerEnv, radius int, coordinate *spot_model_domain.Coordinate) (spot_model_domain.SpotCollection, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getAreaSpotRequest := spot_record.GetSpotCollectionByCoordinateAndRadiusRequest{
		Latitude:  coordinate.GetLatitude(),
		Longitude: coordinate.GetLongitude(),
	}

	getAreaSpotResponse, err := sg.GetSpotCollectionByCoordinateAndRadiusGateway(spotEstimationServerUrl, radius, &getAreaSpotRequest)
	if err != nil {
		return nil, err
	}

	resSpotCollection, err := getAreaSpotResponse.ToDomainSpotCollection()
	if err != nil {
		return nil, err
	}

	return resSpotCollection, nil
}

func CreateSpot(functionServerEnv *env.FunctionServerEnv, rawDataFile *os.File, s *spot_model_domain.Spot) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	createSpotRequest := spot_record.CreateSpotRequest{
		Name:         s.GetName(),
		Floors:       s.GetFloors(),
		LocationType: s.GetLocationType(),
		Latitude:     s.GetCoordinate().GetLatitude(),
		Longitude:    s.GetCoordinate().GetLongitude(),
		RawDataFile:  rawDataFile,
	}

	createSpotResponse, err := sg.CreateSpotGateway(spotEstimationServerUrl, &createSpotRequest)
	if err != nil {
		return nil, err
	}

	resSpot, err := createSpotResponse.ToDomainSpot()
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}
