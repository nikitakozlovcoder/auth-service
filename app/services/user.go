package services

import (
	"context"
	"errors"
	"log/slog"
	"service/auth/app/domain/requests"
)

const (
	UerIdJwtClaim = "userId"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserStorageAccessor interface {
	UserExists(ctx context.Context, username, passwordHash string) (bool, error)
}

type JwtGenerator interface {
	Generate(map[string]string) (string, error)
}

type UserService struct {
	userStorageAccessor UserStorageAccessor
	jwtGenerator        JwtGenerator
}

func NewUserService(userStorageAccessor UserStorageAccessor, jwtGenerator JwtGenerator) *UserService {
	return &UserService{
		userStorageAccessor: userStorageAccessor,
		jwtGenerator:        jwtGenerator,
	}
}

func (s *UserService) Login(ctx context.Context, req requests.LoginRequest) (string, error) {
	exists, err := s.userStorageAccessor.UserExists(ctx, req.Username, req.Password)
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return "", err
	}

	if !exists {
		return "", ErrInvalidCredentials
	}

	jwt, err := s.jwtGenerator.Generate(map[string]string{UerIdJwtClaim: req.Username})
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return "", err
	}

	return jwt, nil
}
