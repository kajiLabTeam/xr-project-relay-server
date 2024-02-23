package infrastructure

import "github.com/kajiLabTeam/xr-project-relay-server/config/env"

func Init() *env.FunctionServerEnv {
	fse, err := env.SetObjectServiceUrl()
	if err != nil {
		panic(err)
	}
	return fse
}
