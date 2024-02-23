package object_model_domain

import "github.com/kajiLabTeam/xr-project-relay-server/config/env"

type ObjectRepository interface {
	CreateObject(functionServerEnv *env.FunctionServerEnv, o *Object) (*Object, error)
	GetObjectsBySpotId(functionServerEnv *env.FunctionServerEnv, coordinate *Object) ([]*Object, error)
}
