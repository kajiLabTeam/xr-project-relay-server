package repository_impl

import (
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	spot_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/spot"
)

type SpotRepositoryImpl interface {
	FindForIdsAndRawDataFile(
		spotIds []string,
		rawDataFile []byte,
		application *application_models_domain.Application,
	) (*spot_models_domain.Spot, error)

	FindForCoordinateAndRadius(
		radius int,
		latitude float64,
		longitude float64,
		application *application_models_domain.Application,
	) (*spot_models_domain.SpotCollection, error)

	Save(
		spot *spot_models_domain.Spot,
		application *application_models_domain.Application,
	) (*spot_models_domain.Spot, error)
}
