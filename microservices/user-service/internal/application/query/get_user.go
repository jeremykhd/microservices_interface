package query

import (
	"context"
	"user-service/internal/domain"
)

type GetUserQuery struct {
	ID string `json:"id"`
}

type GetUserHandler struct {
	repo domain.UserQueryRepository
}

func (h *GetUserHandler) Handle(ctx context.Context, query *GetUserQuery) (*domain.User, error) {
	return h.repo.Get(ctx, query.ID)
}
