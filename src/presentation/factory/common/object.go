package common_factory_presentation

type ObjectResponseDTO struct {
	Id        string       `json:"id" binding:"required,uuid"`
	PosterId  string       `json:"posterId" binding:"required,uuid"`
	Spot      SpotResponseDTO `json:"spot"`
	ViewUrl   string       `json:"viewUrl,omitempty" binding:"required,url"`
	UploadUrl string       `json:"uploadUrl,omitempty" binding:"required"`
}
