package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-relay-server/src/presentation/handler"
)

func main() {
	r := gin.Default()

	handler.GetObjectBySpotRouter(r)
	handler.GetObjectByAreaRouter(r)
	handler.CreateObjectHandler(r)

	r.Run()
}
