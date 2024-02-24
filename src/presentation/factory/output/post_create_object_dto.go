package output_factory_presentation

import (
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	common_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/common"
)

type PostCreateObjectResponseDTO struct {
	ObjectId  string                                      `json:"objectId" binding:"required,uuid"`
	PosterId  string                                      `json:"posterId" binding:"required,uuid"`
	Extension string                                      `json:"extension"`
	Spot      common_factory_presentation.SpotResponseDTO `json:"spot"`
	UploadUrl string                                      `json:"viewUrl" binding:"required,url"`
}

type PostCreateObjectResponseFactory struct{}

func (pcoof *PostCreateObjectResponseFactory) FromCreateObjectOutputDTO(
	createObjectOutputDTO *output_factory_application.CreateObjectOutputDTO,
) PostCreateObjectResponseDTO {
	spotResponseDTO := common_factory_presentation.SpotResponseDTO{
		Id:           createObjectOutputDTO.Spot.Id,
		Name:         createObjectOutputDTO.Spot.Name,
		Floors:       createObjectOutputDTO.Spot.Floors,
		LocationType: createObjectOutputDTO.Spot.LocationType,
		Latitude:     createObjectOutputDTO.Spot.Latitude,
		Longitude:    createObjectOutputDTO.Spot.Longitude,
	}

	return PostCreateObjectResponseDTO{
		ObjectId:  createObjectOutputDTO.ObjectId,
		PosterId:  createObjectOutputDTO.PosterId,
		Extension: createObjectOutputDTO.Extension,
		Spot:      spotResponseDTO,
		UploadUrl: createObjectOutputDTO.UploadUrl,
	}
}
