package spot_model_domain

import (
	"os"

	"github.com/kajiLabTeam/xr-project-relay-server/config/env"
)

type SpotRepository interface {
	GetSpot(functionServerEnv *env.FunctionServerEnv, spotIds *[]string, rawData *os.File) (*Spot, error)
	GetAreaSpots(functionServerEnv *env.FunctionServerEnv, radius int, coordinate *Coordinate) (SpotCollection, error)
	CreateSpot(functionServerEnv *env.FunctionServerEnv, rawDataFile *os.File, s *Spot) (*Spot, error)
}
