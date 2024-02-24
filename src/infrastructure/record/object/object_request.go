package object_record

type GetObjectBySpotIdRequest struct {
	UserId string `json:"userId"  binding:"required,uuid"`
	SpotId string `json:"spotId"  binding:"required,uuid"`
}

type GetObjectCollectionBySpotIdsRequest struct {
	UserId string   `json:"userId"  binding:"required,uuid"`
	SpotId []string `json:"spotId"  binding:"required,uuid"`
}

type CreateObjectRequest struct {
	UserId    string `json:"userId"  binding:"required,uuid"`
	SpotId    string `json:"spotId"  binding:"required,uuid"`
	Extension string `json:"extension"  binding:"required"`
}
