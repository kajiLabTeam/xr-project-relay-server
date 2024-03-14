package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/kajiLabTeam/xr-project-relay-server/application/services"
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	user_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/user"
	"github.com/kajiLabTeam/xr-project-relay-server/infrastructure/repository"
	common_handler "github.com/kajiLabTeam/xr-project-relay-server/presentation/handlers/common"
	"github.com/kajiLabTeam/xr-project-relay-server/presentation/middleware"
)

type GetObjectBySpotRequest struct {
	UserId      string                `form:"userId" binding:"required"`
	Latitude    float64               `form:"latitude" binding:"required"`
	Longitude   float64               `form:"longitude" binding:"required"`
	RawDataFile *multipart.FileHeader `form:"rawDataFile" binding:"required"`
}

type GetObjectBySpotResponse struct {
	UserId      string                      `json:"userId" binding:"required"`
	SpotObject  common_handler.ViewObject   `json:"spotObject" binding:"required"`
	AreaObjects []common_handler.ViewObject `json:"areaObjects" binding:"required"`
}

func GetObjectBySpotHandler(r *gin.Engine) {
	r.POST("api/objects/search/spot", func(c *gin.Context) {
		var req GetObjectBySpotRequest

		viewObjectCollectionFactory := common_handler.ViewObjectCollectionFactory{}
		applicationFactory := application_models_domain.ApplicationFactory{}
		getObjectBySpotService := service.NewGetObjectBySpotService(
			repository.NewObjectRepository(),
			repository.NewSpotRepository(),
		)

		header := c.GetHeader("Authorization")
		application, err := applicationFactory.Create(header)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// リクエストのバリデーション
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ファイルを []byte に変換
		rawDataFile, err := middleware.GetBytesFromMultiPartFile(req.RawDataFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := user_models_domain.NewUser(req.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// サービスを実行
		resUser, resObject, resObjectCollection, err := getObjectBySpotService.Run(
			req.Latitude,
			req.Longitude,
			rawDataFile,
			user,
			application,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if resObjectCollection == nil && resObject == nil {
			res := GetObjectBySpotResponse{
				UserId:      resUser.GetId(),
				SpotObject:  common_handler.ViewObject{},
				AreaObjects: []common_handler.ViewObject{},
			}
			c.JSON(http.StatusNotFound, res)
			return
		}

		if resObjectCollection == nil && resObject != nil {
			res := GetObjectBySpotResponse{
				UserId: resObject.GetId(),
				SpotObject: common_handler.ViewObject{
					Id:       resObject.GetId(),
					PosterId: resObject.GetPosterId(),
					Spot: common_handler.Spot{
						Id:           resObject.GetSpot().GetId(),
						Name:         resObject.GetSpot().GetName(),
						Floor:        resObject.GetSpot().GetFloor(),
						LocationType: resObject.GetSpot().GetLocationType(),
						Latitude:     resObject.GetSpot().GetCoordinate().GetLatitude(),
						Longitude:    resObject.GetSpot().GetCoordinate().GetLongitude(),
					},
					ViewUrl: resObject.GetPreSignedUrl(),
				},
				AreaObjects: []common_handler.ViewObject{},
			}

			c.JSON(http.StatusNotFound, res)
			return
		}

		// レスポンスを生成
		res := GetObjectBySpotResponse{
			UserId: resUser.GetId(),
			SpotObject: common_handler.ViewObject{
				Id:       resObject.GetId(),
				PosterId: resObject.GetPosterId(),
				Spot: common_handler.Spot{
					Id:           resObject.GetSpot().GetId(),
					Name:         resObject.GetSpot().GetName(),
					Floor:        resObject.GetSpot().GetFloor(),
					LocationType: resObject.GetSpot().GetLocationType(),
					Latitude:     resObject.GetSpot().GetCoordinate().GetLatitude(),
					Longitude:    resObject.GetSpot().GetCoordinate().GetLongitude(),
				},
				ViewUrl: resObject.GetPreSignedUrl(),
			},
			AreaObjects: viewObjectCollectionFactory.FromViewObjectCollection(
				resObjectCollection,
			),
		}

		// レスポンスを返却
		c.JSON(http.StatusOK, res)
	})
}