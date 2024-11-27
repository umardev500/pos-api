package repository

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type authRepository struct {
	db *pkg.GormDB
}

func NewAuthRepository(db *pkg.GormDB) contract.AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) GetUserByUsernameOrEmail(ctx context.Context, username string) (*model.User, error) {
	conn := a.db.GetConn(ctx)

	var user model.User
	result := conn.Where("username = ?", username).Or("email = ?", username).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
