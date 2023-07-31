package user

import (
	"testing"

	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/stretchr/testify/assert"
)

const (
	validEmail    = "myemail@mail.com"
	invalidEmail  = "myemail.com"
	validUserName = "myusername"
	newPassword   = "new_password"
)

func TestUserDomain(t *testing.T) {

	type testCase struct {
		testName string
		userName string
		email    string
		password string
		expected error
	}

	testCases := []testCase{
		{
			testName: "CreateUserSuccess",
			userName: validUserName,
			email:    validEmail,
			password: "mypassword",
			expected: nil,
		},
		{
			testName: "EmptyUserName",
			userName: "",
			email:    validEmail,
			password: "mypassword",
			expected: domain.NewAppError("VALIDATION_ERROR", errEmptyUsername),
		},
		{
			testName: "EmptyEmail",
			userName: validUserName,
			email:    "",
			password: "mypassword",
			expected: domain.NewAppError("VALIDATION_ERROR", errEmptyEmail),
		},
		{
			testName: "InvalidEmail",
			userName: validUserName,
			email:    "myemail.com",
			password: "mypassword",
			expected: domain.NewAppError("VALIDATION_ERROR", errInvalidEmail),
		},
		{
			testName: "InvalidPassword",
			userName: validUserName,
			email:    validEmail,
			password: "",
			expected: domain.NewAppError("VALIDATION_ERROR", errEmptyPassword),
		},
		{
			testName: "InvalidPasswordMinimumChars",
			userName: "myuser",
			email:    validEmail,
			password: "passw",
			expected: domain.NewAppError("VALIDATION_ERROR", errInvalidPassowrd),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewUser(
				test.userName,
				test.password,
				test.email,
			)

			assert.Equal(t, err, test.expected)
		})
	}

	t.Run("when_change_passowrd_return_success", func(t *testing.T) {

		user, _ := NewUser(validUserName, "mypassword", validEmail)

		err := user.SetPassword(newPassword)
		assert.Nil(t, err)
		assert.EqualValues(t, user.GetPassword(), newPassword)
	})

	t.Run("when_change_passowrd_return_err_empty_new_password", func(t *testing.T) {
		user, _ := NewUser(validUserName, "mypassword", validEmail)

		err := user.SetPassword("")
		assert.NotNil(t, err)
	})

	t.Run("when_change_passowrd_return_err_minimum_chars", func(t *testing.T) {
		user, _ := NewUser(validUserName, "mypassword", validEmail)

		err := user.SetPassword("1234")
		assert.NotNil(t, err)
	})

	t.Run("when_change_enable", func(t *testing.T) {
		user, _ := NewUser(validUserName, "mypassword", validEmail)

		user.SetEnable(false)
		assert.False(t, user.IsEnable())
	})

}
