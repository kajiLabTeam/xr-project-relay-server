package spot_record

import spot_model_domain "github.com/kajiLabTeam/xr-project-relay-server/src/domain/model/spot"

type SpotResponse struct {
	Id           string  `json:"id"  binding:"required,uuid"`
	Name         string  `json:"name"  binding:"required"`
	Floors       int     `json:"floors"  binding:"required"`
	LocationType string  `json:"locationType"  binding:"required"`
	Latitude     float64 `json:"latitude"  binding:"required"`
	Longitude    float64 `json:"longitude"  binding:"required"`
}

type GetSpotBySpotIdsAndRawDataFileResponse struct {
	Id           string  `json:"id"  binding:"required,uuid"`
	Name         string  `json:"name"  binding:"required"`
	Floors       int     `json:"floors"  binding:"required"`
	LocationType string  `json:"locationType"  binding:"required"`
	Latitude     float64 `json:"latitude"  binding:"required"`
	Longitude    float64 `json:"longitude"  binding:"required"`
}

// TODO : ドメイン層に書くべきかもしれない
func (gsr *GetSpotBySpotIdsAndRawDataFileResponse) ToDomainSpot() (*spot_model_domain.Spot, error) {
	domainCoordinate, err := spot_model_domain.NewCoordinate(
		gsr.Latitude,
		gsr.Longitude,
	)
	if err != nil {
		return nil, err
	}

	domainSpot, err := spot_model_domain.NewSpot(
		gsr.Id,
		gsr.Name,
		gsr.LocationType,
		gsr.Floors,
		domainCoordinate,
	)
	if err != nil {
		return nil, err
	}

	return domainSpot, err
}

type GetSpotCollectionByCoordinateAndRadiusResponse struct {
	Spots []SpotResponse `json:"spots"  binding:"required"`
}

// TODO : ドメイン層に書くべきかもしれない
func (gasr *GetSpotCollectionByCoordinateAndRadiusResponse) ToDomainSpotCollection() (spot_model_domain.SpotCollection, error) {
	var domainSpotCollection spot_model_domain.SpotCollection
	for _, spot := range gasr.Spots {
		coordinate, err := spot_model_domain.NewCoordinate(
			spot.Latitude,
			spot.Longitude,
		)
		if err != nil {
			return nil, err
		}

		domainSpot, err := spot_model_domain.NewSpot(
			spot.Id,
			spot.Name,
			spot.LocationType,
			spot.Floors,
			coordinate,
		)
		if err != nil {
			return nil, err
		}

		domainSpotCollection.AddSpot(domainSpot)
	}

	return domainSpotCollection, nil
}

type CreateSpotResponse struct {
	Id           string  `json:"id"  binding:"required,uuid"`
	Name         string  `json:"name"  binding:"required"`
	Floors       int     `json:"floors"  binding:"required"`
	LocationType string  `json:"locationType"  binding:"required"`
	Latitude     float64 `json:"latitude"  binding:"required"`
	Longitude    float64 `json:"longitude"  binding:"required"`
}

// TODO : ドメイン層に書くべきかもしれない
func (csr *CreateSpotResponse) ToDomainSpot() (*spot_model_domain.Spot, error) {
	domainCoordinate, err := spot_model_domain.NewCoordinate(
		csr.Latitude,
		csr.Longitude,
	)
	if err != nil {
		return nil, err
	}

	domainSpot, err := spot_model_domain.NewSpot(
		csr.Id,
		csr.Name,
		csr.LocationType,
		csr.Floors,
		domainCoordinate,
	)
	if err != nil {
		return nil, err
	}
	return domainSpot, err
}
