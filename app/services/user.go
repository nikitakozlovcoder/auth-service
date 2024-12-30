package services

import (
	"context"
	"errors"
	"log/slog"
	"service/auth/app/domain/dtos"
	"service/auth/app/domain/requests"
	"service/auth/app/repositories"
	"slices"
	"strconv"
)

const (
	UerIdJwtClaim = "userId"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserStorageAccessor interface {
	UserWithHash(ctx context.Context, username string) (dtos.UserWithHash, error)
}

type PasswordComparer interface {
	Compare(passwordHash, password string) (bool, error)
}

type JwtGenerator interface {
	Generate(map[string]string) (string, error)
}

type UserService struct {
	userStorageAccessor UserStorageAccessor
	jwtGenerator        JwtGenerator
	passwordComparer    PasswordComparer
}

func NewUserService(userStorageAccessor UserStorageAccessor, jwtGenerator JwtGenerator, passwordComparer PasswordComparer) *UserService {
	return &UserService{
		userStorageAccessor: userStorageAccessor,
		jwtGenerator:        jwtGenerator,
		passwordComparer:    passwordComparer,
	}
}

func (s *UserService) Login(ctx context.Context, req requests.LoginRequest) (string, error) {
	user, err := s.userStorageAccessor.UserWithHash(ctx, req.Username)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}

		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return "", err
	}

	isValidPassword, err := s.passwordComparer.Compare(user.PasswordHash, req.Password)
	if err != nil {
		return "", err
	}

	if !isValidPassword {
		return "", ErrInvalidCredentials
	}

	jwt, err := s.jwtGenerator.Generate(map[string]string{UerIdJwtClaim: strconv.FormatInt(user.Id, 10)})
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return "", err
	}

	return jwt, nil
}
