package repository

import (
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	object_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/object"
	object_collection_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/object_collection"
	user_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/user"
	"github.com/kajiLabTeam/xr-project-relay-server/domain/repository_impl"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/gateway"
	object_record "github.com/kajiLabTeam/xr-project-relay-server/infrastructure/record/object"
)

var og = gateway.ObjectGateway{}

type ObjectRepository struct{}

func NewObjectRepository() repository_impl.ObjectRepositoryImpl {
	return &ObjectRepository{}
}

func (or *ObjectRepository) FindForSpotId(
	spotId string,
	user *user_models_domain.User,
	application *application_models_domain.Application,
) (*object_models_domain.Object, error) {
	FindForSpotIdRequest := object_record.FindForSpotIdRequest{
		UserId: user.GetId(),
		SpotId: spotId,
	}

	// ゲートウェイを介してスポットを取得
	getObjectBySpotIdResponse, err := og.FindForSpotId(
		&FindForSpotIdRequest,
		application,
	)
	if err != nil {
		return nil, err
	}

	resObject, err := object_models_domain.NewObject(
		&getObjectBySpotIdResponse.Id,
		&getObjectBySpotIdResponse.PosterId,
		getObjectBySpotIdResponse.Extension,
		&getObjectBySpotIdResponse.SpotId,
		nil,
		&getObjectBySpotIdResponse.ViewUrl,
	)
	if err != nil {
		return nil, err
	}

	return resObject, nil

}

func (or *ObjectRepository) FindForSpotIds(
	spotIds []string,
	user *user_models_domain.User,
	application *application_models_domain.Application,
) (*object_collection_models_domain.ObjectCollection, error) {
	findForObjectBySpotIAndRawDataFiledResponseFactory := object_record.FindForObjectBySpotIAndRawDataFiledResponseFactory{}
	FindForSpotIdAndRawData := object_record.FindForSpotIdsRequest{
		UserId:  user.GetId(),
		SpotIds: spotIds,
	}

	// ゲートウェイを介してスポットを取得
	findForSpotIdAndRawDataResponse, err := og.FindForSpotIds(
		&FindForSpotIdAndRawData,
		application,
	)
	if err != nil {
		return nil, err
	}
	if findForSpotIdAndRawDataResponse == nil {
		return nil, nil
	}

	resObjectCollection, err := findForObjectBySpotIAndRawDataFiledResponseFactory.ToDomainObject(
		findForSpotIdAndRawDataResponse,
	)
	if err != nil {
		return nil, err
	}

	return resObjectCollection, nil

}

func (or *ObjectRepository) Save(
	spotId string,
	user *user_models_domain.User,
	object *object_models_domain.Object,
	application *application_models_domain.Application,
) (*object_models_domain.Object, error) {
	createObjectRequest := object_record.SaveRequest{
		UserId:    user.GetId(),
		SpotId:    spotId,
		Extension: object.GetExtension(),
	}

	// ゲートウェイを介してスポットを保存
	saveObjectResponse, err := og.Save(&createObjectRequest, application)
	if err != nil {
		return nil, err
	}

	resObject, err := object_models_domain.NewObject(
		&saveObjectResponse.Id,
		&saveObjectResponse.PosterId,
		saveObjectResponse.Extension,
		&saveObjectResponse.SpotId,
		nil,
		&saveObjectResponse.UploadUrl,
	)
	if err != nil {
		return nil, err
	}

	return resObject, nil
}
