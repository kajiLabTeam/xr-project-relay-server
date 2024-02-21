package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-relay-server/presentation/common"
)

type ObjectPostSearchSpotRequest struct {
	Id         string                    `json:"id" binding:"required,uuid"`
	Coordinate *common.CoordinateRequest `json:"coordinate"`
}

type ObjectPostSearchSpotResponse struct {
	Id         string                  `json:"id" binding:"required,uuid"`
	SpotObject common.ObjectResponse   `json:"spotObject"`
	AreaObject []common.ObjectResponse `json:"areaObject"`
}

func ObjectSearchSpotRouter(r *gin.Engine) {
	// スポット推定を用いたオブジェクト探索
	r.POST("api/users/:userId/objects/search/spot", func(c *gin.Context) {
		var req ObjectPostSearchSpotRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
