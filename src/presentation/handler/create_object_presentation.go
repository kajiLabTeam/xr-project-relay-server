package handler_presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common_handler "github.com/kajiLabTeam/xr-project-relay-server/presentation/handler/common"
)

type ObjectPostUploadRequest struct {
	Id     string                       `json:"id"`
	Object common_handler.ObjectRequest `json:"object"`
}

type ObjectPostUploadResponse struct {
	Id     string                        `json:"id" binding:"required,uuid"`
	Object common_handler.ObjectResponse `json:"object"`
}

// オブジェクトのアップロード
func CreateObjectHandler(r *gin.Engine) {
	r.POST("api/users/:userId/objects/upload", func(c *gin.Context) {
		var req ObjectPostUploadRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
