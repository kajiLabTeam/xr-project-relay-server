package output_factory_application

import (
	common_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/common"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
)

type GetObjectBySpotOutputDTO struct {
	User        *user_model_domain.User
	SpotObject  *common_factory_application.ObjectDTO
	AreaObjects common_factory_application.ObjectDTOCollection
}

type GetObjectBySpotOutputFactory struct{}

func (gobsof *GetObjectBySpotOutputFactory) Create(
	user user_model_domain.User,
	spotObject object_model_domain.Object,
	spot spot_model_domain.Spot,
	areaObjectCollection object_model_domain.ObjectCollection,
	areaSpotCollection spot_model_domain.SpotCollection,
) GetObjectBySpotOutputDTO {
	spotObjectDTO := common_factory_application.ObjectDTO{
		Id:       spotObject.GetId(),
		PosterId: spotObject.GetPosterId(),
		Spot: common_factory_application.SpotOutputDTO{
			Id:           spot.GetId(),
			Name:         spot.GetName(),
			Floors:       spot.GetFloors(),
			LocationType: spot.GetLocationType(),
			Latitude:     spot.GetCoordinate().GetLatitude(),
			Longitude:    spot.GetCoordinate().GetLongitude(),
		},
		ViewUrl: spotObject.GetViewUrl(),
	}

	var areaObjectCollectionDTO common_factory_application.ObjectDTOCollection

	for i, areaObject := range areaObjectCollection {
		areaObjectDTO := common_factory_application.ObjectDTO{
			Id:       areaObject.GetId(),
			PosterId: areaObject.GetPosterId(),
			Spot: common_factory_application.SpotOutputDTO{
				Id:           areaSpotCollection[i].GetId(),
				Name:         areaSpotCollection[i].GetName(),
				Floors:       areaSpotCollection[i].GetFloors(),
				LocationType: areaSpotCollection[i].GetLocationType(),
				Latitude:     areaSpotCollection[i].GetCoordinate().GetLatitude(),
				Longitude:    areaSpotCollection[i].GetCoordinate().GetLongitude(),
			},
			ViewUrl: areaObject.GetViewUrl(),
		}
		areaObjectCollectionDTO = append(areaObjectCollectionDTO, areaObjectDTO)
	}

	// GetObjectBySpotOutputDTOに詰め替え
	result := GetObjectBySpotOutputDTO{
		User:        &user,
		SpotObject:  &spotObjectDTO,
		AreaObjects: areaObjectCollectionDTO,
	}

	return result
}
