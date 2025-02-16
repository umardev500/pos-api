package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type userRepository struct {
	db *pkg.GormDB
}

func NewUserRepository(db *pkg.GormDB) contract.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := r.db.GetConn(ctx).Create(user)

	return result.Error
}

func (r *userRepository) FindAllUsers(ctx context.Context, params pkg.FindRequest) ([]model.User, int64, error) {
	pagination := params.Pagination

	conn := r.db.GetConn(ctx)
	var users = make([]model.User, 0)
	var count int64 = 0

	result := conn.Offset(int(pagination.Offset)).Limit(int(pagination.PerPage)).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	conn.Model(&model.User{}).Count(&count)

	return users, count, nil
}

func (r *userRepository) FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	r.db.GetConn(ctx).First(&user, id)

	return &user, nil
}
