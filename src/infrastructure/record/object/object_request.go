package object_record

type CreateObjectRequest struct {
	Id     string        `json:"id"  binding:"required,uuid"`
	Object ObjectRequest `json:"object"  binding:"required"`
}

type GetObjectsBySpotIdsRequest struct {
	UserId string   `json:"userId"  binding:"required,uuid"`
	SpotId []string `json:"spotId"  binding:"required,uuid"`
}

type ObjectRequest struct {
	Id     string `json:"id"  binding:"required,uuid"`
	SpotId string `json:"spotId"  binding:"required,uuid"`
}
