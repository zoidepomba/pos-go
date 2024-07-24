package database

import (
	"github.com/zoidepomba/pos-go/tree/main/9/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}