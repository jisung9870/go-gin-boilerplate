package models

import (
	"errors"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
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

func (m UserModel) Login(form forms.Login) (User, error) {
	var user User
	db := database.GetDB()

	result := db.Where("Email = ?", form.Email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return user, errors.New("password is wrong")
	}

	user.Password = ""
	return user, nil
}

func (m UserModel) Register(form forms.Register) (User, error) {
	var user User
	db := database.GetDB()

	result := db.Where("Email = ?", form.Email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	user.Email = form.Email
	user.Password = string(hashedPassword)
	user.Name = form.Name

	result = db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	user.Password = ""
	return user, nil
}
