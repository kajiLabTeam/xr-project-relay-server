package spot_model_domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Spot struct {
	id           string
	name         string
	locationType string
	floors       int
	Coordinate   Coordinate
	// RawDataFile  []byte
}

func NewSpot(id string, name string, locationType string, floors int, coordinate *Coordinate) (*Spot, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, fmt.Errorf("invalid id value")
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("invalid name value")
	}

	if locationType != "indoor" && locationType != "outdoor" {
		return nil, fmt.Errorf("invalid locationType value")
	}

	if floors < 0 {
		return nil, fmt.Errorf("invalid floors value")
	}

	if coordinate == nil {
		return nil, fmt.Errorf("invalid coordinate value")
	}

	// if rowDataFile == nil {
	// 	return nil, fmt.Errorf("invalid rowDataFile value")
	// }

	return &Spot{
		id:           id,
		name:         name,
		locationType: locationType,
		floors:       floors,
		Coordinate:   *coordinate,
		// RawDataFile:  rowDataFile,
	}, nil
}

func (s *Spot) GetId() string {
	return s.id
}

func (s *Spot) GetName() string {
	return s.name
}

func (s *Spot) GetLocationType() string {
	return s.locationType
}

func (s *Spot) GetFloors() int {
	return s.floors
}

func (s *Spot) GetCoordinate() *Coordinate {
	return &s.Coordinate
}

// func (s *Spot) GetRawDataFile() []byte {
// 	return s.RawDataFile
// }
