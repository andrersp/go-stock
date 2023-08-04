package utils

import "github.com/labstack/echo/v4"

type RouterModel struct {
	Path         string
	Method       string
	Func         func(echo.Context) error
	AuthRequired bool
}
