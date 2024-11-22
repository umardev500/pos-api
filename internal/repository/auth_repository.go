package repository

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type authRepository struct {
	db *pkg.PGX
}

func NewAuthRepository(db *pkg.PGX) contract.AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) GetUserByUsernameOrEmail(ctx context.Context, username string) (*model.User, error) {
	conn := a.db.GetConn(ctx)

	sql := "SELECT id, username, email ,password_hash FROM users WHERE username = $1 OR email = $2"
	row := conn.QueryRow(ctx, sql, username, username)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
