package jwttoken

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type jwtTokenClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type jwtRefreshTokenClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
