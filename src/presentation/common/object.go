package common

import "mime/multipart"

type ObjectRequest struct {
	Name string                `json:"name" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
	Spot SpotRequest           `json:"spot"`
}

type ObjectResponse struct {
	Id       string       `json:"id" binding:"required,uuid"`
	PosterId string       `json:"posterId" binding:"required,uuid"`
	Spot     SpotResponse `json:"spot"`
	ViewUrl  string       `json:"viewUrl" binding:"required,url"`
}
