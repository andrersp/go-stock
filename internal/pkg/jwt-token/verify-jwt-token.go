package jwttoken

import (
	"net/http"
	"strings"

	"github.com/andrersp/go-stock/internal/config"
	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyJwtToken(r *http.Request) (user *JwtTokenClaims, err error) {

	tokenString, err := extractTokenString(r)
	user, err = parseToken(tokenString)

	if err != nil {
		return
	}

	return

}

func extractTokenString(r *http.Request) (tokenString string, err error) {

	if _, ok := r.Header["Authorization"]; !ok {
		err = getError()
		return
	}

	authorizationString := r.Header.Get("Authorization")

	splitAuthorization := strings.Split(authorizationString, " ")

	if len(splitAuthorization) != 2 {
		err = getError()
		return
	}

	if splitAuthorization[0] != "Bearer" {
		getError()
		return
	}

	tokenString = splitAuthorization[1]
	return

}

func parseToken(tokenString string) (tokenClaims *JwtTokenClaims, err error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = getError()
			return nil, err
		}
		return []byte(config.JWT_ACCESS_TOKEN_SECRET), nil

	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtTokenClaims{}, keyFunc)

	if err != nil {

		verr, _ := err.(*jwt.ValidationError)
		if verr.Errors == jwt.ValidationErrorExpired {
			err = domain.NewAppError("EXPIRED_TOKEN", "")
			return
		}

		err = getError()
		return
	}

	tokenClaims, ok := token.Claims.(*JwtTokenClaims)

	if !ok {
		err = getError()
		return
	}

	return
}

func getError() error {
	return domain.NewAppError("UNAUTHORIZED", "")
}
