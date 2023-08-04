package utils

import (
	"errors"

	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/labstack/echo/v4"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {

		var httpError *echo.HTTPError
		var detail string

		if errors.As(err, &httpError) {
			detail = string(httpError.Internal.Error())
		} else {
			detail = err.Error()
		}

		err = domain.NewAppError("VALIDATION_ERROR", detail)

		return err
	}

	return
}
