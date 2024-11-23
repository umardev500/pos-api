package service

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	user, err := a.repo.GetUserByUsernameOrEmail(ctx, req.Username)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	// Check password
	if !pkg.CheckPasswordHash(req.Password, user.PasswordHash) {
		return pkg.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "Username or password is incorrect",
		}
	}

	// Generate token
	jwtClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(60 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	secret := os.Getenv("JWT_SECRET")

	token, err := pkg.CreateJWT(jwtClaims, secret)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	// Return
	var loginResponse model.LoginResponse = model.LoginResponse{
		Token: token,
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Login successful",
		Data:       loginResponse,
	}
}
