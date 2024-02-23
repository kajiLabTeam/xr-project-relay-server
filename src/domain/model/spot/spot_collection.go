package spot_model_domain

type SpotCollection []Spot

func (sc *SpotCollection) AddSpot(id, name, locationType string, floors int, coordinate Coordinate) error {
	newSpot, err := NewSpot(id, name, locationType, floors, &coordinate)
	if err != nil {
		return err
	}
	*sc = append(*sc, *newSpot)

	return nil
}
