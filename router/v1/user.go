package v1

import (
	"github.com/JisungPark0319/go-gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	userController := new(controllers.UsersController)

	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)
}
