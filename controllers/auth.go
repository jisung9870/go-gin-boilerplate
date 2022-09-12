package controllers

import (
	"net/http"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (au AuthController) Refresh(c *gin.Context) {
	jwtAuth := auth.Get()
	tokens := auth.Tokens{}
	tokens.AccessToken = c.GetHeader("Authorization")
	tokens.RefreshToken = c.GetHeader("Authorization-Refresh")

	reTokens, err := jwtAuth.Refresh(tokens)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  reTokens.AccessToken,
		"refresh_token": reTokens.RefreshToken,
	})
}
