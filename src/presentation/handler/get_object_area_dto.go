package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	"github.com/kajiLabTeam/xr-project-relay-server/src/application/service"
	input_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/input"
	output_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/output"
)

var getObjectByAreaService service.GetObjectByAreaService
var getObjectByAreaInputFactory input_factory_application.GetObjectByAreaInputFactory
var postGetObjectByAreaResponseFactory output_factory_presentation.PostGetObjectByAreaResponseFactory

// エリアを用いた周辺オブジェクト探索
func GetObjectByAreaRouter(r *gin.Engine) {
	r.POST("api/users/:userId/objects/search/area", func(c *gin.Context) {
		var req input_factory_presentation.PostGetObjectByAreaRequestDTO

		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// アプリケーション層に渡すDTOを生成
		getObjectByAreaInputDTO, err := getObjectByAreaInputFactory.Create(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// サービスを実行
		getObjectByAreaOutputDTO, err := getObjectByAreaService.Run(*getObjectByAreaInputDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// レスポンスを生成
		postGetObjectByAreaResponseDTO := postGetObjectByAreaResponseFactory.
			FromGetObjectByAreaOutputDTO(*getObjectByAreaOutputDTO)

		// レスポンスを返却
		c.JSON(http.StatusOK, postGetObjectByAreaResponseDTO)
	})
}
