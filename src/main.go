package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-relay-server/presentation"
)

func main() {
	r := gin.Default()

	presentation.ObjectUploadRouter(r)
	presentation.ObjectSearchSpotRouter(r)
	presentation.ObjectSearchAreaRouter(r)

	r.Run()
}
