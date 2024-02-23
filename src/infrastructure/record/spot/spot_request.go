package spot_record

import (
	"os"
)

type CreateSpotRequest struct {
	Name         string            `json:"name"  binding:"required"`
	Floors       int               `json:"floors"  binding:"required"`
	LocationType string            `json:"locationType"  binding:"required"`
	Coordinate   CoordinateRequest `json:"coordinate"  binding:"required"`
	RawDataFile      *os.File          `json:"rawDataFile"  binding:"required"`
}

type GetSpotRequest struct {
	Ids     []string `json:"ids"  binding:"required"`
	RawDataFile *os.File `json:"rawDataFile"  binding:"required"`
}

type GetAreaSpotRequest struct {
	Coordinate CoordinateRequest `json:"coordinate"  binding:"required"`
}

type CoordinateRequest struct {
	Latitude  float64 `json:"latitude"  binding:"required"`
	Longitude float64 `json:"longitude"  binding:"required"`
}
