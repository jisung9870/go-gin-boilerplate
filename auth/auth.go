package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	access        Token
	refresh       Token
	accessExpire  time.Duration
	refreshExpire time.Duration
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type tokenDetail struct {
	token  Token
	expire time.Duration
}

type Claims map[string]interface{}

var auth *Auth

func New(accessSecret, refreshSecret string) error {
	if auth != nil {
		return errors.New("Auth instance exists")
	}
	auth = &Auth{
		access:        Token{Secret: accessSecret},
		refresh:       Token{Secret: refreshSecret},
		accessExpire:  time.Minute * 10,
		refreshExpire: time.Hour * 1,
	}
	return nil
}

func Get() *Auth {
	return auth
}

func (a *Auth) SetExpire(accessExpire, refreshExpire time.Duration) {
	a.accessExpire = accessExpire
	a.refreshExpire = refreshExpire
}

func (a Auth) Create(claims Claims) (Tokens, error) {
	tokens := Tokens{}

	var err error

	tokens.AccessToken, err = a.CreateToken("access", claims)
	if err != nil {
		return tokens, err
	}
	tokens.AccessToken = "Bearer " + tokens.AccessToken

	tokens.RefreshToken, err = a.CreateToken("refresh", claims)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = "Bearer " + tokens.RefreshToken

	// TODO: Store info in Database

	return tokens, nil
}

func (a Auth) Refresh(tokens Tokens) (Tokens, error) {
	ok, err := a.access.Verify(tokens.AccessToken)
	if err != nil || ok {
		if ok {
			err = errors.New("Token has not expired")
		}
		return Tokens{}, err
	}

	ok, err = a.refresh.Verify(tokens.RefreshToken)
	if err != nil || !ok {
		if !ok {
			err = errors.New("Token is expired")
		}
		return Tokens{}, err
	}

	accessClaims, err := a.access.ExtractTokenMetaData(tokens.AccessToken)
	if err != nil {
		return Tokens{}, err
	}
	refreshClaims, err := a.refresh.ExtractTokenMetaData(tokens.RefreshToken)
	if err != nil {
		return Tokens{}, err
	}
	if accessClaims["access_uuid"] != refreshClaims["access_uuid"] {
		return Tokens{}, errors.New("claim values for access token and refresh token are different")
	}
	// TODO: Compare with database info

	reTokens, err := a.Create(accessClaims)
	if err != nil {
		return Tokens{}, err
	}

	return reTokens, nil
}

func (a Auth) Verify(tokenType string, tokenStr string) (bool, error) {
	token := a.tokenSelect(tokenType).token

	ok, err := token.Verify(tokenStr)
	return ok, err
}

func (a Auth) ExtractClaims(tokenType string, tokenStr string) (Claims, error) {
	token := a.tokenSelect(tokenType).token

	claims, err := token.ExtractTokenMetaData(tokenStr)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// TODO: Delete info in Database
// func DeleteAuth(tokens Tokens) error {
// 	return nil
// }

func (a Auth) CreateToken(tokenType string, claims Claims) (string, error) {
	td := a.tokenSelect(tokenType)

	uuid := uuid.New()

	claims[tokenType+"_uuid"] = uuid.String()
	claims["exp"] = time.Now().Add(td.expire).Unix()

	tk, err := td.token.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return tk, nil
}

func (a Auth) tokenSelect(tokenType string) tokenDetail {
	var td tokenDetail
	switch tokenType {
	case "access":
		td.token = a.access
		td.expire = a.accessExpire
	case "refresh":
		td.token = a.refresh
		td.expire = a.refreshExpire
	}
	return td
}
