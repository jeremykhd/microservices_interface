package grpc

import (
	"user-service/internal/application/command"
	"user-service/internal/application/query"
)

type UserHandler struct {
	command.CreateUserHandler
	query.GetUserHandler
}

func NewUserHandler(createUserHandler command.CreateUserHandler, getUserHandler query.GetUserHandler) *UserHandler {
	return &UserHandler{
		CreateUserHandler: createUserHandler,
		GetUserHandler:    getUserHandler,
	}
}
