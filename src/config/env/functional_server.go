package env

import (
	"os"

	"github.com/joho/godotenv"
)

type FunctionServerEnv struct {
	ObjectServiceURL         string
	SpotEstimationServiceURL string
	AreaEstimationServiceURL string
	ApplicationServiceURL    string
	RegistrationServiceURL   string
}

func SetObjectServiceUrl() (*FunctionServerEnv, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	return &FunctionServerEnv{
		ObjectServiceURL:         os.Getenv("OBJECT_SERVICE_URL"),
		SpotEstimationServiceURL: os.Getenv("SPOT_ESTIMATION_SERVICE_URL"),
		AreaEstimationServiceURL: os.Getenv("AREA_ESTIMATION_SERVICE_URL"),
		ApplicationServiceURL:    os.Getenv("APPLICATION_SERVICE_URL"),
		RegistrationServiceURL:   os.Getenv("REGISTRATION_SERVICE_URL"),
	}, nil
}

func (fsc *FunctionServerEnv) GetObjectServiceUrl() string {
	return fsc.ObjectServiceURL
}

func (fsc *FunctionServerEnv) GetSpotEstimationServiceUrl() string {
	return fsc.SpotEstimationServiceURL
}

func (fsc *FunctionServerEnv) GetAreaEstimationServiceUrl() string {
	return fsc.AreaEstimationServiceURL
}

func (fsc *FunctionServerEnv) GetApplicationServiceUrl() string {
	return fsc.ApplicationServiceURL
}

func (fsc *FunctionServerEnv) GetRegistrationServiceUrl() string {
	return fsc.RegistrationServiceURL
}
