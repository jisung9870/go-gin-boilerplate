package router

import (
	v1 "github.com/JisungPark0319/go-gin-boilerplate/router/v1"
	"github.com/gin-gonic/gin"
)

func Set(router *gin.Engine) {
	routerV1 := router.Group("api/v1")
	{
		v1.HealthCheck(routerV1.Group("/health"))
		v1.User(routerV1.Group("/user"))
		v1.Auth(routerV1.Group("/auth"))
		v1.Board(routerV1.Group("/board"))
	}
}
