package jwttoken

import (
	"time"

	"github.com/andrersp/go-stock/internal/config"
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/golang-jwt/jwt/v4"
)

const (
	JWT_ACCESS_TOKEN_EXPIRED_TIME  = time.Hour * 24
	JWT_REFRESH_TOKEN_EXPIRED_TIME = time.Hour * 48
)

func CreateAccessToken(user *user.User) (token string, err error) {

	atClaims := jwtTokenClaims{
		ID: user.GetId(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_ACCESS_TOKEN_EXPIRED_TIME)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err = accessToken.SignedString([]byte(config.JWT_ACCESS_TOKEN_SECRET))
	return
}

func CreateRefreshToken(user *user.User) (token string, err error) {

	atClaims := jwtRefreshTokenClaims{
		ID: user.GetId(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_REFRESH_TOKEN_EXPIRED_TIME)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err = accessToken.SignedString([]byte(config.JWT_REFRESH_TOKEN_SECRET))
	return
}
