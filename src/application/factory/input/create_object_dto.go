package input_factory_application

import (
	common_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/common"
	input_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/input"
)

type CreateObjectInputDTO struct {
	UserId    string
	Extension string
	Spot      common_factory_application.SpotInputDTO
}

type CreateObjectInputFactory struct{}

func (coif *CreateObjectInputFactory) Create(
	postObjectCreateRequestDTO *input_factory_presentation.PostObjectCreateRequestDTO,
) *CreateObjectInputDTO {
	spotInputDTO := common_factory_application.SpotInputDTO{
		Name:         postObjectCreateRequestDTO.Spot.Name,
		Floors:       postObjectCreateRequestDTO.Spot.Floors,
		LocationType: postObjectCreateRequestDTO.Spot.LocationType,
		Latitude:     postObjectCreateRequestDTO.Spot.Latitude,
		Longitude:    postObjectCreateRequestDTO.Spot.Longitude,
	}

	return &CreateObjectInputDTO{
		UserId:    postObjectCreateRequestDTO.UserId,
		Extension: postObjectCreateRequestDTO.Extension,
		Spot:      spotInputDTO,
	}
}
