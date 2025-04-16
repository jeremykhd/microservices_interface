package postgres

import (
	"context"
	"database/sql"
	"user-service/internal/domain"
)

type UserCommandRepository struct {
	db *sql.DB
}

func NewUserCommandRepository(db *sql.DB) *UserCommandRepository {
	return &UserCommandRepository{db: db}
}

func (r *UserCommandRepository) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	query := `
	INSERT INTO users (id, name, email)
	VALUES ($1, $2, $3)
	RETURNING id, name, email
	`

	var createdUser domain.User
	err := r.db.QueryRowContext(ctx, query, user.ID, user.Name, user.Email).Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}
