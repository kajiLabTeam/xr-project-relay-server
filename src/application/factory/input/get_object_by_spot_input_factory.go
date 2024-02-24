package input_factory_application

import (
	"io"
	"mime/multipart"

	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
)

type GetObjectBySpotInputDTO struct {
	User        *user_model_domain.User
	RawDataFile []byte
	Coordinate  *spot_model_domain.Coordinate
}

type GetObjectBySpotInputFactory struct{}

func (gobsif *GetObjectBySpotInputFactory) Create(
	userId string,
	latitude float64,
	longitude float64,
	rawDataFile multipart.File,
) (*GetObjectBySpotInputDTO, error) {
	user, err := user_model_domain.NewUser(userId)
	if err != nil {
		return nil, err
	}

	coordinate, err := spot_model_domain.NewCoordinate(latitude, longitude)
	if err != nil {
		return nil, err
	}

	file, err := io.ReadAll(rawDataFile)
	if err != nil {
		return nil, err
	}

	return &GetObjectBySpotInputDTO{
		User:        user,
		RawDataFile: file,
		Coordinate:  coordinate,
	}, nil
}
