package api

import (
	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/go-playground/validator"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {
		errorDetail := normalizeError(err)
		err := domain.NewAppError("VALIDATION_ERROR", errorDetail)
		return err

	}

	return nil
}

func normalizeError(err error) (errorDetail string) {

	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		switch e.Tag() {
		case "required":
			errorDetail = e.Field() + " required!"
			return
		case "email":
			errorDetail = e.Field() + " invalid!"
			return
		case "min":
			errorDetail = e.Field() + " must have " + e.Param() + " characters at least!"
			return
		case "eqfield":
			errorDetail = e.Field() + " does not match with " + e.Param()
		default:
			errorDetail += "error on field " + e.Tag() + " " + e.Field()

		}
	}
	return

}
