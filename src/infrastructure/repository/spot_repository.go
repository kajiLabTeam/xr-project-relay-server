package repository

import (
	"os"

	"github.com/kajiLabTeam/xr-project-relay-server/config/env"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/spot"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

var sg = gateway.SpotGateway{}

func GetSpot(functionServerEnv *env.FunctionServerEnv, spotIds *[]string, rawData *os.File) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getSpotRequest := spot_record.GetSpotRequest{
		Ids:         *spotIds,
		RawDataFile: rawData,
	}

	getSpotResponse, err := sg.GetSpot(spotEstimationServerUrl, &getSpotRequest)
	if err != nil {
		return nil, err
	}

	coordinate, err := spot_model_domain.NewCoordinate(
		getSpotResponse.Coordinate.Latitude,
		getSpotResponse.Coordinate.Longitude,
	)
	if err != nil {
		return nil, err
	}

	resSpot, err := spot_model_domain.NewSpot(
		getSpotResponse.Id,
		getSpotResponse.Name,
		getSpotResponse.LocationType,
		getSpotResponse.Floors,
		coordinate,
	)
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}

func GetAreaSpots(functionServerEnv *env.FunctionServerEnv, radius int, coordinate *spot_model_domain.Coordinate) (spot_model_domain.SpotCollection, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	getAreaSpotRequest := spot_record.GetAreaSpotRequest{
		Coordinate: spot_record.CoordinateRequest{
			Latitude:  coordinate.GetLatitude(),
			Longitude: coordinate.GetLongitude(),
		},
	}

	getAreaSpotResponse, err := sg.GetAreaSpots(spotEstimationServerUrl, radius, &getAreaSpotRequest)
	if err != nil {
		return nil, err
	}

	resSpotCollection, err := getAreaSpotResponse.ToDomainSpotCollection()

	return resSpotCollection, nil
}

func CreateSpot(functionServerEnv *env.FunctionServerEnv, rawDataFile *os.File, s *spot_model_domain.Spot) (*spot_model_domain.Spot, error) {
	spotEstimationServerUrl := functionServerEnv.GetSpotEstimationServiceUrl()
	createSpotRequest := spot_record.CreateSpotRequest{
		Name:         s.GetName(),
		Floors:       s.GetFloors(),
		LocationType: s.GetLocationType(),
		Coordinate: spot_record.CoordinateRequest{
			Latitude:  s.GetCoordinate().GetLatitude(),
			Longitude: s.GetCoordinate().GetLongitude(),
		},
		RawDataFile: rawDataFile,
	}

	createSpotResponse, err := sg.CreateSpot(spotEstimationServerUrl, &createSpotRequest)
	if err != nil {
		return nil, err
	}

	coordinate, err := spot_model_domain.NewCoordinate(
		createSpotResponse.Coordinate.Latitude,
		createSpotResponse.Coordinate.Longitude,
	)

	resSpot, err := spot_model_domain.NewSpot(
		createSpotResponse.Id,
		createSpotResponse.Name,
		createSpotResponse.LocationType,
		createSpotResponse.Floors,
		coordinate,
	)

	return resSpot, nil
}
