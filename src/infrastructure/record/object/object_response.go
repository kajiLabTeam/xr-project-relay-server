package object_record

type GetObjectResponse struct {
	Id     string         `json:"id"  binding:"required,uuid"`
	Object ObjectResponse `json:"object"  binding:"required"`
}

type CreateObjectResponse struct {
	Id     string         `json:"id"  binding:"required,uuid"`
	Object ObjectResponse `json:"object"  binding:"required"`
}

type ObjectResponse struct {
	Id        string `json:"id"  binding:"required,uuid"`
	PosterId  string `json:"posterId"  binding:"required,uuid"`
	SpotId    string `json:"spotId"  binding:"required,uuid"`
	UploadUrl string `json:"uploadUrl,omitempty"  binding:"required"`
	ViewUrl   string `json:"viewUrl,omitempty"  binding:"required"`
}
