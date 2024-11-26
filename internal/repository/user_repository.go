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

func (r *userRepository) FindAllUsers(ctx context.Context) ([]model.User, error) {
	sql := "SELECT id, username, email, created_at FROM users"
	rows, err := r.db.GetConn(ctx).Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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
