package jwt

import (
	"errors"
	"fmt"

	jwtPkg "github.com/golang-jwt/jwt/v4"
)

type JWT interface {
	CreateTokenString(userId uint, email string) (string, error)
	ExtractTokenData(tokenString string) (uint, string, error)
}

type jwt struct {
	cfg *Config
}

func New(cfg *Config) JWT {
	return &jwt{cfg: cfg}
}

type payload struct {
	Email  string `json:"email"`
	UserId uint   `json:"userid"`
}

func (payload *payload) Valid() error {
	return nil
}

func (jwt *jwt) CreateTokenString(userId uint, email string) (string, error) {
	p := &payload{UserId: userId, Email: email}
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, p)
	return token.SignedString([]byte(jwt.cfg.SecretKey))
}

const (
	inValidToken        = "invalid token"
	errorMappingPayload = "error mapping the payload"
)

func (jwt *jwt) ExtractTokenData(tokenString string) (uint, string, error) {
	checkSigningMethod := func(token *jwtPkg.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtPkg.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", token.Header["alg"])
		}
		return []byte(jwt.cfg.SecretKey), nil
	}

	token, err := jwtPkg.ParseWithClaims(tokenString, &payload{}, checkSigningMethod)
	if err != nil {
		errStr := fmt.Sprintf("error: %v, token: %s", err, tokenString)
		return 0, "", errors.New(errStr)
	}

	if !token.Valid {
		errStr := fmt.Sprintf("%s, token: %v", inValidToken, token)
		return 0, "", errors.New(errStr)
	}

	p, ok := token.Claims.(*payload)
	if !ok {
		errStr := fmt.Sprintf("%s: %s, token: %v", inValidToken, errorMappingPayload, token)
		return 0, "", errors.New(errStr)
	}

	return p.UserId, p.Email, nil
}
