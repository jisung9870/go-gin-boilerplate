package models

import (
	"context"
	"errors"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"github.com/JisungPark0319/go-gin-boilerplate/trace"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;index"`
	Password  string `gorm:"not null"`
	Name      string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserModel struct{}

func (m UserModel) LoginWithContext(ctx context.Context, form forms.Login) (User, error) {
	var user User
	db := database.GetDB()

	result := db.WithContext(ctx).Where("Email = ?", form.Email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	span := trace.Span{SpanName: "Password-Compare"}
	span.Event(ctx)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return user, errors.New("password is wrong")
	}
	span.Close()

	user.Password = ""
	return user, nil
}

func (m UserModel) RegisterWithContext(ctx context.Context, form forms.Register) (User, error) {
	var user User
	db := database.GetDB()

	result := db.WithContext(ctx).Where("Email = ?", form.Email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected > 0 {
		return user, errors.New("email already exists")
	}

	span := trace.Span{SpanName: "generate password"}
	span.Event(ctx)
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	span.Close()

	user.Email = form.Email
	user.Password = string(hashedPassword)
	user.Name = form.Name

	result = db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	user.Password = ""
	return user, nil
}
