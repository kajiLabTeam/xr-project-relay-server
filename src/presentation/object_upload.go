package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-relay-server/presentation/common"
)

type ObjectPostUploadRequest struct {
	Id     string               `json:"id"`
	Object common.ObjectRequest `json:"object"`
}

type ObjectPostUploadResponse struct {
	Id     string                `json:"id" binding:"required,uuid"`
	Object common.ObjectResponse `json:"object"`
}

func ObjectUploadRouter(r *gin.Engine) {
	// オブジェクトのアップロード
	r.POST("api/users/:userId/objects/upload", func(c *gin.Context) {
		var req ObjectPostUploadRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
