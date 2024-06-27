package entity

import (
	"github.com/zoidepomba/pos-go/tree/main/9/internal/entity"
	"github.com/zoidepomba/pos-go/tree/main/9/pkg/entity"
	"golang.org/x/crypto/bycrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) *User {
	hash, err := bycrypt.GenerateFromPassword([]byte(password), bycrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bycrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
