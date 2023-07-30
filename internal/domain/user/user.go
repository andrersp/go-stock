package user

import (
	"net/mail"
	"strings"

	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/google/uuid"
)

const (
	emptyUsername = "Username not be empty"
	emptyEmail    = "Email not be empty"
	invalidEmail  = "Email not be empty"
	emptyPassowrd = "Password not be empty and mininum 6 chars"
)

type User struct {
	Id       uuid.UUID
	UserName string
	Password string
	Email    string
	Enable   bool
}

func (u *User) Validate() error {

	u.UserName = strings.TrimSpace(u.UserName)
	u.Email = strings.TrimSpace(u.Email)

	if u.UserName == "" {
		err := domain.NewAppError("INVALID_PARAM", emptyUsername)
		return err
	}

	if u.Email == "" {
		err := domain.NewAppError("INVALID_PARAM", emptyEmail)
		return err
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		err := domain.NewAppError("INVALID_PARAM", invalidEmail)
		return err
	}

	if u.Password == "" || len(u.Password) < 6 {
		err := domain.NewAppError("INVALID_PARAM", emptyPassowrd)
		return err

	}
	return nil

}
