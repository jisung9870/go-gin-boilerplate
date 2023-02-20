package controllers

import (
	"context"
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BoardController struct{}

var boardModel = new(models.BoardModel)

func (b BoardController) CreateBoard(c *gin.Context) {
	var createBoardFrom forms.CreateBoard

	if err := c.ShouldBindJSON(&createBoardFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtAuth := auth.Get()
	email, err := jwtAuth.GetAccesssClaims(c.GetHeader("Authorization"), "email")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	createBoardFrom.Author = email
	log.Debug(createBoardFrom)

	board, err := boardModel.CreateBoard(context.Background(), createBoardFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"board":   board,
	})
}

func (b BoardController) GetBoardList(c *gin.Context) {
	boardQuery := forms.BoardQuery{}
	if err := c.ShouldBindQuery(&boardQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boards, err := boardModel.GetBoardList(context.Background(), boardQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"boards":  boards,
	})
}

func (b BoardController) DeleteBoard(c *gin.Context) {
	var deleteBoardFrom forms.DeleteBoard

	if err := c.ShouldBindJSON(&deleteBoardFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, err := boardModel.DeleteBoard(context.Background(), deleteBoardFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"board":   board,
	})
}
