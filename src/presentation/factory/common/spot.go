package common_factory_presentation

import "mime/multipart"

type SpotRequestDTO struct {
	Name         string                `json:"name"`
	Floors       int                   `json:"floors"`
	LocationType string                `json:"locationType"`
	RawDataFile  *multipart.FileHeader `form:"rawDataFile" binding:"required"`
	Latitude     float64               `json:"latitude"`
	Longitude    float64               `json:"longitude"`
}

type SpotResponseDTO struct {
	Id           string  `json:"id" binding:"required,uuid"`
	Name         string  `json:"name"`
	LocationType string  `json:"locationType"`
	Floors       int     `json:"floors"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
