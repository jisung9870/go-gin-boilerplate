package controllers

import (
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/JisungPark0319/go-gin-boilerplate/forms"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

var userModel = new(models.UserModel)

func (u UsersController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var loginForm forms.Login
	jwtAuth := auth.Get()

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userModel.LoginWithContext(ctx, loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := make(auth.Claims)
	claims["email"] = user.Email

	token, err := jwtAuth.Create(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"user":          user,
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

func (u UsersController) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var registerForm forms.Register

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userModel.RegisterWithContext(ctx, registerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
