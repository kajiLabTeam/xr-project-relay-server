package common_factory_application

type SpotInputDTO struct {
	Name         string
	Floors       int
	LocationType string
	RawDataFile  []byte
	Latitude     float64
	Longitude    float64
}

type SpotOutputDTO struct {
	Id           string
	Name         string
	Floors       int
	LocationType string
	Latitude     float64
	Longitude    float64
}

type SpotOutputDTOCollection []SpotOutputDTO
