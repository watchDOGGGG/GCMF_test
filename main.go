package main

import (
	service "gcmf-services/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/test", service.TestApp())
	r.POST("/api/verify_account", service.Verifyuseraccount())

	r.Run(":8787")
}
