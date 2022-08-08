package models

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthModel struct{}

func (a AuthModel) CreateToken(user *User) (string, error) {
	accssUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(time.Minute * 5).Unix()

	at := jwt.MapClaims{}
	at["access_uuid"] = accssUUID.String()
	at["email"] = user.Email
	at["exp"] = exp

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, err := atoken.SignedString([]byte("SecretCode"))
	if err != nil {
		return "", err
	}

	return signedAuthToken, nil
}

func (a AuthModel) ExtractToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	strArr := strings.Split(authorization, " ")
	if len(strArr) == 2 && strArr[0] == "Bearer" {
		return strArr[1]
	}
	return ""
}

func (a AuthModel) VerifyToken(r *http.Request) (bool, error) {
	tokenString := a.ExtractToken(r)

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("SecretCode"), nil
	}

	token, err := jwt.Parse(tokenString, key)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
