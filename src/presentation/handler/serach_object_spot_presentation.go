package handler_presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common_handler "github.com/kajiLabTeam/xr-project-relay-server/presentation/handler/common"
)

type ObjectPostSearchSpotRequest struct {
	Id         string                            `json:"id" binding:"required,uuid"`
	Coordinate *common_handler.CoordinateRequest `json:"coordinate"`
}

type ObjectPostSearchSpotResponse struct {
	Id         string                          `json:"id" binding:"required,uuid"`
	SpotObject common_handler.ObjectResponse   `json:"spotObject"`
	AreaObject []common_handler.ObjectResponse `json:"areaObject"`
}

// スポット推定を用いたオブジェクト探索
func ObjectSearchSpotRouter(r *gin.Engine) {
	r.POST("api/users/:userId/objects/search/spot", func(c *gin.Context) {
		var req ObjectPostSearchSpotRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
