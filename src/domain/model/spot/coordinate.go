package spot_model_domain

import "fmt"

type Coordinate struct {
	latitude  float64
	longitude float64
}

func NewCoordinate(lat float64, log float64) (*Coordinate, error) {
	if lat >= -90 && lat <= 90 {
		return nil, fmt.Errorf("invalid latitude value")
	}

	if log >= -180 && log <= 180 {
		return nil, fmt.Errorf("invalid longitude value")
	}

	return &Coordinate{
		latitude:  lat,
		longitude: log,
	}, nil
}

func (c *Coordinate) GetLatitude() float64 {
	return c.latitude
}

func (c *Coordinate) GetLongitude() float64 {
	return c.longitude
}
