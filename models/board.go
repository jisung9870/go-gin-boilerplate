package models

import (
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
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

func (b BoardModel) CreateBoard(form forms.CreateBoard) (Board, error) {
	var board Board
	db := database.GetDB()

	board.Title = form.Title
	board.Author = form.Author
	board.Content = form.Content

	result := db.Create(&board)
	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}
