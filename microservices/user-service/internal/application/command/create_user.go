package command

import (
	"context"
	"user-service/internal/domain"
)

type CreateUserCommand struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserHandler struct {
	repo domain.UserCommandRepository
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd *CreateUserCommand) (*domain.User, error) {
	user := domain.NewUser(
		cmd.Name,
		cmd.Email,
	)

	return h.repo.Create(ctx, *user)
}
