package input_factory_application

import (
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
	input_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/input"
)

type GetObjectByAreaInputDTO struct {
	User       *user_model_domain.User
	Coordinate *spot_model_domain.Coordinate
}

type GetObjectByAreaInputFactory struct{}

func (gobaif *GetObjectByAreaInputFactory) Create(
	postGetObjectByAreaRequestDTO input_factory_presentation.PostGetObjectByAreaRequestDTO,
) (*GetObjectByAreaInputDTO, error) {
	user, err := user_model_domain.NewUser(postGetObjectByAreaRequestDTO.UserId)
	if err != nil {
		return nil, err
	}

	coordinate, err := spot_model_domain.NewCoordinate(
		postGetObjectByAreaRequestDTO.Latitude,
		postGetObjectByAreaRequestDTO.Longitude,
	)
	if err != nil {
		return nil, err
	}

	return &GetObjectByAreaInputDTO{
		User:       user,
		Coordinate: coordinate,
	}, nil
}
