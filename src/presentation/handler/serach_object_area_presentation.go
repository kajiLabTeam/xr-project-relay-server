package handler_presentation

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	common_handler "github.com/kajiLabTeam/xr-project-relay-server/presentation/handler/common"
)

type ObjectPostSearchAreaRequest struct {
	Id          string                    `json:"id" binding:"required,uuid"`
	Coordinate  *common_handler.CoordinateRequest `json:"coordinate"`
	RawDataFile *multipart.FileHeader     `form:"rawDataFile" binding:"required"`
}

type ObjectPostSearchAreaResponse struct {
	Id      string                `json:"id" binding:"required,uuid"`
	Objects []common_handler.SpotResponse `json:"objects"`
}

// エリアを用いた周辺オブジェクト探索
func ObjectSearchAreaRouter(r *gin.Engine) {
	r.POST("api/users/:userId/objects/search/area", func(c *gin.Context) {
		var req ObjectPostSearchAreaRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
