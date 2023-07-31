package user

import (
	"testing"

	domain "github.com/andrersp/go-stock/internal/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServiceUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repository := NewMockUserRepository(control)
	service := NewUserService(repository)
	userDomain := User{
		id:       uuid.New(),
		userName: validUserName,
		password: "mypassword",
		email:    validEmail,
	}

	t.Run("when_user_already_exists_return_error", func(t *testing.T) {

		repository.EXPECT().GetUserByUserName(userDomain.userName).Return(&userDomain, nil)
		_, err := service.CreateUser(&userDomain)
		assert.NotNil(t, err)

	})

	t.Run("when_user_is_not_registered_return_success", func(t *testing.T) {

		repository.EXPECT().GetUserByUserName(userDomain.userName).Return(nil, nil)
		repository.EXPECT().CreateUser(&userDomain).Return(&userDomain, nil)
		_, err := service.CreateUser(&userDomain)
		assert.Nil(t, err)

	})

	t.Run("when_get_user_by_id_return_not_found", func(t *testing.T) {

		repository.EXPECT().GetUserByID(userDomain.id).Return(nil, domain.NewAppError("NOT_FOUND", ""))
		sut, err := service.GetUserByID(userDomain.id)
		assert.Nil(t, sut)
		assert.NotNil(t, err)

	})

	t.Run("when_get_user_by_id_return_success", func(t *testing.T) {

		userID := uuid.New()
		userDomain.id = userID
		repository.EXPECT().GetUserByID(userDomain.id).Return(&userDomain, nil)
		sut, _ := service.GetUserByID(userID)
		assert.NotNil(t, sut)

	})

}
