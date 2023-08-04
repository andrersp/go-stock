package utils

import (
	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.Validator.Struct(i); err != nil {
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
		default:
			errorDetail += "error on field " + e.Tag() + " " + e.Field()

		}
	}
	return

}
