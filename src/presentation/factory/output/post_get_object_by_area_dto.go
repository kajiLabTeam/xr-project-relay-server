package output_factory_presentation

import (
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	common_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/common"
)

type PostGetObjectByAreaResponseDTO struct {
	UserId  string                                          `json:"userId" binding:"required,uuid"`
	Objects []common_factory_presentation.ObjectResponseDTO `json:"objects"`
}

type PostGetObjectByAreaResponseFactory struct{}

func (pgobasrf *PostGetObjectByAreaResponseFactory) FromGetObjectByAreaOutputDTO(
	getObjectByAreaOutputDTO output_factory_application.GetObjectByAreaOutputDTO,
) PostGetObjectByAreaResponseDTO {
	var objectResponseDTOs []common_factory_presentation.ObjectResponseDTO

	for _, object := range getObjectByAreaOutputDTO.Objects {
		objectResponseDTO := common_factory_presentation.ObjectResponseDTO{
			Id:       object.Id,
			PosterId: object.PosterId,
			Spot: common_factory_presentation.SpotResponseDTO{
				Id:           object.Spot.Id,
				Name:         object.Spot.Name,
				Floors:       object.Spot.Floors,
				LocationType: object.Spot.LocationType,
				Latitude:     object.Spot.Latitude,
				Longitude:    object.Spot.Longitude,
			},
			ViewUrl: object.ViewUrl,
		}
		objectResponseDTOs = append(objectResponseDTOs, objectResponseDTO)
	}

	return PostGetObjectByAreaResponseDTO{
		UserId:  getObjectByAreaOutputDTO.User.GetId(),
		Objects: objectResponseDTOs,
	}
}
