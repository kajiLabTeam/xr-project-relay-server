package output_factory_application

import (
	common_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/common"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
)

type GetObjectByAreaOutputDTO struct {
	User    *user_model_domain.User
	Objects common_factory_application.ObjectDTOCollection
}

type GetObjectByAreaOutputFactory struct{}

func (gobaof *GetObjectByAreaOutputFactory) Create(
	user user_model_domain.User,
	objectCollection object_model_domain.ObjectCollection,
	spotCollection spot_model_domain.SpotCollection,
) GetObjectByAreaOutputDTO {
	var objectCollectionDTO common_factory_application.ObjectDTOCollection

	for i, object := range objectCollection {
		objectDTO := common_factory_application.ObjectDTO{
			Id:       object.GetId(),
			PosterId: object.GetPosterId(),
			Spot: common_factory_application.SpotOutputDTO{
				Id:           spotCollection[i].GetId(),
				Name:         spotCollection[i].GetName(),
				Floors:       spotCollection[i].GetFloors(),
				LocationType: spotCollection[i].GetLocationType(),
				Latitude:     spotCollection[i].GetCoordinate().GetLatitude(),
				Longitude:    spotCollection[i].GetCoordinate().GetLongitude(),
			},
			ViewUrl: object.GetViewUrl(),
		}
		objectCollectionDTO = append(objectCollectionDTO, objectDTO)
	}

	return GetObjectByAreaOutputDTO{
		User:    &user,
		Objects: objectCollectionDTO,
	}
}
