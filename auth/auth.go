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
	tokens.AccessToken, err = createAccessToken(user)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = createRefreshToken(user)
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
	if err != nil {
		return reTokens, err
	} else if ok {
		return reTokens, fmt.Errorf("access token not expired")
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

	accessMetadata, _ := accessToken.ExtractTokenMetaData(tokens.AccessToken)
	atExpire := time.Now().Add(time.Minute * 15).Unix()
	accessMetadata["exp"] = atExpire
	reTokens.AccessToken, _ = accessToken.CreateToken(accessMetadata)

	refreshMetadata, _ := refreshToken.ExtractTokenMetaData(tokens.RefreshToken)
	rtExpire := time.Now().Add(time.Minute * 1).Unix()
	refreshMetadata["exp"] = rtExpire
	reTokens.RefreshToken, _ = refreshToken.CreateToken(refreshMetadata)

	return reTokens, nil
}

func VerifyAccessToken(token string) (bool, error) {
	if ok, err := accessToken.Verify(token); err != nil || !ok {
		return false, err
	}
	return true, nil
}

func createAccessToken(user models.User) (string, error) {
	metadata := make(map[string]interface{})

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(time.Minute * 15).Unix()

	metadata["access_uuid"] = uuid.String()
	metadata["exp"] = expire
	metadata["email"] = user.Email
	metadata["user_id"] = user.ID

	tk, err := accessToken.CreateToken(metadata)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func createRefreshToken(user models.User) (string, error) {
	metadata := make(map[string]interface{})

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(time.Hour * 1).Unix()

	metadata["refresh_uuid"] = uuid.String()
	metadata["exp"] = expire
	metadata["email"] = user.Email
	metadata["user_id"] = user.ID

	tk, err := accessToken.CreateToken(metadata)
	if err != nil {
		return "", err
	}
	return tk, nil
}
