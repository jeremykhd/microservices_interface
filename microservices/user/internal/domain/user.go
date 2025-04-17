package domain

import (
	"context"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCommandRepository interface {
	Create(ctx context.Context, user User) (*User, error)
}

type UserQueryRepository interface {
	Get(ctx context.Context, id string) (*User, error)
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
