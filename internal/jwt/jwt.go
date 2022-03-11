package jwt

import (
	"errors"
	"fmt"

	jwtPkg "github.com/golang-jwt/jwt/v4"
)

type JWT interface {
	CreateTokenString(userId uint, email string) (string, error)
	ExtractTokenData(tokenString string) (*Payload, error)
}

type jwt struct {
	cfg *Config
}

func New(cfg *Config) JWT {
	return &jwt{cfg: cfg}
}

func (jwt *jwt) CreateTokenString(userId uint, email string) (string, error) {
	payload := &Payload{UserId: userId, Email: email}
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, payload)
	return token.SignedString([]byte(jwt.cfg.SecretKey))
}

const (
	inValidToken        = "invalid token"
	errorMappingPayload = "error mapping the payload"
	errorUnmarshalData  = "error unmarshaling the data"
)

func (jwt *jwt) ExtractTokenData(tokenString string) (*Payload, error) {
	checkSigningMethod := func(token *jwtPkg.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtPkg.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", token.Header["alg"])
		}
		return jwt.cfg.SecretKey, nil
	}

	token, err := jwtPkg.ParseWithClaims(tokenString, &Payload{}, checkSigningMethod)
	if err != nil {
		errStr := fmt.Sprintf("error: %v, token: %s", err, tokenString)
		return nil, errors.New(errStr)
	}

	if !token.Valid {
		errStr := fmt.Sprintf("%s, token: %v", inValidToken, token)
		return nil, errors.New(errStr)
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		errStr := fmt.Sprintf("%s: %s, token: %v", inValidToken, errorMappingPayload, token)
		return nil, errors.New(errStr)
	}

	return payload, nil
}
