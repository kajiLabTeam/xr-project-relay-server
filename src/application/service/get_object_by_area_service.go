package service

import (
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	output_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/output"
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/constant"
	object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/object"
	spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"
	"github.com/kajiLabTeam/xr-project-relay-server/src/infrastructure"
)

var gobaof output_factory_application.GetObjectByAreaOutputFactory

type GetObjectByAreaService struct {
	objectRepo object_model_domain.ObjectRepository
	spotRepo   spot_model_domain.SpotRepository
}

func (gobas *GetObjectByAreaService) Run(
	gobaidto input_factory_application.GetObjectByAreaInputDTO,
) (
	*output_factory_application.GetObjectByAreaOutputDTO,
	error,
) {
	// サーバ環境を初期化
	functionalServerEnv := infrastructure.Init()

	// エリア探索を用いて周辺スポットを取得
	areaSpotCollection, err := gobas.spotRepo.
		GetSpotCollectionByCoordinateAndRadiusRepository(
			functionalServerEnv,
			constant.AREA_THRESHOLD,
			gobaidto.Coordinate,
		)
	if err != nil {
		return nil, err
	}

	// 周辺スポットのIDを取得
	spotIdCollection := areaSpotCollection.
		ExtractIdCollection()

	// 周辺スポットを元にスポットに紐づくオブジェクトを取得
	areaObjectCollection, err := gobas.objectRepo.
		GetObjectCollectionBySpotIdsRepository(
			functionalServerEnv,
			spotIdCollection,
			gobaidto.User,
		)
	if err != nil {
		return nil, err
	}

	// 戻り値を生成
	getObjectByAreaOutputDTO := gobaof.Create(
		*gobaidto.User,
		areaObjectCollection,
		areaSpotCollection,
	)
	return &getObjectByAreaOutputDTO, nil

}
