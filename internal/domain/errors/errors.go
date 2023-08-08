package domain

import "fmt"

var ERRORS = map[string]map[string]string{
	"VALIDATION_ERROR": {"name": "VALIDATION_ERROR", "description": "indicates a error in input validation.", "detail": ""},
	"RESOURCE_EXISTS":  {"name": "RESOURCE_EXISTS", "description": "indicates a duplicate / already existing record", "detail": ""},
	"NOT_FOUND":        {"name": "NOT_FOUND", "description": " indicates a missing / not found record", "detail": ""},
	"INTERNAL_ERROR":   {"name": "INTERNAL_ERROR", "description": "internal error", "detail": ""},
	"INVALID_PARAM":    {"name": "INVALID_PARAM", "description": "invalid param on request", "detail": ""},
	"INVALID_USERNAME": {"name": "INVALID_USERNAME", "description": "invalid username", "detail": ""},
	"USER_NOT_FOUND":   {"name": "USER_NOT_FOUND", "description": "user not found", "detail": ""},
	"LOGIN_FAIL":       {"name": "LOGIN_FAIL", "description": "invalid username or password", "detail": ""},
	"UNAUTHORIZED":     {"name": "UNAUTHORIZED", "description": "missing or malformed token.", "detail": "token invalid, missing or malformed"},
	"EXPIRED_TOKEN":    {"name": "EXPIRED_TOKEN", "description": "the token has expired", "detail": ""},
}

type AppError struct {
	ErrorName   string `json:"errorName" example:"GENERIC_ERROR"`
	Description string `json:"description" example:"Generic error os proccess request."`
	Detail      string `json:"detail"`
} //@name AppError

func (e *AppError) Error() string {
	return fmt.Sprintf("errorName: %s, description: %s", e.ErrorName, e.Description)
}

func NewAppError(errorName, detail string) *AppError {
	appEroor := &AppError{}

	if _, ok := ERRORS[errorName]; !ok {
		errorName = "GENERIC_ERROR"
	}

	err := ERRORS[errorName]

	appEroor.ErrorName = err["name"]
	appEroor.Description = err["description"]

	if detail != "" {
		appEroor.Detail = detail
	} else {
		appEroor.Detail = err["detail"]
	}

	return appEroor

}
