package repository

import (
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/google/uuid"
)

type UserModel struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;autoIncrement:false;index;"`
	UserName string    `gorm:"size:40;"`
	Password string
	Email    string `gorm:"size:200;"`
	Enable   bool   `gorm:"default:false"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) toDomainModel() *user.User {

	var user user.User

	user.SetEmail(u.Email)
	user.SetUserName(u.UserName)
	user.SetEnable(u.Enable)
	user.SetPassword(u.Password)
	user.SetID(u.Id)
	return &user
}

func userDomainToModel(user *user.User) *UserModel {
	return &UserModel{
		Id:       user.GetId(),
		UserName: user.GetUserName(),
		Password: user.GetPassword(),
		Email:    user.GetEmail(),
		Enable:   user.IsEnable(),
	}
}
