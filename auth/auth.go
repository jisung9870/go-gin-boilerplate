package auth

import (
	"fmt"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/google/uuid"
)

type Auth struct {
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var accessToken = Token{
	Secret: "qwerasdf",
}

var refreshToken = Token{
	Secret: "qwerasdf",
}

func CreateAuth(user models.User) (Tokens, error) {
	tokens := Tokens{}

	var err error
	tokens.AccessToken, err = createAccessToken()
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = createRefreshToken()
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func DeleteAuth(tokens Tokens) error {
	return nil
}

func Refresh(tokens Tokens) (Tokens, error) {
	reTokens := Tokens{}

	ok, err := accessToken.Verify(tokens.AccessToken)
	if ok {
		return reTokens, fmt.Errorf("Token is not expired")
	}
	if err != nil && err.Error() != "Token is expired" {
		return reTokens, err
	}

	ok, err = refreshToken.Verify(tokens.RefreshToken)
	if err != nil {
		return reTokens, err
	} else if !ok {
		return reTokens, fmt.Errorf("refresh token verification failed")
	}

	if err = DeleteAuth(tokens); err != nil {
		return reTokens, err
	}

	reTokens.AccessToken, err = createAccessToken()
	if err != nil {
		return reTokens, err
	}
	reTokens.RefreshToken, _ = createRefreshToken()
	if err != nil {
		return reTokens, err
	}

	return reTokens, nil
}

func VerifyAccessToken(token string) (bool, error) {
	if ok, err := accessToken.Verify(token); err != nil || !ok {
		return false, err
	}
	return true, nil
}

func createAccessToken() (string, error) {
	metadata := make(map[string]interface{})

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(time.Minute * 1).Unix()

	metadata["access_uuid"] = uuid.String()
	metadata["exp"] = expire

	tk, err := accessToken.CreateToken(metadata)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func createRefreshToken() (string, error) {
	metadata := make(map[string]interface{})

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(time.Hour * 1).Unix()

	metadata["refresh_uuid"] = uuid.String()
	metadata["exp"] = expire

	tk, err := refreshToken.CreateToken(metadata)
	if err != nil {
		return "", err
	}
	return tk, nil
}
