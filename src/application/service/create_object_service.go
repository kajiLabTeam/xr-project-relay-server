package service

import (
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	"github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure"
	input_dto_infrastructure "github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure/dto/input"
)

var createObjectOutputFactory output_factory_application.CreateObjectOutputFactory

type CreateObjectService struct {
	objectRepo object_model_domain.ObjectRepository
	spotRepo   spot_model_domain.SpotRepository
}

func (createObjectService *CreateObjectService) Run(
	createObjectInputDTO input_factory_application.CreateObjectInputDTO,
) (*output_factory_application.CreateObjectOutputDTO, error) {
	// サーバ環境を初期化
	functionServerUrl := infrastructure.Init()

	// スポットをDBに登録
	spot, err := createObjectService.spotRepo.
		CreateSpotRepository(
			&input_dto_infrastructure.CreateSpotRepositoryInputDTO{
				FunctionServerEnv: functionServerUrl,
				RawDataFile:       createObjectInputDTO.Spot.RawDataFile,
				Name:              createObjectInputDTO.Spot.Name,
				Floors:            createObjectInputDTO.Spot.Floors,
				LocationType:      createObjectInputDTO.Spot.LocationType,
				Latitude:          createObjectInputDTO.Spot.Latitude,
				Longitude:         createObjectInputDTO.Spot.Longitude,
			})
	if err != nil {
		return nil, err
	}

	// 登録したスポットのIDを取得
	spotId := spot.GetId()

	// オブジェクトをDBに登録
	object, err := createObjectService.objectRepo.
		CreateObjectRepository(
			functionServerUrl,
			createObjectInputDTO.UserId,
			createObjectInputDTO.Extension,
			spotId,
		)
	if err != nil {
		return nil, err
	}

	// 戻り値を生成
	createObjectOutputDTO := createObjectOutputFactory.
		Create(
			createObjectInputDTO.Extension,
			object,
			spot,
		)

	return createObjectOutputDTO, nil
}
