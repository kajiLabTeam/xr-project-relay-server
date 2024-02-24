package output_factory_presentation

import (
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	common_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/common"
)

type PostGetObjectBySpotResponseDTO struct {
	Id         string                                          `json:"id" binding:"required,uuid"`
	SpotObject common_factory_presentation.ObjectResponseDTO   `json:"spotObject"`
	AreaObject []common_factory_presentation.ObjectResponseDTO `json:"areaObject"`
}

type PostGetObjectBySpotResponseFactory struct{}

func (pgobsrf *PostGetObjectBySpotResponseFactory) FromGetObjectBySpotOutputDTO(
	gobsodto output_factory_application.GetObjectBySpotOutputDTO,
) PostGetObjectBySpotResponseDTO {
	spotObject := common_factory_presentation.ObjectResponseDTO{
		Id:       gobsodto.SpotObject.Id,
		PosterId: gobsodto.SpotObject.PosterId,
		Spot: common_factory_presentation.SpotResponseDTO{
			Id:           gobsodto.SpotObject.Spot.Id,
			Name:         gobsodto.SpotObject.Spot.Name,
			Floors:       gobsodto.SpotObject.Spot.Floors,
			LocationType: gobsodto.SpotObject.Spot.LocationType,
			Latitude:     gobsodto.SpotObject.Spot.Latitude,
			Longitude:    gobsodto.SpotObject.Spot.Longitude,
		},
		ViewUrl: gobsodto.SpotObject.ViewUrl,
	}

	var areaObjectCollection []common_factory_presentation.ObjectResponseDTO

	for _, areaObject := range gobsodto.AreaObjects {
		areaObjectResponseDTO := common_factory_presentation.ObjectResponseDTO{
			Id:       areaObject.Id,
			PosterId: areaObject.PosterId,
			Spot: common_factory_presentation.SpotResponseDTO{
				Id:           areaObject.Spot.Id,
				Name:         areaObject.Spot.Name,
				Floors:       areaObject.Spot.Floors,
				LocationType: areaObject.Spot.LocationType,
				Latitude:     areaObject.Spot.Latitude,
				Longitude:    areaObject.Spot.Longitude,
			},
			ViewUrl: areaObject.ViewUrl,
		}
		areaObjectCollection = append(areaObjectCollection, areaObjectResponseDTO)
	}

	return PostGetObjectBySpotResponseDTO{
		Id:         gobsodto.User.GetId(),
		SpotObject: spotObject,
		AreaObject: areaObjectCollection,
	}
}
