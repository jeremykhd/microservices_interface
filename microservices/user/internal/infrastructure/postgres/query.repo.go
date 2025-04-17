package postgres

import (
	"context"
	"database/sql"
	"user-service/internal/domain"
)

type UserQueryRepository struct {
	db *sql.DB
}

func NewUserQueryRepository(db *sql.DB) *UserQueryRepository {
	return &UserQueryRepository{db: db}
}

func (r *UserQueryRepository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	query := `
	SELECT 	*
	FROM users
	WHERE id = $1
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserQueryRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
	SELECT *
	FROM users
	WHERE email = $1
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
