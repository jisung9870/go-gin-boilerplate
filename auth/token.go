package auth

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Secret string
}

func (tk *Token) CreateToken(metadata map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}

	for key, val := range metadata {
		claims[key] = val
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwt.SignedString([]byte(tk.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (tk *Token) ExtractToken(bearToken string) (*jwt.Token, error) {
	strArr := strings.Split(bearToken, " ")
	fmt.Println(strArr)
	if len(strArr) != 2 || !(strArr[0] == "Bearer" || strArr[0] == "bearer") {
		return nil, fmt.Errorf("bearer token not in format")
	}
	tokenString := strArr[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tk.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (tk *Token) Verify(bearToken string) (bool, error) {
	token, err := tk.ExtractToken(bearToken)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (tk *Token) ExtractTokenMetaData(bearToken string) (map[string]interface{}, error) {
	token, err := tk.ExtractToken(bearToken)

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("extraction of token information failed")
	}
}
