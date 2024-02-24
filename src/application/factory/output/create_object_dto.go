package output_factory_application

import (
	common_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/common"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
)

type CreateObjectOutputDTO struct {
	ObjectId  string
	PosterId  string
	Extension string
	Spot      common_factory_application.SpotOutputDTO
	UploadUrl string
}

type CreateObjectOutputFactory struct{}

func (coof *CreateObjectOutputFactory) Create(
	extension string,
	o *object_model_domain.Object,
	s *spot_model_domain.Spot,
) *CreateObjectOutputDTO {
	spotDTO := common_factory_application.SpotOutputDTO{
		Id:           s.GetId(),
		Name:         s.GetName(),
		Floors:       s.GetFloors(),
		LocationType: s.GetLocationType(),
		Latitude:     s.GetCoordinate().GetLatitude(),
		Longitude:    s.GetCoordinate().GetLongitude(),
	}

	return &CreateObjectOutputDTO{
		ObjectId:  o.GetId(),
		PosterId:  o.GetPosterId(),
		Extension: extension,
		Spot:      spotDTO,
		UploadUrl: o.GetViewUrl(),
	}
}
