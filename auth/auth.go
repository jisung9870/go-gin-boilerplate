package auth

import (
	"errors"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/config"
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

func New(cfg config.AuthConfig) error {
	if auth != nil {
		return errors.New("Auth instance exists")
	}
	auth = &Auth{
		access:        Token{Secret: cfg.AccessSecret},
		refresh:       Token{Secret: cfg.RefreshSecret},
		accessExpire:  time.Minute * 10,
		refreshExpire: time.Hour * 1,
	}
	return nil
}

func Get() *Auth {
	return auth
}

// Token expiration date setting
func (a *Auth) SetExpire(accessExpire, refreshExpire time.Duration) {
	a.accessExpire = accessExpire
	a.refreshExpire = refreshExpire
}

// Create auth token(access, refresh)
func (a Auth) Create(claims Claims) (Tokens, error) {
	tokens := Tokens{}

	var err error

	tokens.AccessToken, err = a.CreateAccessToken(claims)
	if err != nil {
		return tokens, err
	}
	tokens.AccessToken = "Bearer " + tokens.AccessToken

	tokens.RefreshToken, err = a.CreateRefreshToken(claims)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = "Bearer " + tokens.RefreshToken

	// TODO: Store info in Database

	return tokens, nil
}

// Token Refresh
func (a Auth) Refresh(tokens Tokens) (Tokens, error) {
	ok, err := a.AccessVerify(tokens.AccessToken)
	if err != nil || ok {
		if ok {
			err = errors.New("Token has not expired")
		}
		return Tokens{}, err
	}

	ok, err = a.RefreshVerify(tokens.RefreshToken)
	if err != nil || !ok {
		if !ok {
			err = errors.New("Token is expired")
		}
		return Tokens{}, err
	}

	accessClaims, err := a.AccessExtractClaims(tokens.AccessToken)
	if err != nil {
		return Tokens{}, err
	}
	refreshClaims, err := a.RefreshExtractClaims(tokens.RefreshToken)
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

// TODO: Delete info in Database
// func DeleteAuth(tokens Tokens) error {
// 	return nil
// }

/*********************************
Abstraction functions
*********************************/
// Create access token abstract
func (a Auth) CreateAccessToken(claims Claims) (string, error) {
	return a.createToken("access", claims)
}

// Create refresh token abstract
func (a Auth) CreateRefreshToken(claims Claims) (string, error) {
	return a.createToken("refresh", claims)
}

// Access token verfication abstract
func (a Auth) AccessVerify(tokenStr string) (bool, error) {
	return a.verify("access", tokenStr)
}

// Refresh token verification abstract
func (a Auth) RefreshVerify(tokenStr string) (bool, error) {
	return a.verify("refresh", tokenStr)
}

// Extract claim data from access token
func (a Auth) AccessExtractClaims(tokenStr string) (Claims, error) {
	return a.extractClaims("access", tokenStr)
}

// Extract claim data from refresh token
func (a Auth) RefreshExtractClaims(tokenStr string) (Claims, error) {
	return a.extractClaims("refresh", tokenStr)
}

func (a Auth) GetAccesssClaims(jwt string, key string) (string, error) {
	claims, err := a.AccessExtractClaims(jwt)
	if err != nil {
		return "", err
	}

	return claims[key].(string), nil
}

// Token expiration verify
func (a Auth) verify(tokenType string, tokenStr string) (bool, error) {
	token := a.tokenSelect(tokenType).token

	ok, err := token.Verify(tokenStr)
	return ok, err
}

// Extract claim data forn token
func (a Auth) extractClaims(tokenType string, tokenStr string) (Claims, error) {
	token := a.tokenSelect(tokenType).token

	claims, err := token.ExtractTokenMetaData(tokenStr)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// Create token
func (a Auth) createToken(tokenType string, claims Claims) (string, error) {
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

// Select structure to use between access, refresh token
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
