package spot_record

type GetSpotBySpotIdsAndRawDataFileRequest struct {
	SpotIds     []string `json:"spotIds"  binding:"required"`
	RawDataFile []byte   `json:"rawDataFile"  binding:"required"`
}

type GetSpotCollectionByCoordinateAndRadiusRequest struct {
	Latitude  float64 `json:"latitude"  binding:"required"`
	Longitude float64 `json:"longitude"  binding:"required"`
}

type CreateSpotRequest struct {
	Name         string  `json:"name"  binding:"required"`
	Floors       int     `json:"floors"  binding:"required"`
	LocationType string  `json:"locationType"  binding:"required"`
	Latitude     float64 `json:"latitude"  binding:"required"`
	Longitude    float64 `json:"longitude"  binding:"required"`
	RawDataFile  []byte  `json:"rawDataFile"  binding:"required"`
}
