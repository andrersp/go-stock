package jwttoken

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtTokenClaims struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

type JwtRefreshTokenClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
