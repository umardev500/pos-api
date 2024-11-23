package repository

import (
	"context"

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

func (r *userRepository) FindUserById(ctx context.Context, id string) (*model.User, error) {
	sql := "SELECT id, username, email FROM users WHERE id = $1"
	row := r.db.GetConn(ctx).QueryRow(ctx, sql, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
