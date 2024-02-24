package common_factory_application

type ObjectDTO struct {
	Id       string
	PosterId string
	Spot     SpotOutputDTO
	ViewUrl  string
}

type ObjectDTOCollection []ObjectDTO
