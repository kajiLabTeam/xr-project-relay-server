package repository

import (
	"github.com/kajiLabTeam/xr-project-relay-server/config/env"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/spot"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/user"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/object"
)

var og = gateway.ObjectGateway{}

func GetObjectsBySpotIds(functionServerEnv *env.FunctionServerEnv, u *user_model_domain.User, spotId []string) (*object_model_domain.ObjectCollection, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	getObjectRequest := object_record.GetObjectsBySpotIdsRequest{
		UserId: u.GetId(),
		SpotId: spotId,
	}

	getObjectResponse, err := og.GetObjectsBySpotIdsGateway(objectServerUrl, &getObjectRequest)
	if err != nil {
		return nil, err
	}

	resObjectCollection, err := getObjectResponse.ToDomainObjectCollection()

	if err != nil {
		return nil, err
	}

	return &resObjectCollection, nil

}

func CreateObject(functionServerEnv *env.FunctionServerEnv, o *object_model_domain.Object, s *spot_model_domain.Spot, u *user_model_domain.User) (*object_model_domain.Object, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	createObjectRequest := object_record.CreateObjectRequest{
		Id: u.GetId(),
		Object: object_record.ObjectRequest{
			Id:     o.GetId(),
			SpotId: s.GetId(),
		},
	}

	createObjectResponse, err := og.CreateObjectGateway(objectServerUrl, &createObjectRequest)
	if err != nil {
		return nil, err
	}

	resObject, err := createObjectResponse.ToDomainObject()
	if err != nil {
		return nil, err
	}

	return resObject, nil
}
