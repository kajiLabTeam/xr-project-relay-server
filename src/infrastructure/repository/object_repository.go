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

func GetObject(functionServerEnv *env.FunctionServerEnv, o *object_model_domain.Object, s *spot_model_domain.Spot, u *user_model_domain.User) (*object_model_domain.Object, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	getObjectRequest := object_record.GetObjectRequest{
		UserId: u.GetId(),
		SpotId: s.GetId(),
	}

	getObjectResponse, err := og.GetObject(objectServerUrl, &getObjectRequest)
	if err != nil {
		return nil, err
	}

	resObject, err := object_model_domain.NewObject(
		getObjectResponse.Object.Id,
		getObjectResponse.Object.PosterId,
		getObjectResponse.Object.ViewUrl,
		getObjectResponse.Object.UploadUrl,
	)
	if err != nil {
		return nil, err
	}

	return resObject, nil

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

	createObjectResponse, err := og.CreateObject(objectServerUrl, &createObjectRequest)
	if err != nil {
		return nil, err
	}

	resObject, err := object_model_domain.NewObject(
		createObjectResponse.Object.Id,
		createObjectResponse.Object.PosterId,
		createObjectResponse.Object.ViewUrl,
		createObjectResponse.Object.UploadUrl,
	)
	if err != nil {
		return nil, err
	}

	return resObject, nil

}
