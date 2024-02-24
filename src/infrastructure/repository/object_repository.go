package repository

import (
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/env"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
	"github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/gateway"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/record/object"
)

var og = gateway.ObjectGateway{}

func GetObjectBySpotIdRepository(
	functionServerEnv *env.FunctionServerEnv,
	spotId string,
	u *user_model_domain.User,
) (*object_model_domain.Object, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	getObjectRequest := object_record.
		GetObjectBySpotIdRequest{
		UserId: u.GetId(),
		SpotId: spotId,
	}

	getObjectBySpotIdResponse, err := og.
		GetObjectBySpotIdGateway(
			objectServerUrl,
			&getObjectRequest,
		)
	if err != nil {
		return nil, err
	}

	resObject, err := getObjectBySpotIdResponse.ToDomainObject()

	if err != nil {
		return nil, err
	}

	return resObject, nil

}

func GetObjectCollectionBySpotIdsRepository(
	functionServerEnv *env.FunctionServerEnv,
	spotId []string,
	u *user_model_domain.User,
) (object_model_domain.ObjectCollection, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	getObjectRequest := object_record.
		GetObjectCollectionBySpotIdsRequest{
		UserId: u.GetId(),
		SpotId: spotId,
	}

	getObjectResponse, err := og.
		GetObjectCollectionBySpotIdsGateway(
			objectServerUrl,
			&getObjectRequest,
		)
	if err != nil {
		return nil, err
	}

	resObjectCollection, err := getObjectResponse.ToDomainObjectCollection()

	if err != nil {
		return nil, err
	}

	return resObjectCollection, nil

}

func CreateObjectRepository(
	functionServerEnv *env.FunctionServerEnv,
	userId string,
	extension string,
	spotId string,
) (*object_model_domain.Object, error) {
	objectServerUrl := functionServerEnv.GetObjectServiceUrl()
	createObjectRequest := object_record.
		CreateObjectRequest{
		UserId:    userId,
		SpotId:    spotId,
		Extension: extension,
	}

	createObjectResponse, err := og.
		CreateObjectGateway(
			objectServerUrl,
			&createObjectRequest,
		)
	if err != nil {
		return nil, err
	}

	resObject, err := createObjectResponse.ToDomainObject()
	if err != nil {
		return nil, err
	}

	return resObject, nil
}
