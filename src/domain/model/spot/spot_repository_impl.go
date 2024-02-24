package spot_model_domain

import (
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/env"
	input_dto_infrastructure "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/dto/input"
)

type SpotRepository interface {
	GetSpotBySpotIdsAndRawDataFileRepository(
		functionServerEnv *env.FunctionServerEnv,
		spotIds []string,
		rawData []byte,
	) (*Spot, error)
	GetSpotCollectionByCoordinateAndRadiusRepository(
		functionServerEnv *env.FunctionServerEnv,
		radius int,
		c *Coordinate,
	) (SpotCollection, error)
	CreateSpotRepository(
		csidto *input_dto_infrastructure.CreateSpotRepositoryInputDTO,
	) (*Spot, error)
}
