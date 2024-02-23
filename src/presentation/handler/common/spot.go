package common_handler

import "mime/multipart"

type SpotRequest struct {
	Name         string                `json:"name"`
	LocationType string                `json:"locationType"`
	Floors       int                   `json:"floors"`
	RawDataFile  *multipart.FileHeader `form:"rawDataFile" binding:"required"`
	Coordinate   CoordinateRequest     `json:"coordinate"`
}

type SpotResponse struct {
	Id           string            `json:"id" binding:"required,uuid"`
	Name         string            `json:"name"`
	LocationType string            `json:"locationType"`
	Floors       int               `json:"floors"`
	Coordinate   CoordinateRequest `json:"coordinate"`
}
