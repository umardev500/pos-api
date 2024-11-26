package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type userRepository struct {
	db *pkg.PGX
}

func NewUserRepository(db *pkg.PGX) contract.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	sql := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)"
	_, err := r.db.GetConn(ctx).Exec(ctx, sql, user.Username, user.Email, user.PasswordHash)
	return err
}

func (r *userRepository) FindAllUsers(ctx context.Context, params pkg.FindRequest) ([]model.User, int64, error) {
	pagination := params.Pagination

	sql := `
	SELECT
		u.id, u.username, u.email, u.created_at 
	FROM users u
	LIMIT $1 OFFSET $2
	`
	rows, err := r.db.GetConn(ctx).Query(ctx, sql, pagination.PerPage, pagination.Offset)
	if err != nil {
		return nil, 0, err
	}

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	// Count
	sql = "SELECT COUNT(*) FROM users"
	var count int64
	err = r.db.GetConn(ctx).QueryRow(ctx, sql).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *userRepository) FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	sql := "SELECT id, username, email, created_at FROM users WHERE id = $1"
	row := r.db.GetConn(ctx).QueryRow(ctx, sql, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
