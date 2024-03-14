package repository

import (
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	spot_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/spot"
	"github.com/kajiLabTeam/xr-project-relay-server/domain/repository_impl"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway"
	spot_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/spot"
)

var sg = gateway.SpotGateway{}

type SpotRepository struct{}

func NewSpotRepository() repository_impl.SpotRepositoryImpl {
	return &SpotRepository{}
}

func (sr *SpotRepository) FindForIdsAndRawDataFile(
	spotIds []string,
	rawDataFile []byte,
	application *application_models_domain.Application,
) (*spot_models_domain.Spot, error) {
	findForIdsAndRawDataFileResponse, err := sg.FindForIdsAndRawDataFile(
		spotIds,
		rawDataFile,
		application,
	)
	if err != nil {
		return nil, err
	}
	if findForIdsAndRawDataFileResponse == nil {
		return nil, nil
	}

	resSpot, err := spot_models_domain.NewSpot(
		&findForIdsAndRawDataFileResponse.Id,
		findForIdsAndRawDataFileResponse.Name,
		&findForIdsAndRawDataFileResponse.LocationType,
		findForIdsAndRawDataFileResponse.Floor,
		findForIdsAndRawDataFileResponse.Latitude,
		findForIdsAndRawDataFileResponse.Longitude,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}

func (sr *SpotRepository) FindForCoordinateAndRadius(
	radius int,
	latitude float64,
	longitude float64,
	application *application_models_domain.Application,
) (*spot_models_domain.SpotCollection, error) {
	findForObjectBySpotIAndRawDataFiledResponseFactory := spot_record.FindForCoordinateAndRadiusResponseFactory{}
	findForCoordinateAndRadiusRequest := spot_record.FindForCoordinateAndRadiusRequest{
		Latitude:  latitude,
		Longitude: longitude,
	}

	findForCoordinateAndRadiusResponse, err := sg.FindForCoordinateAndRadius(
		radius,
		&findForCoordinateAndRadiusRequest,
		application,
	)
	if err != nil {
		return nil, err
	}
	if findForCoordinateAndRadiusResponse == nil {
		return nil, nil
	}

	resSpotCollection, err := findForObjectBySpotIAndRawDataFiledResponseFactory.ToDomainSpotCollection(
		findForCoordinateAndRadiusResponse,
	)
	if err != nil {
		return nil, err
	}

	return resSpotCollection, nil
}

func (sr *SpotRepository) Save(
	spot *spot_models_domain.Spot,
	application *application_models_domain.Application,
) (*spot_models_domain.Spot, error) {
	createSpotRequest := spot_record.SaveRequest{
		Name:         spot.GetName(),
		LocationType: spot.GetLocationType(),
		Floor:        spot.GetFloor(),
		Latitude:     spot.GetCoordinate().GetLatitude(),
		Longitude:    spot.GetCoordinate().GetLongitude(),
	}

	createSpotResponse, err := sg.Save(
		spot.GetRawDataFile(),
		&createSpotRequest,
		application,
	)
	if err != nil {
		return nil, err
	}

	resSpot, err := spot_models_domain.NewSpot(
		&createSpotResponse.Id,
		createSpotResponse.Name,
		&createSpotResponse.LocationType,
		createSpotResponse.Floor,
		createSpotResponse.Latitude,
		createSpotResponse.Longitude,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return resSpot, nil
}
