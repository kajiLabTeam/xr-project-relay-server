package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	input_factory_application "github.com/kajiLabTeam/xr-project-relay-server/src/application/factory/input"
	"github.com/kajiLabTeam/xr-project-relay-server/src/application/service"
	input_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/input"
	output_factory_presentation "github.com/kajiLabTeam/xr-project-relay-server/src/presentation/factory/output"
)

var createObjectService service.CreateObjectService
var createObjectInputFactory input_factory_application.CreateObjectInputFactory
var postCreateObjectResponseFactory output_factory_presentation.PostCreateObjectResponseFactory

// オブジェクトのアップロード
func CreateObjectHandler(r *gin.Engine) {
	r.POST("api/users/:userId/objects/upload", func(c *gin.Context) {
		var req input_factory_presentation.PostObjectCreateRequestDTO

		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// アプリケーション層に渡すDTOを生成
		createObjectInputDTO := createObjectInputFactory.Create(&req)

		// サービスを実行
		createObjectOutputDTO, err := createObjectService.Run(*createObjectInputDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// レスポンスを生成
		postCreateObjectResponseDTO := postCreateObjectResponseFactory.FromCreateObjectOutputDTO(createObjectOutputDTO)

		// レスポンスを返却
		c.JSON(http.StatusCreated, postCreateObjectResponseDTO)
	})
}
