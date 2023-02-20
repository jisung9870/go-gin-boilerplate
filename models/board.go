package models

import (
	"context"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"gorm.io/gorm"
)

type Board struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Content   string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BoardModel struct{}

func (b BoardModel) CreateBoard(ctx context.Context, form forms.CreateBoard) (Board, error) {
	var board Board
	db := database.GetDB()

	board.Title = form.Title
	board.Author = form.Author
	board.Content = form.Content

	result := db.WithContext(ctx).Create(&board)
	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}

func (b BoardModel) GetBoardList(ctx context.Context, querys forms.BoardQuery) ([]Board, error) {
	var board []Board
	db := database.GetDB()

	var result *gorm.DB
	if querys.Id != "" {
		result = db.WithContext(ctx).Where("ID = ?", querys.Id).First(&board)
	} else {
		result = db.WithContext(ctx).Limit(querys.Limit).Find(&board)
	}
	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}

func (b BoardModel) DeleteBoard(ctx context.Context, form forms.DeleteBoard) (Board, error) {
	var board Board
	db := database.GetDB()

	board.ID = form.ID
	result := db.WithContext(ctx).Delete(&board)

	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}
