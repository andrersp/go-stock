package userrepository

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

	model := fromEntity(user)
	err := ur.db.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model.toEntity(), nil
}

func (ur *userRepository) GetUserByID(uuid.UUID) (*user.User, error) {
	return nil, nil
}

func (ur *userRepository) GetUserByUserName(userName string) (user *user.User, err error) {

	var model User
	err = ur.db.Where("user_name = ?", userName).First(&model).Error
	if err != nil {
		return
	}

	user = model.toEntity()

	return
}

func (ur *userRepository) GetUserByEmail(email string) (user *user.User, err error) {

	var model User
	err = ur.db.Where("email = ?", email).First(&model).Error
	if err != nil {
		return
	}

	user = model.toEntity()

	return
}

func (ur *userRepository) ListUsers() ([]user.User, error) {

	var model []User
	users := make([]user.User, 0)
	err := ur.db.Find(&model).Error
	if err != nil {
		return nil, err
	}

	for _, u := range model {
		users = append(users,
			*u.toEntity())
	}

	return users, err
}
