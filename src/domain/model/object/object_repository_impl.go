package object_model_domain

import (
	"github.com/kajiLabTeam/xr-project-relay-server/src/config/env"
	user_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/user"
)

type ObjectRepository interface {
	GetObjectBySpotIdRepository(
		functionServerEnv *env.FunctionServerEnv,
		spotId string,
		u *user_model_domain.User,
	) (*Object, error)
	GetObjectCollectionBySpotIdsRepository(
		functionServerEnv *env.FunctionServerEnv,
		spotId []string,
		u *user_model_domain.User,
	) (ObjectCollection, error)
	CreateObjectRepository(
		functionServerEnv *env.FunctionServerEnv,
		userId string,
		extension string,
		spotId string,
	) (*Object, error)
}
