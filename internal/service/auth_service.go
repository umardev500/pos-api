package service

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type authService struct {
	repo contract.AuthRepository
	v    pkg.Validator
}

func NewAuthService(repo contract.AuthRepository, v pkg.Validator) contract.AuthService {
	return &authService{
		repo: repo,
		v:    v,
	}
}

func (a *authService) Login(ctx context.Context, req *model.LoginRequest) pkg.Response {
	// Validate
	fields, err := a.v.Struct(req)
	if err != nil {
		return pkg.ValidationErrorResponse(fields)
	}

	// Get user
	// Generate token

	return pkg.Response{}
}
