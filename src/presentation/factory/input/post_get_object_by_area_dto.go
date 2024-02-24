package input_factory_presentation

type PostGetObjectByAreaRequestDTO struct {
	UserId    string  `json:"userId" binding:"required,uuid"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
