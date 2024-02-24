package service

import (
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/constant"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"

	"github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure"
)

var getObjectBySpotOutputFactory output_factory_application.GetObjectBySpotOutputFactory

type GetObjectBySpotService struct {
	objectRepo object_model_domain.ObjectRepository
	spotRepo   spot_model_domain.SpotRepository
}

func (getObjectBySpotService *GetObjectBySpotService) Run(
	getObjectBySpotInputDTO input_factory_application.GetObjectBySpotInputDTO,
) (*output_factory_application.GetObjectBySpotOutputDTO, error) {
	// サーバ環境を初期化
	functionalServerEnv := infrastructure.Init()

	// エリア探索を用いて周辺スポットを取得
	areaSpotCollection, err := getObjectBySpotService.
		spotRepo.
		GetSpotCollectionByCoordinateAndRadiusRepository(
			functionalServerEnv,
			constant.AREA_THRESHOLD,
			getObjectBySpotInputDTO.Coordinate,
		)
	if err != nil {
		return nil, err
	}

	// 周辺スポットのIDを取得
	spotIdCollection := areaSpotCollection.ExtractIdCollection()

	// 周辺スポットを元にスポットに紐づくオブジェクトを取得
	areaObjectCollection, err := getObjectBySpotService.objectRepo.
		GetObjectCollectionBySpotIdsRepository(
			functionalServerEnv,
			spotIdCollection,
			getObjectBySpotInputDTO.User,
		)
	if err != nil {
		return nil, err
	}

	// 周辺スポットをヒントにピンポイントのスポットを取得
	spot, err := getObjectBySpotService.spotRepo.
		GetSpotBySpotIdsAndRawDataFileRepository(
			functionalServerEnv,
			spotIdCollection,
			getObjectBySpotInputDTO.RawDataFile,
		)
	if err != nil {
		return nil, err
	}

	// 屋内推定をしたユーザのピンポイント一点のスポットIDを取得
	spotId := spot.GetId()

	// ピンポイントのスポットを元にスポットに紐づくオブジェクトを取得
	spotObject, err := getObjectBySpotService.objectRepo.
		GetObjectBySpotIdRepository(
			functionalServerEnv,
			spotId,
			getObjectBySpotInputDTO.User,
		)
	if err != nil {
		return nil, err
	}

	// 戻り値を生成
	getObjectBySpotOutputDTO := getObjectBySpotOutputFactory.
		Create(
			*getObjectBySpotInputDTO.User,
			*spotObject,
			*spot,
			areaObjectCollection,
			areaSpotCollection,
		)

	return &getObjectBySpotOutputDTO, nil
}
