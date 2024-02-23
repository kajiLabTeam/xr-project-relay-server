package spot_record

import spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/spot"

type CreateSpotResponse struct {
	SpotResponse `json:"spot"  binding:"required"`
}

type GetSpotResponse struct {
	SpotResponse `json:"spot"  binding:"required"`
}

type GetAreaSpotResponse struct {
	Spots []SpotResponse `json:"spots"  binding:"required"`
}

type SpotResponse struct {
	Id           string             `json:"id"  binding:"required,uuid"`
	Name         string             `json:"name"  binding:"required"`
	Floors       int                `json:"floors"  binding:"required"`
	LocationType string             `json:"locationType"  binding:"required"`
	Coordinate   CoordinateResponse `json:"coordinate"  binding:"required"`
}

type CoordinateResponse struct {
	Latitude  float64 `json:"latitude"  binding:"required"`
	Longitude float64 `json:"longitude"  binding:"required"`
}

// TODO : ドメイン層に書くべきかもしれない
func (gasr *GetAreaSpotResponse) ToDomainSpotCollection() (spot_model_domain.SpotCollection, error) {
	var domainSpotCollection spot_model_domain.SpotCollection
	for _, spot := range gasr.Spots {
		coordinate, err := spot_model_domain.NewCoordinate(
			spot.Coordinate.Latitude,
			spot.Coordinate.Longitude,
		)
		if err != nil {
			return nil, err
		}

		domainSpotCollection.AddSpot(spot.Id, spot.Name, spot.LocationType, spot.Floors, *coordinate)
	}

	return domainSpotCollection, nil
}
