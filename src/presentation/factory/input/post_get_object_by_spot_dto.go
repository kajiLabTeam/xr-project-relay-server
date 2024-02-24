package input_factory_presentation

import "mime/multipart"

type PostGetObjectBySpotRequestDTO struct {
	UserId      string          `json:"userId" binding:"required,uuid"`
	Latitude    float64         `json:"latitude" binding:"required"`
	Longitude   float64         `json:"longitude" binding:"required"`
	RawDataFile *multipart.File `form:"rawDataFile" binding:"required"`
}
