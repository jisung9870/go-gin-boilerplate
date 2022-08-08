package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(router *gin.RouterGroup) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
}
