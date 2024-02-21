package presentation

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-relay-server/presentation/common"
)

type ObjectPostSearchAreaRequest struct {
	Id          string                    `json:"id" binding:"required,uuid"`
	Coordinate  *common.CoordinateRequest `json:"coordinate"`
	RawDataFile *multipart.FileHeader     `form:"rawDataFile" binding:"required"`
}

type ObjectPostSearchAreaResponse struct {
	Id      string                `json:"id" binding:"required,uuid"`
	Objects []common.SpotResponse `json:"objects"`
}

func ObjectSearchAreaRouter(r *gin.Engine) {
	// エリアを用いた周辺オブジェクト探索
	r.POST("api/users/:userId/objects/search/area", func(c *gin.Context) {
		var req ObjectPostSearchAreaRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
