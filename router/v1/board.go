package v1

import (
	"github.com/JisungPark0319/go-gin-boilerplate/controllers"
	"github.com/JisungPark0319/go-gin-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func Board(router *gin.RouterGroup) {
	boardController := new(controllers.BoardController)
	router.Use(middlewares.Authorization())

	router.POST("/", boardController.CreateBoard)
	router.GET("/", boardController.GetBoardList)
	router.DELETE("/", boardController.DeleteBoard)
}
