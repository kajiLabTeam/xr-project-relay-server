package spot_model_domain

type SpotCollection []Spot

func (sc *SpotCollection) AddSpot(s *Spot) {
	*sc = append(*sc, *s)
}

func (sc *SpotCollection) ExtractIdCollection() []string {
	var idCollection []string
	for _, spot := range *sc {
		idCollection = append(idCollection, spot.id)
	}
	return idCollection
}
