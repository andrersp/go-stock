package userrepository

import (
	"github.com/andrersp/go-stock/internal/domain/user"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;autoIncrement:false;index;"`
	UserName string    `gorm:"size:40;"`
	Password string
	Email    string `gorm:"size:200;"`
	Enable   bool   `gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) toEntity() *user.User {

	var user user.User

	user.SetEmail(u.Email)
	user.SetUserName(u.UserName)
	user.SetEnable(u.Enable)
	user.SetPassword(u.Password)
	user.SetID(u.Id)
	return &user
}

func fromEntity(user *user.User) *User {
	return &User{
		Id:       user.GetId(),
		UserName: user.GetUserName(),
		Password: user.GetPassword(),
		Email:    user.GetEmail(),
		Enable:   user.IsEnable(),
	}
}
