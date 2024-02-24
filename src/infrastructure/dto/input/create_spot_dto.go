package input_dto_infrastructure

import "github.com/kajiLabTeam/xr-project-relay-server/src/config/env"

type CreateSpotRepositoryInputDTO struct {
	FunctionServerEnv *env.FunctionServerEnv
	RawDataFile       []byte
	Name              string
	Floors            int
	LocationType      string
	Latitude          float64
	Longitude         float64
}
