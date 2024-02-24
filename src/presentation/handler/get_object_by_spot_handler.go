package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	"github.com/kajiLabTeam/xr-project-relay-server/src/application/service"
	input_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/input"
	output_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/output"
)

var getObjectBySpotService service.GetObjectBySpotService
var getObjectBySpotInputFactory input_factory_application.GetObjectBySpotInputFactory
var postGetObjectBySpotResponseFactory output_factory_presentation.PostGetObjectBySpotResponseFactory

// スポット推定を用いたオブジェクト探索
func GetObjectBySpotRouter(r *gin.Engine) {
	r.POST("api/users/:userId/objects/search/spot", func(c *gin.Context) {
		var req input_factory_presentation.PostGetObjectBySpotRequestDTO

		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// リクエストから生データファイルを取得
		rawDataFile, _, err := c.Request.FormFile("rawDataFile")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer rawDataFile.Close()

		// アプリケーション層に渡すDTOを生成
		getObjectBySpotInputDTO, err := getObjectBySpotInputFactory.Create(
			req.UserId,
			req.Latitude,
			req.Longitude,
			rawDataFile,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// サービスを実行
		getObjectBySpotOutputDTO, err := getObjectBySpotService.
			Run(*getObjectBySpotInputDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// レスポンスを生成
		postGetObjectBySpotResponseDTO := postGetObjectBySpotResponseFactory.FromGetObjectBySpotOutputDTO(*getObjectBySpotOutputDTO)

		// レスポンスを返却
		c.JSON(http.StatusOK, postGetObjectBySpotResponseDTO)
	})
}
