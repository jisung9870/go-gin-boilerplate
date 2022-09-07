package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
}
