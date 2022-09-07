package v1

import (
	"github.com/JisungPark0319/go-gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	authController := new(controllers.AuthController)

	router.POST("/refresh", authController.Refresh)
}
