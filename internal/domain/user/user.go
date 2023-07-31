package user

import (
	"net/mail"
	"strings"

	domain "github.com/andrersp/go-stock/internal/domain/errors"

	"github.com/google/uuid"
)

const (
	errEmptyUsername     = "Username not be empty"
	errEmptyEmail        = "Email not be empty"
	errInvalidEmail      = "Invalid email"
	errEmptyPassword     = "Password not be empty"
	errInvalidPassowrd   = "password must be at least 6 characters long"
	errPasswordDontMatch = "old password dont macth"
)

type User struct {
	id       uuid.UUID
	userName string
	password string
	email    string
	enable   bool
}

func (u *User) GetId() uuid.UUID {
	return u.id
}

func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) IsEnable() bool {
	return u.enable
}

func (u *User) SetUserName(userName string) error {
	userName = strings.TrimSpace(userName)
	if userName == "" {
		err := domain.NewAppError("VALIDATION_ERROR", errEmptyUsername)
		return err
	}

	u.userName = userName
	return nil

}

func (u *User) SetEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		err := domain.NewAppError("VALIDATION_ERROR", errEmptyEmail)
		return err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		err := domain.NewAppError("VALIDATION_ERROR", errInvalidEmail)
		return err
	}

	u.email = email
	return nil

}
func (u *User) SetPassword(password string) error {
	if password == "" {
		err := domain.NewAppError("VALIDATION_ERROR", errEmptyPassword)
		return err
	}
	if len(password) < 6 {
		err := domain.NewAppError("VALIDATION_ERROR", errInvalidPassowrd)
		return err
	}

	u.password = password
	return nil
}

func (u *User) SetEnable(enable bool) {
	u.enable = enable
}

func NewUser(userName, password, email string) (user *User, err error) {

	userName = strings.TrimSpace(userName)
	email = strings.TrimSpace(email)

	if userName == "" {
		err = domain.NewAppError("VALIDATION_ERROR", errEmptyUsername)
		return
	}

	if email == "" {
		err = domain.NewAppError("VALIDATION_ERROR", errEmptyEmail)
		return
	}

	if _, err = mail.ParseAddress(email); err != nil {
		err = domain.NewAppError("VALIDATION_ERROR", errInvalidEmail)
		return
	}

	if password == "" {
		err = domain.NewAppError("VALIDATION_ERROR", errEmptyPassword)
		return
	}
	if len(password) < 6 {
		err = domain.NewAppError("VALIDATION_ERROR", errInvalidPassowrd)
		return
	}

	return &User{
		id:       uuid.New(),
		userName: userName,
		password: password,
		email:    email,
		enable:   true,
	}, nil

}
