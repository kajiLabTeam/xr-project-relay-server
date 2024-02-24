package input_factory_presentation

import common_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/common"

type PostObjectCreateRequestDTO struct {
	UserId    string                                     `json:"userId" binding:"required,uuid"`
	Extension string                                     `json:"extension"`
	Spot      common_factory_presentation.SpotRequestDTO `json:"spot"`
}
