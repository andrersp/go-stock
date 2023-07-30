package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDomain(t *testing.T) {

	t.Run("TesteUserDomainSuccess", func(t *testing.T) {

		user := User{
			UserName: "rspandre",
			Email:    "rspandre@email.com",
			Password: "minhasenha123",
		}

		err := user.Validate()

		assert.Equal(t, err, nil)

	})

}
