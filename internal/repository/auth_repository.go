package repository

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/pkg"
)

type authRepository struct {
	db *pkg.PGX
}

func NewAuthRepository(db *pkg.PGX) contract.AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) GetUserByUsernameOrEmail(ctx context.Context, username string, email string) (string, error) {
	conn := a.db.GetConn(ctx)

	sql := "SELECT password_hash FROM users WHERE username = $1 OR email = $2"
	row := conn.QueryRow(ctx, sql, username, email)

	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
