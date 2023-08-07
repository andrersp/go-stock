package repository

import (
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user *user.User) (*user.User, error) {

	model := userDomainToModel(user)

	err := ur.db.Create(&model).Error

	if err != nil {
		return nil, err
	}

	return model.toDomainModel(), nil
}

func (ur *userRepository) GetUserByID(uuid.UUID) (*user.User, error) {
	return nil, nil
}

func (ur *userRepository) GetUserByUserName(userName string) (*user.User, error) {
	var user UserModel

	err := ur.db.Where("user_name = ?", userName).First(&user).Error

	return user.toDomainModel(), err
}

func (ur *userRepository) GetUserByEmail(email string) (*user.User, error) {

	var user UserModel

	err := ur.db.Where("email = ?", email).First(&user).Error

	return user.toDomainModel(), err
}

func (ur *userRepository) ListUsers() ([]user.User, error) {

	var model []UserModel
	users := make([]user.User, 0)
	err := ur.db.Find(&model).Error

	for _, u := range model {
		users = append(users,
			*u.toDomainModel())
	}

	return users, err
}
