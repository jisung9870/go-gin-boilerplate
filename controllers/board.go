package controllers

import (
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type BoardController struct{}

var boardModel = new(models.BoardModel)

func (b BoardController) CreateBoard(c *gin.Context) {
	var createBoardFrom forms.CreateBoard

	if err := c.ShouldBindJSON(&createBoardFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, err := boardModel.CreateBoard(createBoardFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"board":   board,
	})
}
